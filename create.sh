# Create the app deployments
kubectl create deployment frontend --image=nginx:alpine --replicas=3 --port=80
kubectl create deployment api --image=nginx:alpine --replicas=3 --port=80
kubectl create deployment frontend --image=nginx:alpine --replicas=1 --port=80

# Create the app services
kubectl expose deploy frontend --port=80 --target-port=80
kubectl expose deploy api --port=80 --target-port=80
kubectl expose deploy db --port=80 --target-port=80
