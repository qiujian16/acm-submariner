package submarineragent

import (
	"context"
	"fmt"
	clientset "github.com/open-cluster-management/api/client/cluster/clientset/versioned"
	clusterinformerv1 "github.com/open-cluster-management/api/client/cluster/informers/externalversions/cluster/v1"
	clusterinformerv1alpha1 "github.com/open-cluster-management/api/client/cluster/informers/externalversions/cluster/v1alpha1"
	clusterlisterv1 "github.com/open-cluster-management/api/client/cluster/listers/cluster/v1"
	clusterlisterv1alpha1 "github.com/open-cluster-management/api/client/cluster/listers/cluster/v1alpha1"
	workv1client "github.com/open-cluster-management/api/client/work/clientset/versioned"
	workinformer "github.com/open-cluster-management/api/client/work/informers/externalversions/work/v1"
	worklister "github.com/open-cluster-management/api/client/work/listers/work/v1"
	clusterv1 "github.com/open-cluster-management/api/cluster/v1"
	clusterv1alpha1 "github.com/open-cluster-management/api/cluster/v1alpha1"
	"k8s.io/client-go/dynamic"
	"path/filepath"

	"github.com/openshift/library-go/pkg/assets"
	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/resource/resourceapply"
	operatorhelpers "github.com/openshift/library-go/pkg/operator/v1helpers"

	"github.com/qiujian16/acm-submariner/pkg/helpers"
	"github.com/qiujian16/acm-submariner/pkg/hub/submarineragent/bindata"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
)

const (
	agentFinalizer      = "cluster.open-cluster-management.io/submariner-agent-cleanup"
	serviceAccountLabel = "cluster.open-cluster-management.io/submariner-cluster-sa"
)

var clusterRBACFiles = []string{
	"manifests/agent/rbac/submariner-cluster-serviceaccount.yaml",
	"manifests/agent/rbac/submariner-cluster-rolebinding.yaml",
}

type clusterRBACConfig struct {
	ManagedClusterName        string
	SubmarinerBrokerNamespace string
}

// submarinerAgentController reconciles instances of ManagedCluster on the hub to deploy/remove
// corresponding submariner agent manifestworks
type submarinerAgentController struct {
	kubeClient         kubernetes.Interface
	dynamicClient      dynamic.Interface
	clusterClient      clientset.Interface
	manifestWorkClient workv1client.Interface
	clusterLister      clusterlisterv1.ManagedClusterLister
	clusterSetLister   clusterlisterv1alpha1.ManagedClusterSetLister
	manifestWorkLister worklister.ManifestWorkLister
	eventRecorder      events.Recorder
}

// NewSubmarinerAgentController returns a submarinerAgentController instance
func NewSubmarinerAgentController(
	kubeClient kubernetes.Interface,
	dynamicClient dynamic.Interface,
	clusterClient clientset.Interface,
	manifestWorkClient workv1client.Interface,
	clusterInformer clusterinformerv1.ManagedClusterInformer,
	clusterSetInformer clusterinformerv1alpha1.ManagedClusterSetInformer,
	manifestWorkInformer workinformer.ManifestWorkInformer,
	recorder events.Recorder) factory.Controller {
	c := &submarinerAgentController{
		kubeClient:         kubeClient,
		dynamicClient:      dynamicClient,
		clusterClient:      clusterClient,
		manifestWorkClient: manifestWorkClient,
		clusterLister:      clusterInformer.Lister(),
		clusterSetLister:   clusterSetInformer.Lister(),
		manifestWorkLister: manifestWorkInformer.Lister(),
		eventRecorder:      recorder.WithComponentSuffix("submariner-agent-controller"),
	}
	return factory.New().
		WithInformersQueueKeyFunc(func(obj runtime.Object) string {
			accessor, _ := meta.Accessor(obj)
			return accessor.GetName()
		}, clusterInformer.Informer()).
		WithInformers(clusterSetInformer.Informer()).
		WithInformers(manifestWorkInformer.Informer()).
		WithSync(c.sync).
		ToController("SubmarinerAgentController", recorder)
}

func (c *submarinerAgentController) sync(ctx context.Context, syncCtx factory.SyncContext) error {
	key := syncCtx.QueueKey()

	// if the sync is triggered by change of ManagedClusterSet or ManifestWork, reconcile all managed clusters
	if key == "key" {
		if err := c.syncAllManagedClusters(ctx); err != nil {
			return err
		}
		return nil
	}

	managedCluster, err := c.clusterLister.Get(key)
	if errors.IsNotFound(err) {
		// managed cluster not found, could have been deleted, do nothing.
		return nil
	}
	if err != nil {
		return err
	}

	if err := c.syncManagedCluster(ctx, managedCluster); err != nil {
		return err
	}

	return nil
}

// syncAllManagedClusters syncs all managed clusters
func (c *submarinerAgentController) syncAllManagedClusters(ctx context.Context) error {
	managedClusters, err := c.clusterLister.List(labels.Everything())
	if err != nil {
		return err
	}

	errs := []error{}
	for _, managedCluster := range managedClusters {
		if err = c.syncManagedCluster(ctx, managedCluster); err != nil {
			errs = append(errs, err)
		}
	}

	return operatorhelpers.NewMultiLineAggregate(errs)
}

// syncManagedCluster syncs one managed cluster
func (c *submarinerAgentController) syncManagedCluster(ctx context.Context, managedCluster *clusterv1.ManagedCluster) error {
	//TODO: selected the cluster with label

	// add a submariner agent finalizer to a managed cluster
	if managedCluster.DeletionTimestamp.IsZero() {
		hasFinalizer := false
		for i := range managedCluster.Finalizers {
			if managedCluster.Finalizers[i] == agentFinalizer {
				hasFinalizer = true
				break
			}
		}
		if !hasFinalizer {
			managedCluster.Finalizers = append(managedCluster.Finalizers, agentFinalizer)
			_, err := c.clusterClient.ClusterV1().ManagedClusters().Update(ctx, managedCluster, metav1.UpdateOptions{})
			return err
		}
	}

	// managed cluster is deleting, we remove its related resources
	if !managedCluster.DeletionTimestamp.IsZero() {
		// remove the submariner agent from this managedCluster
		if err := c.removeSubmarinerAgent(ctx, managedCluster); err != nil {
			return err
		}
		return c.removeAgentFinalizer(ctx, managedCluster)
	}

	// find the clustersets that contains this managed cluster
	matchedClusterSets, err := c.findClusterSet(managedCluster)
	if err != nil {
		return err
	}
	switch {
	case len(matchedClusterSets) == 0:
		// the managed cluster original clusterset may be deleted or changed, try to remove the submariner
		// agent from the managed cluter
		return c.removeSubmarinerAgent(ctx, managedCluster)
	case len(matchedClusterSets) == 1:
		// deploy the submariner agent on this managed cluster
		return c.deploySubmarinerAgent(ctx, matchedClusterSets[0], managedCluster)
	case len(matchedClusterSets) > 1:
		// case1: the managed cluter belongs more than one clusterset at the beginning, do nothing
		// case2: the managed cluter belongs more than one clusterset due to a certain clusterset is changed,
		// we keep the submariner agent, do nothing
		c.eventRecorder.Warning("SubmarinerUndeployed", fmt.Sprintf("There are one more than clustersets for managed cluster %q", managedCluster.Name))
		return nil
	}
	return nil
}

// removeAgentFinalizer removes the agent finalizer from a clusterset
func (c *submarinerAgentController) removeAgentFinalizer(ctx context.Context, managedCluster *clusterv1.ManagedCluster) error {
	copiedFinalizers := []string{}
	for i := range managedCluster.Finalizers {
		if managedCluster.Finalizers[i] == agentFinalizer {
			continue
		}
		copiedFinalizers = append(copiedFinalizers, managedCluster.Finalizers[i])
	}

	if len(managedCluster.Finalizers) != len(copiedFinalizers) {
		managedCluster.Finalizers = copiedFinalizers
		_, err := c.clusterClient.ClusterV1().ManagedClusters().Update(ctx, managedCluster, metav1.UpdateOptions{})
		return err
	}

	return nil
}

func (c *submarinerAgentController) findClusterSet(managedCluster *clusterv1.ManagedCluster) ([]*clusterv1alpha1.ManagedClusterSet, error) {
	clusterSets, err := c.clusterSetLister.List(labels.Everything())
	if err != nil {
		return nil, err
	}

	matchedSets := []*clusterv1alpha1.ManagedClusterSet{}
	for _, clusterSet := range clusterSets {
		selectedCluster := c.selectAllClusters(clusterSet)
		if selectedCluster.Has(managedCluster.Name) {
			matchedSets = append(matchedSets, clusterSet)
		}
	}

	return matchedSets, nil
}

func (c *submarinerAgentController) deploySubmarinerAgent(ctx context.Context, clusterSet *clusterv1alpha1.ManagedClusterSet, managedCluster *clusterv1.ManagedCluster) error {
	// generate service account and bind it to `submariner-k8s-broker-cluster` role
	brokerNamespace := fmt.Sprintf("submariner-clusterset-%s-broker", clusterSet.Name)
	if err := c.applyClusterRBACFiles(brokerNamespace, managedCluster.Name); err != nil {
		return err
	}

	config, err := NewSubmarinerConfig(c.kubeClient, c.dynamicClient, managedCluster.Name, brokerNamespace)
	if err != nil {
		c.eventRecorder.Warning("SubmarinerAgentDeployedFailed",
			fmt.Sprintf("failed to get config of submariner agent on managed cluster %v: %v", managedCluster.Name, err))
		return err
	}

	if err := ApplySubmarinerManifestWorks(config, c.manifestWorkClient, ctx); err != nil {
		c.eventRecorder.Warning("SubmarinerAgentDeployedFailed",
			fmt.Sprintf("failed to deploy submariner agent on managed cluster %v: %v", managedCluster.Name, err))
		return err
	}

	c.eventRecorder.Event("SubmarinerAgentDeployed", fmt.Sprintf("submariner agent was deployed on managed cluster %q", managedCluster.Name))
	return nil
}

func (c *submarinerAgentController) removeSubmarinerAgent(ctx context.Context, managedCluster *clusterv1.ManagedCluster) error {
	errs := []error{}
	if err := RemoveSubmarinerManifestWorks(managedCluster.Name, c.manifestWorkClient, ctx); err != nil {
		errs = append(errs, fmt.Errorf("failed to remove submariner agent from managed cluster %v: %v", managedCluster.Name, err))
	}

	// remove service account and its rolebinding from broker namespace
	if err := c.removeClusterRBACFiles(managedCluster.Name, ctx); err != nil {
		errs = append(errs, err)
	}

	if len(errs) == 0 {
		c.eventRecorder.Event("SubmarinerAgentRemoved", fmt.Sprintf("submariner agent was removed from managed cluster %q", managedCluster.Name))
	}
	return operatorhelpers.NewMultiLineAggregate(errs)
}

func (c *submarinerAgentController) applyClusterRBACFiles(brokerNamespace, managedClusterName string) error {
	config := &clusterRBACConfig{
		ManagedClusterName:        managedClusterName,
		SubmarinerBrokerNamespace: brokerNamespace,
	}
	clientHolder := resourceapply.NewKubeClientHolder(c.kubeClient)
	applyResults := resourceapply.ApplyDirectly(
		clientHolder,
		c.eventRecorder,
		func(name string) ([]byte, error) {
			return assets.MustCreateAssetFromTemplate(name, bindata.MustAsset(filepath.Join("", name)), config).Data, nil
		},
		clusterRBACFiles...,
	)
	errs := []error{}
	for _, result := range applyResults {
		if result.Error != nil {
			errs = append(errs, fmt.Errorf("%q (%T): %v", result.File, result.Type, result.Error))
		}
	}
	return operatorhelpers.NewMultiLineAggregate(errs)
}

func (c *submarinerAgentController) removeClusterRBACFiles(managedClusterName string, ctx context.Context) error {
	serviceAccounts, err := c.kubeClient.CoreV1().ServiceAccounts(metav1.NamespaceAll).List(ctx, metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", serviceAccountLabel, managedClusterName),
	})
	if errors.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}

	if len(serviceAccounts.Items) != 1 {
		return fmt.Errorf("one more than service accounts are found for %q", managedClusterName)
	}

	config := &clusterRBACConfig{
		ManagedClusterName:        managedClusterName,
		SubmarinerBrokerNamespace: serviceAccounts.Items[0].Namespace,
	}

	return helpers.CleanUpSubmarinerManifests(
		context.TODO(),
		c.kubeClient,
		c.eventRecorder,
		func(name string) ([]byte, error) {
			return assets.MustCreateAssetFromTemplate(name, bindata.MustAsset(filepath.Join("", name)), config).Data, nil
		},
		clusterRBACFiles...,
	)
}

// selectAllClusters selects all of clusters from a clusterset
func (c *submarinerAgentController) selectAllClusters(clusterSet *clusterv1alpha1.ManagedClusterSet) sets.String {
	selectedClusters := sets.NewString()
	for _, selector := range clusterSet.Spec.ClusterSelectors {
		clusters, err := c.selectClusters(selector)
		if err != nil {
			c.eventRecorder.Warning("ClustersSelectedFailed", fmt.Sprintf("failed to select clusters with %v: %v", selector, err))
		}
		selectedClusters.Insert(clusters...)
	}

	return selectedClusters
}

// selectClusters selects clusters with a clusterset selector
func (c *submarinerAgentController) selectClusters(selector clusterv1alpha1.ClusterSelector) (clusterNames []string, err error) {
	switch {
	case len(selector.ClusterNames) > 0 && selector.LabelSelector != nil:
		// return error if both ClusterNames and LabelSelector is specified for they are mutually exclusive
		// This case should be handled by validating webhook
		return nil, fmt.Errorf("both ClusterNames and LabelSelector is specified in ClusterSelector: %v", selector.LabelSelector)
	case len(selector.ClusterNames) > 0:
		// select clusters with cluster names
		for _, clusterName := range selector.ClusterNames {
			_, err = c.clusterLister.Get(clusterName)
			switch {
			case errors.IsNotFound(err):
				continue
			case err != nil:
				return nil, fmt.Errorf("unable to fetch ManagedCluster %q: %w", clusterName, err)
			default:
				clusterNames = append(clusterNames, clusterName)
			}
		}
		return clusterNames, nil
	case selector.LabelSelector != nil:
		// select clusters with label selector
		labelSelector, err := convertLabels(selector.LabelSelector)
		if err != nil {
			// This case should be handled by validating webhook
			return nil, fmt.Errorf("invalid label selector: %v, %w", selector.LabelSelector, err)
		}
		clusters, err := c.clusterLister.List(labelSelector)
		if err != nil {
			return nil, fmt.Errorf("unable to list ManagedClusters with label selector: %v, %w", selector.LabelSelector, err)
		}
		for _, cluster := range clusters {
			clusterNames = append(clusterNames, cluster.Name)
		}
		return clusterNames, nil
	default:
		// no cluster selected if neither ClusterNames nor LabelSelector is specified
		return clusterNames, nil
	}
}

// convertLabels returns label
func convertLabels(labelSelector *metav1.LabelSelector) (labels.Selector, error) {
	if labelSelector != nil {
		selector, err := metav1.LabelSelectorAsSelector(labelSelector)
		if err != nil {
			return labels.Nothing(), err
		}

		return selector, nil
	}

	return labels.Everything(), nil
}
