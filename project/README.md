Create a k3d cluster with

```Bash
gcloud container clusters create dwk-cluster --zone=europe-north1-b --cluster-version=1.33 --disk-size=32 --num-nodes=3 --machine-type=e2-micro
```

Create the project namespace with

```Bash
kubectl create namespace project
```

Deploy with 

```Bash
kubectl apply -k .
```

Then get the address of the app with

```Bash
kubectl get ing -n project project
```
