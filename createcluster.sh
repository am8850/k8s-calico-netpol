kind create cluster --config=config.yml

# Install Calico
kubectl apply -f https://docs.projectcalico.org/v3.8/manifests/calico.yaml

# Configure Calico for development
kubectl -n kube-system set env daemonset/calico-node FELIX_IGNORELOOSERPF=true
