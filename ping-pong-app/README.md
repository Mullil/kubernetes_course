Create a cluster with

```Bash
gcloud container clusters create dwk-cluster --zone=europe-north1-b --cluster-version=1.33 --disk-size=32 --num-nodes=3 --machine-type=e2-micro
```

Create a namespace with

```Bash
kubectl create namespace exercises
```

Deploy the ping-pong-app with 

```Bash
kubectl apply -f manifests/
```

