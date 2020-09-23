# acm-submariner
An integration between acm and submariner

## Locally Testing With KIND

Below steps can be used to run this repo at a local environment

> Note: [`kind`](https://kind.sigs.k8s.io/) and [`kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl/) are required

1. Build the `acm-submariner` image on local by `make images`
2. Prepare clusters by `make clusters`, this will
    - Create three clustrs: `cluster1`, `clustr2` and `cluster3`, `cluster1` is the hub clsuter
      and the others are managed clusters
    - Load the local docker images to the kind cluster `cluster1`
    - Delopy the `ClusterManager` on hub cluster `cluster1` to deploy hub cluster components
    - Delopy the `Klusterlet` on `cluster2` and `cluster3` to deploy managed cluster agents
    - Make the `cluster2` and `cluster3` join to the hub cluster `cluster1`
3. Run the demo by `make demo`, this will
    - Label the managed cluter with `cluster.open-cluster-management.io/submariner-agent`
    - Create a `Clusterset` that contains `cluster2` and `cluster3` on hub cluster `cluster1`, with
      this `Clusterset`, the `acm-submariner` controller will deploy the submariner broker on the
      hub cluser and deploy the submariner agents on the managed clusters by `ManifestWorks`
    - Create a service on managed cluster `cluster3` and export it, after the service is created,
      the submariner will import this service to managed clusters
    - Access the exported service on managed cluster `cluster2`
