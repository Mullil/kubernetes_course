Delete existing k3d cluster with

```Bash
k3d cluster delete
```

Create a k3d cluster with

```Bash
k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
```

Then run

```Bash
docker exec k3d-k3s-default-agent-0 mkdir -p /tmp/kube
```

Apply volumes with

```Bash
kubectl apply -f ../volumes
```

Deploy with 

```Bash
kubectl apply -f manifests/
```

Now the application can be accessed in http://localhost:8081/
