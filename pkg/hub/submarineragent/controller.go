package submarineragent

import (
	"context"

	clusterclientset "github.com/open-cluster-management/api/client/cluster/clientset/versioned/typed/cluster/v1"
	clusterinformerv1 "github.com/open-cluster-management/api/client/cluster/informers/externalversions/cluster/v1"
	clusterlisterv1 "github.com/open-cluster-management/api/client/cluster/listers/cluster/v1"
	workv1client "github.com/open-cluster-management/api/client/work/clientset/versioned"
	workinformer "github.com/open-cluster-management/api/client/work/informers/externalversions/work/v1"
	worklister "github.com/open-cluster-management/api/client/work/listers/work/v1"
	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/events"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
)

type submarinerAgentController struct {
	kubeClient         kubernetes.Interface
	clusterClient      clusterclientset.ManagedClusterInterface
	manifestWorkClient workv1client.Interface
	manifestWorkLister worklister.ManifestWorkLister
	clusterLister      clusterlisterv1.ManagedClusterLister
	eventRecorder      events.Recorder
}

func NewSubmarinerAgentController(
	clusterClient clusterclientset.ManagedClusterInterface,
	kubeClient kubernetes.Interface,
	manifestWorkClient workv1client.Interface,
	clusterInformer clusterinformerv1.ManagedClusterInformer,
	manifestWorkInformer workinformer.ManifestWorkInformer,
	recorder events.Recorder) factory.Controller {
	c := &submarinerAgentController{
		kubeClient:         kubeClient,
		manifestWorkClient: manifestWorkClient,
		manifestWorkLister: manifestWorkInformer.Lister(),
		clusterClient:      clusterClient,
		clusterLister:      clusterInformer.Lister(),
		eventRecorder:      recorder.WithComponentSuffix("submariner-agent-controller"),
	}
	return factory.New().
		WithInformersQueueKeyFunc(func(obj runtime.Object) string {
			accessor, _ := meta.Accessor(obj)
			return accessor.GetName()
		}, clusterInformer.Informer()).
		WithSync(c.sync).
		ToController("SubmarinerAgentController", recorder)
}

func (c *submarinerAgentController) sync(ctx context.Context, syncCtx factory.SyncContext) error {
	// get clusterset
	// get service account token/rolebinding
	// generate submariner CR
	// install operator manifestwork
	// install submarinere CR manifestwork
	return nil
}
