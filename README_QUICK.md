# Amazon Linux on Azure done Quick!
```
git clone https://github.com/listonb/azure-amazonlinux-demo.git
cd azure-amazonlinux-demo
```

```
# Set reusable environment variables
export RESOURCE_GROUP=amazonlinuxdemo
export LOCATION=eastus
# Change this to something unique
export CLUSTER_NAME=amazonlinuxdemo
export DNS_PREFIX=$CLUSTER_NAME
```

```
# Azure CLI Commands
az group create -l $LOCATION --name $RESOURCE_GROUP
az acr create --resource-group $RESOURCE_GROUP --name $CLUSTER_NAME --sku Basic --admin-enabled true
az acr login --resource-group $RESOURCE_GROUP --name $CLUSTER_NAME
az acs create --orchestrator-type=kubernetes --resource-group $RESOURCE_GROUP --name=$CLUSTER_NAME --dns-prefix=$DNS_PREFIX --generate-ssh-keys
az acs kubernetes get-credentials --resource-group=$RESOURCE_GROUP --name=$CLUSTER_NAME
```

```
# Docker Commands
docker build . -t azure-amazon-linux
docker tag azure-amazon-linux $CLUSTER_NAME.azurecr.io/azure-amazon-linux:latest
docker push $CLUSTER_NAME.azurecr.io/azure-amazon-linux:latest
```

```
# Kubernetes Commands
kubectl run $CLUSTER_NAME --replicas=2 --image $CLUSTER_NAME.azurecr.io/azure-amazon-linux --port=80
kubectl expose deployments $CLUSTER_NAME --port=80 --type=LoadBalancer

# Look for the external IP Assigned to the Service
kubectl get svc $CLUSTER_NAME
```
