Create a k3d cluster with
```Bash
k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
```

Deploy with 

```Bash
kubectl apply -f manifests/deployment.yaml

kubectl apply -f manifests/service.yaml
```

Now the application can be accessed in http://localhost:8082/
