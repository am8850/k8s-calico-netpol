# Kubernets Network Policies with Calico

## Setup


```bash
kind create cluster --config=config.yml
```

```yaml
```

```bash
```

## Create a 3-tier app

```bash
kubectl create deployment frontend --image=nginx:alpine --replicas=3 --port=80
kubectl expose deploy frontend --port=80 --target-port=
kubectl create deployment api --image=nginx:alpine --replicas=3 --port=80
kubectl expose deploy api --port=80 --target-port=80
kubectl create deployment frontend --image=nginx:alpine --replicas=1 --port=80
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
