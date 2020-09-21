package submarineragent

import (
	"context"
	"fmt"
	"github.com/ghodss/yaml"
	workv1client "github.com/open-cluster-management/api/client/work/clientset/versioned"
	workv1 "github.com/open-cluster-management/api/work/v1"
	"github.com/openshift/library-go/pkg/assets"
	operatorhelpers "github.com/openshift/library-go/pkg/operator/v1helpers"
	"github.com/qiujian16/acm-submariner/pkg/helpers"
	hubbindata "github.com/qiujian16/acm-submariner/pkg/hub/bindata"
	"github.com/qiujian16/acm-submariner/pkg/hub/submarineragent/bindata"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog/v2"
	"path/filepath"
)

type wrapperInfo struct {
	name      string
	files     []string
	mustAsset func(name string) []byte
}

var (
	wrappers = []wrapperInfo{
		{
			name:      "submariner-agent-crds",
			mustAsset: hubbindata.MustAsset,
			files: []string{
				"manifests/crds/lighthouse.submariner.io_multiclusterservices_crd.yaml",
				"manifests/crds/lighthouse.submariner.io_serviceexports_crd.yaml",
				"manifests/crds/lighthouse.submariner.io_serviceimports_crd.yaml",
				"manifests/crds/submariner.io_clusters_crd.yaml",
				"manifests/crds/submariner.io_endpoints_crd.yaml",
				"manifests/crds/submariner.io_gateways_crd.yaml",
				"manifests/crds/submariner.io_servicediscoveries_crd.yaml",
				"manifests/crds/submariner.io_submariners_crd.yaml",
			},
		},
		{
			name:      "submariner-agent-rbac",
			mustAsset: bindata.MustAsset,
			files: []string{
				"manifests/agent/rbac/submariner-admin-aggeragate-clusterrole.yaml",
				"manifests/agent/rbac/submariner-lighthouse-clusterrole.yaml",
				"manifests/agent/rbac/submariner-lighthouse-clusterrolebinding.yaml",
				"manifests/agent/rbac/submariner-lighthouse-serviceaccount.yaml",
				"manifests/agent/rbac/submariner-operator-clusterrole.yaml",
				"manifests/agent/rbac/submariner-operator-clusterrolebinding.yaml",
				"manifests/agent/rbac/submariner-operator-namespace.yaml",
				"manifests/agent/rbac/submariner-operator-role.yaml",
				"manifests/agent/rbac/submariner-operator-rolebinding.yaml",
				"manifests/agent/rbac/submariner-operator-serviceaccount.yaml",
			},
		},
		{
			name:      "submariner-agent-operator",
			mustAsset: bindata.MustAsset,
			files: []string{
				"manifests/agent/operator/submariner-operator-deployment.yaml",
				"manifests/agent/operator/submariner.io-servicediscoveries-cr.yaml",
				"manifests/agent/operator/submariner.io-submariners-cr.yaml",
			},
		},
	}
)

type SubmarinerConfig struct {
	BrokerAPIServer string
	BrokerNamespace string
	BrokerToken     string
	BrokerCA        string
	IPSecPSK        string
	ClusterName     string
	ClusterCIDR     string
	ServiceCIDR     string
}

func NewSubmarinerConfig(
	client kubernetes.Interface,
	dynamicClient dynamic.Interface,
	clusterName, brokeNamespace string) (*SubmarinerConfig, error) {
	config := &SubmarinerConfig{
		BrokerNamespace: brokeNamespace,
		ClusterName:     clusterName,
	}

	apiServer, err := helpers.GetBrokerAPIServer(dynamicClient)
	if err != nil {
		return config, err
	}
	config.BrokerAPIServer = apiServer

	IPSecPSK, err := helpers.GetIPSecPSK(client, brokeNamespace)
	if err != nil {
		return config, err
	}
	config.IPSecPSK = IPSecPSK

	token, ca, err := helpers.GetBrokerTokenAndCA(client, brokeNamespace, clusterName)
	if err != nil {
		return config, err
	}
	config.BrokerCA = ca
	config.BrokerToken = token

	return config, nil
}

func wrapManifestWorks(config *SubmarinerConfig) ([]*workv1.ManifestWork, error) {
	var manifestWorks []*workv1.ManifestWork
	klog.V(4).Infof("config: %+v", config)
	for _, w := range wrappers {
		work := &workv1.ManifestWork{
			TypeMeta: metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{
				Name:      w.name,
				Namespace: config.ClusterName,
			},
			Spec: workv1.ManifestWorkSpec{},
		}

		var manifests []workv1.Manifest
		for _, file := range w.files {
			yamlData := assets.MustCreateAssetFromTemplate(file, w.mustAsset(filepath.Join("", file)), config).Data
			jsonData, err := yaml.YAMLToJSON(yamlData)
			if err != nil {
				klog.Errorf("failed to YAMLToJSON %+v: %+v : %+v", file, jsonData, err)
				return manifestWorks, err
			}
			manifest := workv1.Manifest{RawExtension: runtime.RawExtension{Raw: jsonData}}
			manifests = append(manifests, manifest)
		}
		work.Spec.Workload.Manifests = manifests

		manifestWorks = append(manifestWorks, work)
	}

	return manifestWorks, nil
}

func ApplySubmarinerManifestWorks(config *SubmarinerConfig, client workv1client.Interface, ctx context.Context) error {
	var errs []error

	manifestWorks, err := wrapManifestWorks(config)
	if err != nil {
		return fmt.Errorf("failed to wrap mainfestWorks:%+v", err)
	}

	for _, work := range manifestWorks {
		err := ApplyManifestWork(work, client, ctx)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to apply manifestWork %+v: %+v", work.Name, err))
			klog.Errorf("failed to apply manifestWork %+v: %+v : %+v", work.Name, work, err)
			continue
		}
	}

	return operatorhelpers.NewMultiLineAggregate(errs)
}

func RemoveSubmarinerManifestWorks(namespace string, client workv1client.Interface, ctx context.Context) error {
	var errs []error
	for _, w := range wrappers {
		if err := client.WorkV1().ManifestWorks(namespace).
			Delete(ctx, w.name, metav1.DeleteOptions{}); err != nil {
			if errors.IsNotFound(err) {
				continue
			}
			errs = append(errs, err)
		}
	}
	return operatorhelpers.NewMultiLineAggregate(errs)
}

func ApplyManifestWork(required *workv1.ManifestWork, client workv1client.Interface, ctx context.Context) error {
	err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		existing, err := client.WorkV1().ManifestWorks(required.Namespace).
			Get(ctx, required.Name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				_, err = client.WorkV1().ManifestWorks(required.Namespace).
					Create(ctx, required, metav1.CreateOptions{})
			}
			return err
		}

		if !equality.Semantic.DeepEqual(existing.Spec, required.Spec) {
			_, err = client.WorkV1().ManifestWorks(required.Namespace).
				Update(ctx, required, metav1.UpdateOptions{})
		}

		return err
	})

	return err
}
