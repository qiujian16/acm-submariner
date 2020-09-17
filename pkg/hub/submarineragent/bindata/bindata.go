// Code generated for package bindata by go-bindata DO NOT EDIT. (@generated)
// sources:
// manifests/agent/crds/lighthouse.submariner.io_multiclusterservices_crd.yaml
// manifests/agent/crds/lighthouse.submariner.io_serviceexports_crd.yaml
// manifests/agent/crds/lighthouse.submariner.io_serviceimports_crd.yaml
// manifests/agent/crds/submariner.io_clusters_crd.yaml
// manifests/agent/crds/submariner.io_endpoints_crd.yaml
// manifests/agent/crds/submariner.io_gateways_crd.yaml
// manifests/agent/crds/submariner.io_servicediscoveries_crd.yaml
// manifests/agent/crds/submariner.io_submariners_crd.yaml
// manifests/agent/operator/submariner-operator-deployment.yaml
// manifests/agent/operator/submariner.io-servicediscoveries-cr.yaml
// manifests/agent/operator/submariner.io-submariners-cr.yaml
// manifests/agent/rbac/submariner-lighthouse-clusterrole.yaml
// manifests/agent/rbac/submariner-lighthouse-clusterrolebinding.yaml
// manifests/agent/rbac/submariner-lighthouse-serviceaccount.yaml
// manifests/agent/rbac/submariner-operator-clusterrole.yaml
// manifests/agent/rbac/submariner-operator-clusterrolebinding.yaml
// manifests/agent/rbac/submariner-operator-namespace.yaml
// manifests/agent/rbac/submariner-operator-role.yaml
// manifests/agent/rbac/submariner-operator-rolebinding.yaml
// manifests/agent/rbac/submariner-operator-serviceaccount.yaml
package bindata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _manifestsAgentCrdsLighthouseSubmarinerIo_multiclusterservices_crdYaml = []byte(`---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: multiclusterservices.lighthouse.submariner.io
spec:
  group: lighthouse.submariner.io
  version: v1
  names:
    kind: MultiClusterService
    plural: multiclusterservices
    singular: multiclusterservice
  scope: Namespaced
  validation:
    openAPIV3Schema:
      properties:
        spec:
          properties:
            clusterServiceInfo:
              properties:
                clusterID:
                  type: "string"
                  minimum: 1
                clusterDomain:
                  type: "string"
                  minimum: 0
                serviceIP:
                  type: "string"
                  minimum: 1
                port:
                  type: "integer"
                  minimum: 0
`)

func manifestsAgentCrdsLighthouseSubmarinerIo_multiclusterservices_crdYamlBytes() ([]byte, error) {
	return _manifestsAgentCrdsLighthouseSubmarinerIo_multiclusterservices_crdYaml, nil
}

func manifestsAgentCrdsLighthouseSubmarinerIo_multiclusterservices_crdYaml() (*asset, error) {
	bytes, err := manifestsAgentCrdsLighthouseSubmarinerIo_multiclusterservices_crdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/crds/lighthouse.submariner.io_multiclusterservices_crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentCrdsLighthouseSubmarinerIo_serviceexports_crdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: serviceexports.lighthouse.submariner.io
spec:
  group: lighthouse.submariner.io
  version: v2alpha1
  names:
    kind: ServiceExport
    plural: serviceexports
    singular: serviceexport
  scope: Namespaced
`)

func manifestsAgentCrdsLighthouseSubmarinerIo_serviceexports_crdYamlBytes() ([]byte, error) {
	return _manifestsAgentCrdsLighthouseSubmarinerIo_serviceexports_crdYaml, nil
}

func manifestsAgentCrdsLighthouseSubmarinerIo_serviceexports_crdYaml() (*asset, error) {
	bytes, err := manifestsAgentCrdsLighthouseSubmarinerIo_serviceexports_crdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/crds/lighthouse.submariner.io_serviceexports_crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentCrdsLighthouseSubmarinerIo_serviceimports_crdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: serviceimports.lighthouse.submariner.io
spec:
  group: lighthouse.submariner.io
  version: v2alpha1
  names:
    kind: ServiceImport
    plural: serviceimports
    singular: serviceimport
  scope: Namespaced
`)

func manifestsAgentCrdsLighthouseSubmarinerIo_serviceimports_crdYamlBytes() ([]byte, error) {
	return _manifestsAgentCrdsLighthouseSubmarinerIo_serviceimports_crdYaml, nil
}

func manifestsAgentCrdsLighthouseSubmarinerIo_serviceimports_crdYaml() (*asset, error) {
	bytes, err := manifestsAgentCrdsLighthouseSubmarinerIo_serviceimports_crdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/crds/lighthouse.submariner.io_serviceimports_crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentCrdsSubmarinerIo_clusters_crdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: clusters.submariner.io
spec:
  group: submariner.io
  version: v1
  names:
    kind: Cluster
    listKind: ClusterList
    plural: clusters
    singular: cluster
  scope: Namespaced
`)

func manifestsAgentCrdsSubmarinerIo_clusters_crdYamlBytes() ([]byte, error) {
	return _manifestsAgentCrdsSubmarinerIo_clusters_crdYaml, nil
}

func manifestsAgentCrdsSubmarinerIo_clusters_crdYaml() (*asset, error) {
	bytes, err := manifestsAgentCrdsSubmarinerIo_clusters_crdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/crds/submariner.io_clusters_crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentCrdsSubmarinerIo_endpoints_crdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: endpoints.submariner.io
spec:
  group: submariner.io
  version: v1
  names:
    kind: Endpoint
    listKind: EndpointList
    plural: endpoints
    singular: endpoint
  scope: Namespaced
`)

func manifestsAgentCrdsSubmarinerIo_endpoints_crdYamlBytes() ([]byte, error) {
	return _manifestsAgentCrdsSubmarinerIo_endpoints_crdYaml, nil
}

func manifestsAgentCrdsSubmarinerIo_endpoints_crdYaml() (*asset, error) {
	bytes, err := manifestsAgentCrdsSubmarinerIo_endpoints_crdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/crds/submariner.io_endpoints_crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentCrdsSubmarinerIo_gateways_crdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: gateways.submariner.io
spec:
  group: submariner.io
  version: v1
  names:
    kind: Gateway
    listKind: GatewayList
    plural: gateways
    singular: gateway
  scope: Namespaced
  additionalPrinterColumns:
    - name: ha-status
      type: string
      description: High Availability Status of the Gateway
      JSONPath: .status.haStatus
`)

func manifestsAgentCrdsSubmarinerIo_gateways_crdYamlBytes() ([]byte, error) {
	return _manifestsAgentCrdsSubmarinerIo_gateways_crdYaml, nil
}

func manifestsAgentCrdsSubmarinerIo_gateways_crdYaml() (*asset, error) {
	bytes, err := manifestsAgentCrdsSubmarinerIo_gateways_crdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/crds/submariner.io_gateways_crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentCrdsSubmarinerIo_servicediscoveries_crdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: servicediscoveries.submariner.io
spec:
  group: submariner.io
  names:
    kind: ServiceDiscovery
    listKind: ServiceDiscoveryList
    plural: servicediscoveries
    singular: servicediscovery
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: ServiceDiscovery is the Schema for the servicediscoveries API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: ServiceDiscoverySpec defines the desired state of ServiceDiscovery
          properties:
            brokerK8sApiServer:
              type: string
            brokerK8sApiServerToken:
              type: string
            brokerK8sCA:
              type: string
            brokerK8sRemoteNamespace:
              type: string
            clusterID:
              type: string
            debug:
              type: boolean
            globalnetEnabled:
              type: boolean
            namespace:
              type: string
            repository:
              type: string
            version:
              type: string
          required:
          - brokerK8sApiServer
          - brokerK8sApiServerToken
          - brokerK8sCA
          - brokerK8sRemoteNamespace
          - clusterID
          - debug
          - namespace
          type: object
        status:
          description: ServiceDiscoveryStatus defines the observed state of ServiceDiscovery
          type: object
      type: object
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
`)

func manifestsAgentCrdsSubmarinerIo_servicediscoveries_crdYamlBytes() ([]byte, error) {
	return _manifestsAgentCrdsSubmarinerIo_servicediscoveries_crdYaml, nil
}

func manifestsAgentCrdsSubmarinerIo_servicediscoveries_crdYaml() (*asset, error) {
	bytes, err := manifestsAgentCrdsSubmarinerIo_servicediscoveries_crdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/crds/submariner.io_servicediscoveries_crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentCrdsSubmarinerIo_submariners_crdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: submariners.submariner.io
spec:
  group: submariner.io
  names:
    kind: Submariner
    listKind: SubmarinerList
    plural: submariners
    singular: submariner
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: Submariner is the Schema for the submariners API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: SubmarinerSpec defines the desired state of Submariner
          properties:
            broker:
              type: string
            brokerK8sApiServer:
              type: string
            brokerK8sApiServerToken:
              type: string
            brokerK8sCA:
              type: string
            brokerK8sRemoteNamespace:
              type: string
            cableDriver:
              type: string
            ceIPSecDebug:
              type: boolean
            ceIPSecIKEPort:
              type: integer
            ceIPSecNATTPort:
              type: integer
            ceIPSecPSK:
              type: string
            clusterCIDR:
              type: string
            clusterID:
              type: string
            colorCodes:
              type: string
            debug:
              type: boolean
            globalCIDR:
              type: string
            namespace:
              type: string
            natEnabled:
              type: boolean
            repository:
              type: string
            serviceCIDR:
              type: string
            serviceDiscoveryEnabled:
              type: boolean
            version:
              type: string
          required:
          - broker
          - brokerK8sApiServer
          - brokerK8sApiServerToken
          - brokerK8sCA
          - brokerK8sRemoteNamespace
          - ceIPSecDebug
          - ceIPSecPSK
          - clusterCIDR
          - clusterID
          - debug
          - namespace
          - natEnabled
          - serviceCIDR
          type: object
        status:
          description: SubmarinerStatus defines the observed state of Submariner
          properties:
            clusterCIDR:
              type: string
            clusterID:
              type: string
            colorCodes:
              type: string
            engineDaemonSetStatus:
              description: DaemonSetStatus represents the current status of a daemon
                set.
              properties:
                collisionCount:
                  description: Count of hash collisions for the DaemonSet. The DaemonSet
                    controller uses this field as a collision avoidance mechanism
                    when it needs to create the name for the newest ControllerRevision.
                  format: int32
                  type: integer
                conditions:
                  description: Represents the latest available observations of a DaemonSet's
                    current state.
                  items:
                    description: DaemonSetCondition describes the state of a DaemonSet
                      at a certain point.
                    properties:
                      lastTransitionTime:
                        description: Last time the condition transitioned from one
                          status to another.
                        format: date-time
                        type: string
                      message:
                        description: A human readable message indicating details about
                          the transition.
                        type: string
                      reason:
                        description: The reason for the condition's last transition.
                        type: string
                      status:
                        description: Status of the condition, one of True, False,
                          Unknown.
                        type: string
                      type:
                        description: Type of DaemonSet condition.
                        type: string
                    required:
                    - status
                    - type
                    type: object
                  type: array
                currentNumberScheduled:
                  description: 'The number of nodes that are running at least 1 daemon
                    pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/'
                  format: int32
                  type: integer
                desiredNumberScheduled:
                  description: 'The total number of nodes that should be running the
                    daemon pod (including nodes correctly running the daemon pod).
                    More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/'
                  format: int32
                  type: integer
                numberAvailable:
                  description: The number of nodes that should be running the daemon
                    pod and have one or more of the daemon pod running and available
                    (ready for at least spec.minReadySeconds)
                  format: int32
                  type: integer
                numberMisscheduled:
                  description: 'The number of nodes that are running the daemon pod,
                    but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/'
                  format: int32
                  type: integer
                numberReady:
                  description: The number of nodes that should be running the daemon
                    pod and have one or more of the daemon pod running and ready.
                  format: int32
                  type: integer
                numberUnavailable:
                  description: The number of nodes that should be running the daemon
                    pod and have none of the daemon pod running and available (ready
                    for at least spec.minReadySeconds)
                  format: int32
                  type: integer
                observedGeneration:
                  description: The most recent generation observed by the daemon set
                    controller.
                  format: int64
                  type: integer
                updatedNumberScheduled:
                  description: The total number of nodes that are running updated
                    daemon pod
                  format: int32
                  type: integer
              required:
              - currentNumberScheduled
              - desiredNumberScheduled
              - numberMisscheduled
              - numberReady
              type: object
            gateways:
              items:
                properties:
                  connections:
                    items:
                      properties:
                        endpoint:
                          properties:
                            backend:
                              type: string
                            backend_config:
                              additionalProperties:
                                type: string
                              type: object
                            cable_name:
                              type: string
                            cluster_id:
                              type: string
                            hostname:
                              type: string
                            nat_enabled:
                              type: boolean
                            private_ip:
                              type: string
                            public_ip:
                              type: string
                            subnets:
                              items:
                                type: string
                              type: array
                          required:
                          - backend
                          - cable_name
                          - cluster_id
                          - hostname
                          - nat_enabled
                          - private_ip
                          - public_ip
                          - subnets
                          type: object
                        status:
                          type: string
                        statusMessage:
                          type: string
                      required:
                      - endpoint
                      - status
                      - statusMessage
                      type: object
                    type: array
                  haStatus:
                    type: string
                  localEndpoint:
                    properties:
                      backend:
                        type: string
                      backend_config:
                        additionalProperties:
                          type: string
                        type: object
                      cable_name:
                        type: string
                      cluster_id:
                        type: string
                      hostname:
                        type: string
                      nat_enabled:
                        type: boolean
                      private_ip:
                        type: string
                      public_ip:
                        type: string
                      subnets:
                        items:
                          type: string
                        type: array
                    required:
                    - backend
                    - cable_name
                    - cluster_id
                    - hostname
                    - nat_enabled
                    - private_ip
                    - public_ip
                    - subnets
                    type: object
                  statusFailure:
                    type: string
                  version:
                    type: string
                required:
                - connections
                - haStatus
                - localEndpoint
                - statusFailure
                - version
                type: object
              type: array
            globalCIDR:
              type: string
            globalnetDaemonSetStatus:
              description: DaemonSetStatus represents the current status of a daemon
                set.
              properties:
                collisionCount:
                  description: Count of hash collisions for the DaemonSet. The DaemonSet
                    controller uses this field as a collision avoidance mechanism
                    when it needs to create the name for the newest ControllerRevision.
                  format: int32
                  type: integer
                conditions:
                  description: Represents the latest available observations of a DaemonSet's
                    current state.
                  items:
                    description: DaemonSetCondition describes the state of a DaemonSet
                      at a certain point.
                    properties:
                      lastTransitionTime:
                        description: Last time the condition transitioned from one
                          status to another.
                        format: date-time
                        type: string
                      message:
                        description: A human readable message indicating details about
                          the transition.
                        type: string
                      reason:
                        description: The reason for the condition's last transition.
                        type: string
                      status:
                        description: Status of the condition, one of True, False,
                          Unknown.
                        type: string
                      type:
                        description: Type of DaemonSet condition.
                        type: string
                    required:
                    - status
                    - type
                    type: object
                  type: array
                currentNumberScheduled:
                  description: 'The number of nodes that are running at least 1 daemon
                    pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/'
                  format: int32
                  type: integer
                desiredNumberScheduled:
                  description: 'The total number of nodes that should be running the
                    daemon pod (including nodes correctly running the daemon pod).
                    More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/'
                  format: int32
                  type: integer
                numberAvailable:
                  description: The number of nodes that should be running the daemon
                    pod and have one or more of the daemon pod running and available
                    (ready for at least spec.minReadySeconds)
                  format: int32
                  type: integer
                numberMisscheduled:
                  description: 'The number of nodes that are running the daemon pod,
                    but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/'
                  format: int32
                  type: integer
                numberReady:
                  description: The number of nodes that should be running the daemon
                    pod and have one or more of the daemon pod running and ready.
                  format: int32
                  type: integer
                numberUnavailable:
                  description: The number of nodes that should be running the daemon
                    pod and have none of the daemon pod running and available (ready
                    for at least spec.minReadySeconds)
                  format: int32
                  type: integer
                observedGeneration:
                  description: The most recent generation observed by the daemon set
                    controller.
                  format: int64
                  type: integer
                updatedNumberScheduled:
                  description: The total number of nodes that are running updated
                    daemon pod
                  format: int32
                  type: integer
              required:
              - currentNumberScheduled
              - desiredNumberScheduled
              - numberMisscheduled
              - numberReady
              type: object
            natEnabled:
              type: boolean
            routeAgentDaemonSetStatus:
              description: DaemonSetStatus represents the current status of a daemon
                set.
              properties:
                collisionCount:
                  description: Count of hash collisions for the DaemonSet. The DaemonSet
                    controller uses this field as a collision avoidance mechanism
                    when it needs to create the name for the newest ControllerRevision.
                  format: int32
                  type: integer
                conditions:
                  description: Represents the latest available observations of a DaemonSet's
                    current state.
                  items:
                    description: DaemonSetCondition describes the state of a DaemonSet
                      at a certain point.
                    properties:
                      lastTransitionTime:
                        description: Last time the condition transitioned from one
                          status to another.
                        format: date-time
                        type: string
                      message:
                        description: A human readable message indicating details about
                          the transition.
                        type: string
                      reason:
                        description: The reason for the condition's last transition.
                        type: string
                      status:
                        description: Status of the condition, one of True, False,
                          Unknown.
                        type: string
                      type:
                        description: Type of DaemonSet condition.
                        type: string
                    required:
                    - status
                    - type
                    type: object
                  type: array
                currentNumberScheduled:
                  description: 'The number of nodes that are running at least 1 daemon
                    pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/'
                  format: int32
                  type: integer
                desiredNumberScheduled:
                  description: 'The total number of nodes that should be running the
                    daemon pod (including nodes correctly running the daemon pod).
                    More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/'
                  format: int32
                  type: integer
                numberAvailable:
                  description: The number of nodes that should be running the daemon
                    pod and have one or more of the daemon pod running and available
                    (ready for at least spec.minReadySeconds)
                  format: int32
                  type: integer
                numberMisscheduled:
                  description: 'The number of nodes that are running the daemon pod,
                    but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/'
                  format: int32
                  type: integer
                numberReady:
                  description: The number of nodes that should be running the daemon
                    pod and have one or more of the daemon pod running and ready.
                  format: int32
                  type: integer
                numberUnavailable:
                  description: The number of nodes that should be running the daemon
                    pod and have none of the daemon pod running and available (ready
                    for at least spec.minReadySeconds)
                  format: int32
                  type: integer
                observedGeneration:
                  description: The most recent generation observed by the daemon set
                    controller.
                  format: int64
                  type: integer
                updatedNumberScheduled:
                  description: The total number of nodes that are running updated
                    daemon pod
                  format: int32
                  type: integer
              required:
              - currentNumberScheduled
              - desiredNumberScheduled
              - numberMisscheduled
              - numberReady
              type: object
            serviceCIDR:
              type: string
          required:
          - clusterID
          - natEnabled
          type: object
      type: object
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
`)

func manifestsAgentCrdsSubmarinerIo_submariners_crdYamlBytes() ([]byte, error) {
	return _manifestsAgentCrdsSubmarinerIo_submariners_crdYaml, nil
}

func manifestsAgentCrdsSubmarinerIo_submariners_crdYaml() (*asset, error) {
	bytes, err := manifestsAgentCrdsSubmarinerIo_submariners_crdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/crds/submariner.io_submariners_crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentOperatorSubmarinerOperatorDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: submariner-operator
  namespace: submariner-operator
spec:
  selector:
    matchLabels:
      name: submariner-operator
  template:
    metadata:
      labels:
        name: submariner-operator
    spec:
      containers:
      - command:
        - submariner-operator
        env:
        - name: WATCH_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: POD_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.name
        - name: OPERATOR_NAME
          value: submariner-operator
        image: quay.io/submariner/submariner-operator:0.6.1
        imagePullPolicy: Always
        name: submariner-operator
      serviceAccount: submariner-operator
      serviceAccountName: submariner-operator
`)

func manifestsAgentOperatorSubmarinerOperatorDeploymentYamlBytes() ([]byte, error) {
	return _manifestsAgentOperatorSubmarinerOperatorDeploymentYaml, nil
}

func manifestsAgentOperatorSubmarinerOperatorDeploymentYaml() (*asset, error) {
	bytes, err := manifestsAgentOperatorSubmarinerOperatorDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/operator/submariner-operator-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentOperatorSubmarinerIoServicediscoveriesCrYaml = []byte(`apiVersion: submariner.io/v1alpha1
kind: ServiceDiscovery
metadata:
  name: service-discovery
  namespace: submariner-operator
spec:
  brokerK8sApiServer: {{ .BrokerAPIServer }}
  brokerK8sApiServerToken: {{ .BrokerToken }}
  brokerK8sCA: {{ .BrokerCA }}
  brokerK8sRemoteNamespace: submariner-k8s-broker
  clusterID: {{ .ManagedClusterName }}
  debug: false
  namespace: submariner-operator
  repository: quay.io/submariner
  version: 0.6.1
`)

func manifestsAgentOperatorSubmarinerIoServicediscoveriesCrYamlBytes() ([]byte, error) {
	return _manifestsAgentOperatorSubmarinerIoServicediscoveriesCrYaml, nil
}

func manifestsAgentOperatorSubmarinerIoServicediscoveriesCrYaml() (*asset, error) {
	bytes, err := manifestsAgentOperatorSubmarinerIoServicediscoveriesCrYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/operator/submariner.io-servicediscoveries-cr.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentOperatorSubmarinerIoSubmarinersCrYaml = []byte(`apiVersion: submariner.io/v1alpha1
kind: Submariner
metadata:
  name: submariner
  namespace: submariner-operator
spec:
  broker: k8s
  brokerK8sApiServer: {{ .BrokerAPIServer }}
  brokerK8sApiServerToken: {{ .BrokerToken }}
  brokerK8sCA: {{ .BrokerCA }}
  brokerK8sRemoteNamespace: submariner-k8s-broker
  ceIPSecDebug: false
  ceIPSecIKEPort: 500
  ceIPSecNATTPort: 4500
  ceIPSecPSK: {{ .IPSecPSK }}
  clusterCIDR: ""
  clusterID: {{ .ManagedClusterName }}
  colorCodes: blue
  debug: false
  namespace: submariner-operator
  natEnabled: true
  repository: quay.io/submariner
  serviceCIDR: ""
  serviceDiscoveryEnabled: true
  version: 0.6.1
`)

func manifestsAgentOperatorSubmarinerIoSubmarinersCrYamlBytes() ([]byte, error) {
	return _manifestsAgentOperatorSubmarinerIoSubmarinersCrYaml, nil
}

func manifestsAgentOperatorSubmarinerIoSubmarinersCrYaml() (*asset, error) {
	bytes, err := manifestsAgentOperatorSubmarinerIoSubmarinersCrYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/operator/submariner.io-submariners-cr.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentRbacSubmarinerLighthouseClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: submariner-lighthouse
rules:
  - apiGroups:
      - ""
    resources:
      - services
      - namespaces
      - endpoints
    verbs:
      - get
      - list
      - watch
      - update
  - apiGroups:
      - discovery.k8s.io
    resources:
      - endpointslices
    verbs:
      - create
      - get
      - list
      - watch
      - update
      - delete
      - deletecollection
  - apiGroups:
      - lighthouse.submariner.io
    resources:
      - "*"
    verbs:
      - create
      - get
      - list
      - watch
      - update
      - delete
  - apiGroups:
      - submariner.io
    resources:
      - "gateways"
    verbs:
      - get
      - list
      - watch
`)

func manifestsAgentRbacSubmarinerLighthouseClusterroleYamlBytes() ([]byte, error) {
	return _manifestsAgentRbacSubmarinerLighthouseClusterroleYaml, nil
}

func manifestsAgentRbacSubmarinerLighthouseClusterroleYaml() (*asset, error) {
	bytes, err := manifestsAgentRbacSubmarinerLighthouseClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/rbac/submariner-lighthouse-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentRbacSubmarinerLighthouseClusterrolebindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: submariner-lighthouse
subjects:
  - kind: ServiceAccount
    name: submariner-lighthouse
    namespace: submariner-operator
roleRef:
  kind: ClusterRole
  name: submariner-lighthouse
  apiGroup: rbac.authorization.k8s.io
`)

func manifestsAgentRbacSubmarinerLighthouseClusterrolebindingYamlBytes() ([]byte, error) {
	return _manifestsAgentRbacSubmarinerLighthouseClusterrolebindingYaml, nil
}

func manifestsAgentRbacSubmarinerLighthouseClusterrolebindingYaml() (*asset, error) {
	bytes, err := manifestsAgentRbacSubmarinerLighthouseClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/rbac/submariner-lighthouse-clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentRbacSubmarinerLighthouseServiceaccountYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: submariner-lighthouse
  namespace: submariner-operator
`)

func manifestsAgentRbacSubmarinerLighthouseServiceaccountYamlBytes() ([]byte, error) {
	return _manifestsAgentRbacSubmarinerLighthouseServiceaccountYaml, nil
}

func manifestsAgentRbacSubmarinerLighthouseServiceaccountYaml() (*asset, error) {
	bytes, err := manifestsAgentRbacSubmarinerLighthouseServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/rbac/submariner-lighthouse-serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentRbacSubmarinerOperatorClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: submariner-operator
rules:
  - apiGroups:
      - ""
    resources:
      - configmaps
    verbs:
      - get
      - list
      - watch
      - update
  - apiGroups:
      - ""
    resources:
      - pods
      - services
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - operator.openshift.io
    resources:
      - dnses
    verbs:
      - get
      - list
      - watch
      - update
  - apiGroups:
      - config.openshift.io
    resources:
      - networks
    verbs:
      - get
      - list
  - apiGroups:
      - network.openshift.io
    resources:
      - clusternetworks
    verbs:
      - get
      - list
  - apiGroups:
      - ""
    resources:
      - nodes
    verbs:
      - get
      - list
      - watch
      - update
`)

func manifestsAgentRbacSubmarinerOperatorClusterroleYamlBytes() ([]byte, error) {
	return _manifestsAgentRbacSubmarinerOperatorClusterroleYaml, nil
}

func manifestsAgentRbacSubmarinerOperatorClusterroleYaml() (*asset, error) {
	bytes, err := manifestsAgentRbacSubmarinerOperatorClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/rbac/submariner-operator-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentRbacSubmarinerOperatorClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: submariner-operator
subjects:
  - kind: ServiceAccount
    name: submariner-operator
    namespace: submariner-operator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: submariner-operator
`)

func manifestsAgentRbacSubmarinerOperatorClusterrolebindingYamlBytes() ([]byte, error) {
	return _manifestsAgentRbacSubmarinerOperatorClusterrolebindingYaml, nil
}

func manifestsAgentRbacSubmarinerOperatorClusterrolebindingYaml() (*asset, error) {
	bytes, err := manifestsAgentRbacSubmarinerOperatorClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/rbac/submariner-operator-clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentRbacSubmarinerOperatorNamespaceYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: submariner-operator
`)

func manifestsAgentRbacSubmarinerOperatorNamespaceYamlBytes() ([]byte, error) {
	return _manifestsAgentRbacSubmarinerOperatorNamespaceYaml, nil
}

func manifestsAgentRbacSubmarinerOperatorNamespaceYaml() (*asset, error) {
	bytes, err := manifestsAgentRbacSubmarinerOperatorNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/rbac/submariner-operator-namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentRbacSubmarinerOperatorRoleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: submariner-operator
  namespace: submariner-operator
rules:
  - apiGroups:
      - ""
    resources:
      - pods
      - services
      - services/finalizers
      - endpoints
      - persistentvolumeclaims
      - events
      - configmaps
      - secrets
    verbs:
      - '*'
  - apiGroups:
      - apps
    resources:
      - deployments
      - daemonsets
      - replicasets
      - statefulsets
    verbs:
      - '*'
  - apiGroups:
      - monitoring.coreos.com
    resources:
      - servicemonitors
    verbs:
      - get
      - create
  - apiGroups:
      - apps
    resourceNames:
      - submariner-operator
    resources:
      - deployments/finalizers
    verbs:
      - update
  - apiGroups:
      - ""
    resources:
      - pods
    verbs:
      - get
  - apiGroups:
      - apps
    resources:
      - replicasets
    verbs:
      - get
  - apiGroups:
      - submariner.io
    resources:
      - '*'
      - servicediscoveries
    verbs:
      - '*'
`)

func manifestsAgentRbacSubmarinerOperatorRoleYamlBytes() ([]byte, error) {
	return _manifestsAgentRbacSubmarinerOperatorRoleYaml, nil
}

func manifestsAgentRbacSubmarinerOperatorRoleYaml() (*asset, error) {
	bytes, err := manifestsAgentRbacSubmarinerOperatorRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/rbac/submariner-operator-role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentRbacSubmarinerOperatorRolebindingYaml = []byte(`kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: submariner-operator
  namespace: submariner-operator
subjects:
  - kind: ServiceAccount
    name: submariner-operator
roleRef:
  kind: Role
  name: submariner-operator
  apiGroup: rbac.authorization.k8s.io
`)

func manifestsAgentRbacSubmarinerOperatorRolebindingYamlBytes() ([]byte, error) {
	return _manifestsAgentRbacSubmarinerOperatorRolebindingYaml, nil
}

func manifestsAgentRbacSubmarinerOperatorRolebindingYaml() (*asset, error) {
	bytes, err := manifestsAgentRbacSubmarinerOperatorRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/rbac/submariner-operator-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsAgentRbacSubmarinerOperatorServiceaccountYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: submariner-operator
  namespace: submariner-operator
`)

func manifestsAgentRbacSubmarinerOperatorServiceaccountYamlBytes() ([]byte, error) {
	return _manifestsAgentRbacSubmarinerOperatorServiceaccountYaml, nil
}

func manifestsAgentRbacSubmarinerOperatorServiceaccountYaml() (*asset, error) {
	bytes, err := manifestsAgentRbacSubmarinerOperatorServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/agent/rbac/submariner-operator-serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"manifests/agent/crds/lighthouse.submariner.io_multiclusterservices_crd.yaml": manifestsAgentCrdsLighthouseSubmarinerIo_multiclusterservices_crdYaml,
	"manifests/agent/crds/lighthouse.submariner.io_serviceexports_crd.yaml":       manifestsAgentCrdsLighthouseSubmarinerIo_serviceexports_crdYaml,
	"manifests/agent/crds/lighthouse.submariner.io_serviceimports_crd.yaml":       manifestsAgentCrdsLighthouseSubmarinerIo_serviceimports_crdYaml,
	"manifests/agent/crds/submariner.io_clusters_crd.yaml":                        manifestsAgentCrdsSubmarinerIo_clusters_crdYaml,
	"manifests/agent/crds/submariner.io_endpoints_crd.yaml":                       manifestsAgentCrdsSubmarinerIo_endpoints_crdYaml,
	"manifests/agent/crds/submariner.io_gateways_crd.yaml":                        manifestsAgentCrdsSubmarinerIo_gateways_crdYaml,
	"manifests/agent/crds/submariner.io_servicediscoveries_crd.yaml":              manifestsAgentCrdsSubmarinerIo_servicediscoveries_crdYaml,
	"manifests/agent/crds/submariner.io_submariners_crd.yaml":                     manifestsAgentCrdsSubmarinerIo_submariners_crdYaml,
	"manifests/agent/operator/submariner-operator-deployment.yaml":                manifestsAgentOperatorSubmarinerOperatorDeploymentYaml,
	"manifests/agent/operator/submariner.io-servicediscoveries-cr.yaml":           manifestsAgentOperatorSubmarinerIoServicediscoveriesCrYaml,
	"manifests/agent/operator/submariner.io-submariners-cr.yaml":                  manifestsAgentOperatorSubmarinerIoSubmarinersCrYaml,
	"manifests/agent/rbac/submariner-lighthouse-clusterrole.yaml":                 manifestsAgentRbacSubmarinerLighthouseClusterroleYaml,
	"manifests/agent/rbac/submariner-lighthouse-clusterrolebinding.yaml":          manifestsAgentRbacSubmarinerLighthouseClusterrolebindingYaml,
	"manifests/agent/rbac/submariner-lighthouse-serviceaccount.yaml":              manifestsAgentRbacSubmarinerLighthouseServiceaccountYaml,
	"manifests/agent/rbac/submariner-operator-clusterrole.yaml":                   manifestsAgentRbacSubmarinerOperatorClusterroleYaml,
	"manifests/agent/rbac/submariner-operator-clusterrolebinding.yaml":            manifestsAgentRbacSubmarinerOperatorClusterrolebindingYaml,
	"manifests/agent/rbac/submariner-operator-namespace.yaml":                     manifestsAgentRbacSubmarinerOperatorNamespaceYaml,
	"manifests/agent/rbac/submariner-operator-role.yaml":                          manifestsAgentRbacSubmarinerOperatorRoleYaml,
	"manifests/agent/rbac/submariner-operator-rolebinding.yaml":                   manifestsAgentRbacSubmarinerOperatorRolebindingYaml,
	"manifests/agent/rbac/submariner-operator-serviceaccount.yaml":                manifestsAgentRbacSubmarinerOperatorServiceaccountYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"manifests": {nil, map[string]*bintree{
		"agent": {nil, map[string]*bintree{
			"crds": {nil, map[string]*bintree{
				"lighthouse.submariner.io_multiclusterservices_crd.yaml": {manifestsAgentCrdsLighthouseSubmarinerIo_multiclusterservices_crdYaml, map[string]*bintree{}},
				"lighthouse.submariner.io_serviceexports_crd.yaml":       {manifestsAgentCrdsLighthouseSubmarinerIo_serviceexports_crdYaml, map[string]*bintree{}},
				"lighthouse.submariner.io_serviceimports_crd.yaml":       {manifestsAgentCrdsLighthouseSubmarinerIo_serviceimports_crdYaml, map[string]*bintree{}},
				"submariner.io_clusters_crd.yaml":                        {manifestsAgentCrdsSubmarinerIo_clusters_crdYaml, map[string]*bintree{}},
				"submariner.io_endpoints_crd.yaml":                       {manifestsAgentCrdsSubmarinerIo_endpoints_crdYaml, map[string]*bintree{}},
				"submariner.io_gateways_crd.yaml":                        {manifestsAgentCrdsSubmarinerIo_gateways_crdYaml, map[string]*bintree{}},
				"submariner.io_servicediscoveries_crd.yaml":              {manifestsAgentCrdsSubmarinerIo_servicediscoveries_crdYaml, map[string]*bintree{}},
				"submariner.io_submariners_crd.yaml":                     {manifestsAgentCrdsSubmarinerIo_submariners_crdYaml, map[string]*bintree{}},
			}},
			"operator": {nil, map[string]*bintree{
				"submariner-operator-deployment.yaml":      {manifestsAgentOperatorSubmarinerOperatorDeploymentYaml, map[string]*bintree{}},
				"submariner.io-servicediscoveries-cr.yaml": {manifestsAgentOperatorSubmarinerIoServicediscoveriesCrYaml, map[string]*bintree{}},
				"submariner.io-submariners-cr.yaml":        {manifestsAgentOperatorSubmarinerIoSubmarinersCrYaml, map[string]*bintree{}},
			}},
			"rbac": {nil, map[string]*bintree{
				"submariner-lighthouse-clusterrole.yaml":        {manifestsAgentRbacSubmarinerLighthouseClusterroleYaml, map[string]*bintree{}},
				"submariner-lighthouse-clusterrolebinding.yaml": {manifestsAgentRbacSubmarinerLighthouseClusterrolebindingYaml, map[string]*bintree{}},
				"submariner-lighthouse-serviceaccount.yaml":     {manifestsAgentRbacSubmarinerLighthouseServiceaccountYaml, map[string]*bintree{}},
				"submariner-operator-clusterrole.yaml":          {manifestsAgentRbacSubmarinerOperatorClusterroleYaml, map[string]*bintree{}},
				"submariner-operator-clusterrolebinding.yaml":   {manifestsAgentRbacSubmarinerOperatorClusterrolebindingYaml, map[string]*bintree{}},
				"submariner-operator-namespace.yaml":            {manifestsAgentRbacSubmarinerOperatorNamespaceYaml, map[string]*bintree{}},
				"submariner-operator-role.yaml":                 {manifestsAgentRbacSubmarinerOperatorRoleYaml, map[string]*bintree{}},
				"submariner-operator-rolebinding.yaml":          {manifestsAgentRbacSubmarinerOperatorRolebindingYaml, map[string]*bintree{}},
				"submariner-operator-serviceaccount.yaml":       {manifestsAgentRbacSubmarinerOperatorServiceaccountYaml, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
