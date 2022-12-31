kubectl create namespace mongodb
kubectl apply -k config/rbac --namespace mongodb
kubectl apply -k config/rbac --namespace default

kubectl create -f config/manager/manager.yaml --namespace mongodb