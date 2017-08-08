# Checkout Azure AmazonLinux Demo
```
git clone https://github.com/listonb/azure-amazonlinux-demo.git
cd azure-amazonlinux-demo
```

# Set reusable environment variables
```
export RESOURCE_GROUP=amazonlinuxdemo
export LOCATION=eastus
export CLUSTER_NAME=amazonlinux
export DNS_PREFIX=$CLUSTER_NAME
```

# Create resource group to contain demo
`az group create -l $LOCATION --name $RESOURCE_GROUP`

# Create Azure Container Registry to store images
`az acr create --resource-group $RESOURCE_GROUP --name $CLUSTER_NAME --sku Basic --admin-enabled true`

# Docker Login to the new Azure container Registry
`az acr login --resource-group $RESOURCE_GROUP --name $CLUSTER_NAME`

# Create Azure Container Service with Kubernetes orchestrator
`az acs create --orchestrator-type=kubernetes --resource-group $RESOURCE_GROUP --name=$CLUSTER_NAME --dns-prefix=$DNS_PREFIX --generate-ssh-keys`

# Install the Kubernetes CLI if you don't have it
`az acs kubernetes install-cli`

# Login to Kubernetes cluster
`az acs kubernetes get-credentials --resource-group=$RESOURCE_GROUP --name=$CLUSTER_NAME`

# See existing nodes
`kubectl get nodes`

# Build docker container that has basic HTTP server written in go outputing the contents of /etc/system-release
`docker build . -t azure-amazon-linux`

# Tag the built image to our Azure Container Registry
`docker tag azure-amazon-linux $CLUSTER_NAME.azurecr.io/azure-amazon-linux:latest`

# Push the new image to Azure Container Registry for use in the Kubernetes Cluster
`docker push $CLUSTER_NAME.azurecr.io/azure-amazon-linux:latest`

# See docker images
`docker images`

# Run the AmazonLinux Container
`kubectl run $CLUSTER_NAME --replicas=2 --image $CLUSTER_NAME.azurecr.io/azure-amazon-linux --port=80`

# See Deployment
`kubectl describe deployment $CLUSTER_NAME`

# Expose a Load balancer to the Kubernetes nodes running our image
`kubectl expose deployments $CLUSTER_NAME --port=80 --type=LoadBalancer`

# Look for the external IP Assigned to the Service
`kubectl get svc $CLUSTER_NAME`

# Scale Azure Container Service
`az acs scale --name $CLUSTER_NAME --resource-group $RESOURCE_GROUP --new-agent-count 6`

# See new nodes
`kubectl get nodes`

# Scale Kubernetes Deployment
`kubectl scale --replicas=6 deployment/$CLUSTER_NAME`

# See Deployment
`kubectl describe deployment $CLUSTER_NAME`

# Author
Bryan Liston
