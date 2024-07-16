# Requirements
To run needs golang >= 1.22 and Gcloud SDK. Your account must have Kubernetes Engine Cluster Admin permissions.

# Previous
You need a cluster configured, when you send `down` as param this code reduce to 0 replicas all deployments in one namespace. Is necesary all deployments associated with an environment must have the nodeSelector in the same node-pool and namespace.

Whe you send `up` as param, recreates a node pool and set al deployment replicas to 1

In this case you can create a test node-pool, namespace and deployment with next steps:

```bash
kubectl config use-context ${your-gke-context}
kubectl create namespace tiny-nspace
gcloud container node-pools create tiny-pool --cluster ${your-cluster} --machine-type=g1-small --disk-size=20 --num-nodes=1 --zone ${cluster-zone} --async
kubectl apply -f ./nginx-deployment.yaml 
```

# Test 
To run in local use (default port is 5432): 
```bash
go mod tidy
go run main.go
```

# Request
Send a `POST` method to `${your-domain}/k8s/:action` where `:action` can be `down` or `up`

## Body Examples
### Down
```JSON
{
	"cluster": "tiny-cluster",
	"zone": "us-central1-a",
	"project": "tiny-project",
	"namespace": "tiny-nspace",
	"nodepool": "tiny-pool"
}
```

### Up

```JSON
	"cluster": "tiny-cluster",
	"zone": "us-central1-a",
	"project": "tiny-project",
	"namespace": "tiny-nspace",
	"nodepool": "tiny-pool",
	"machineSpecs": {
		"type": "g1-small",
		"disk": "20",
		"numNodes": "1"
}
```

# Schedule

Add your `cloudbuild.yaml`, configure a Cloud Build Trigger and deploy as Cloud Run, use Cloud Schdeuler to call API when you need.


Good Luck Have Fun.

