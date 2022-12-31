kubectl delete mongodbcommunity example-mongodb --namespace mongodb
https://github.com/mongodb/mongodb-kubernetes-operator/blob/master/docs/install-upgrade.md#procedure-using-kubectl
kubectl create namespace mongodb
kubectl apply -k config/rbac --namespace mongodb
kubectl apply -k config/rbac --namespace default

kubectl create -f config/manager/manager.yaml --namespace mongodb
https://github.com/mongodb/mongodb-kubernetes-operator/blob/master/docs/deploy-configure.md

kubectl apply -f config/samples/mongodb.com_v1_mongodbcommunity_cr.yaml --namespace mongodb
kubectl get mongodbcommunity --namespace mongodb