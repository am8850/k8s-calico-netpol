# Kubernets Network Policies with Calico

## Description

In this demo, we will create a 3-tier app and setup network policies following the following best practices:

- Deny all traffic by default
- Open all traffic to the frontend only
  - This could be restricted further to a port and protocl
- Open traffic between the fronend and the api and between the api and db, but there should not be traffic from the frontend to the db
  - This could further be restricted by a port and protocl 

![foxdemo](images/)

## Create a 3-tier app

> Note: we will be using nginx and port 80, but this could be change to use other ports and applications

```bash
# Create the deployments
kubectl create deployment frontend --image=nginx:alpine --replicas=3 --port=80
kubectl create deployment api --image=nginx:alpine --replicas=3 --port=80
kubectl create deployment frontend --image=nginx:alpine --replicas=1 --port=80

# Create the services
kubectl expose deploy frontend --port=80 --target-port=
kubectl expose deploy api --port=80 --target-port=80
kubectl expose deploy db --port=80 --target-port=80
```

### Test Connectivity

```
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://frontend -T 2 # pass
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://api -T 2 # pass
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://db -T 2 # pass
```

## Block all Traffic

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: default-deny-ingress
spec:
  podSelector: {}
  policyTypes:
  - Ingress
```

### Test it

```
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://frontend -T 2 # fail
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://api -T 2 # fail
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://db -T 2 # fail
```

## Open All traffic to frontend

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-all-ingress
spec:
  podSelector:
    matchLabels:
      app : frontend
  ingress:
  - {}
  policyTypes:
  - Ingress
```

### Test it

```
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://frontend -T 2 # pass
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://api -T 2 # fail
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://db -T 2 #fail
```

## Open traffic between frontend and api

```yaml
kind: NetworkPolicy
apiVersion: networking.k8s.io/v1
metadata:
  name: frontendtoapi
spec:
  podSelector:
    matchLabels:
      app: api
  ingress:
  - from:
    - podSelector:
        matchLabels:
          app: frontend
```         

### Test it

```
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://frontend -T 2 # pass
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://api -T 2 # fail
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://db -T 2 #fail


kubectl exec <FRONTEND-NODE> -it -- curl http://api # pass
kubectl exec <FRONTEND-NODE> -it -- curl http://db # fail
kubectl exec <API-NODE> -it -- curl http://db # should fail
```

## Open traffic between api and db

```yaml
kind: NetworkPolicy
apiVersion: networking.k8s.io/v1
metadata:
  name: apitodb
spec:
  podSelector:
    matchLabels:
      app: db
  ingress:
  - from:
    - podSelector:
        matchLabels:
          app: api
```          

### Test it

```
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://frontend -T 2 # pass
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://api -T 2 # fail
kubectl run busybox --image=busybox --restart=Never --rm -it -- wget -O- http://db -T 2 # fail

kubectl exec <FRONTEND-NODE> -it -- curl http://api # pass
kubectl exec <FRONTEND-NODE> -it -- curl http://db # fail
kubectl exec <API-NODE> -it -- curl http://db # pass
```

## Other Configuration


### Kind cluster with Calico


```bash
kind create cluster --config=config.yml
```

```yaml
```

```bash
```
