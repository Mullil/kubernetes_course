First run 
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

Deploy the ping-pong-app with 

```Bash
kubectl apply -f manifests/
```

and deploy the log-output app with
```Bash
kubectl apply -f ../log_output/manifests
```

Now the log-output app can be accessed in http://localhost:8081/ and ping-pong app in http://localhost:8081/pingpong
