
export KUBECONFIG=/var/run/kubernetes/admin.kubeconfig
alias kubectl=cluster/kubectl.sh

cat <<EOF | kubectl create -f -
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: tapps.apps.tkestack.io
spec:
  group: apps.tkestack.io
  version: v1
  names:
    kind: TApp
    listKind: TAppList
    plural: tapps
    singular: tapp
  scope: Namespaced
  subresources:
    status: {}
    scale:
      labelSelectorPath: .status.scaleLabelSelector
      specReplicasPath: .spec.replicas
      statusReplicasPath: .status.replicas
EOF

cat <<EOF | kubectl create -f -
apiVersion: apps.tkestack.io/v1
kind: TApp
metadata:
  name: example-tapp
spec:
  replicas: 3
  template:
    metadata:
      labels:
        app: example-tapp
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
EOF