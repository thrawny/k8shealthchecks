### k8s healthchecks demo

Build the docker image and deploy the kubernetes manifests to your local dev cluster
```bash
docker build -t k8shealthcheck .
kubectl apply -f kubernetes/k8shealthcheck.yaml
kubectl scale deployment k8shealthcheck --replicas 3
kc exec k8shealthcheck-5468f88d57-csfmr -- touch notready
kubectl get pods
```




