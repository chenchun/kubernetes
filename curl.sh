#kubectl config view -o jsonpath='{"Cluster name\tServer\n"}{range .clusters[*]}{.name}{"\t"}{.cluster
# .server}{"\n"}{end}'

# Select name of cluster you want to interact with from above output:
export CLUSTER_NAME="local"

# Point to the API server refering the cluster name
APISERVER=$(kubectl config view -o jsonpath="{.clusters[?(@.name==\"$CLUSTER_NAME\")].cluster.server}")

# Gets the token value
#TOKEN=$(kubectl get secrets -o jsonpath="{.items[?(@.metadata.annotations['kubernetes\.io/service-account\
# .name']=='default')].data.token}"|base64 -d)

#curl -k -H "Authorization: Bearer ${TOKEN}" "$APISERVER/apis/apps.tkestack
# .io/v1/namespaces/default/tapps/example-tapp/scale"

#  GET /api/v1/namespaces/test/pods?watch=1&resourceVersion=10245
# https://kubernetes.io/docs/reference/using-api/api-concepts/

curl -k -v --cert /var/run/kubernetes/client-admin.crt \
  --key  /var/run/kubernetes/client-admin.key \
  --cacert /var/run/kubernetes/server-ca.crt \
"$APISERVER/apis/apps.tkestack.io/v1/namespaces/default/tapps//scale"
#"$APISERVER/apis/apps.tkestack.io/v1/watch/namespaces/default/tapps//scale"
#"$APISERVER/apis/apps.tkestack.io/v1/watch/namespaces/default/tapps/example-tapp/scale"

# watch for a single tapp scale
#"$APISERVER/apis/apps.tkestack.io/v1/watch/namespaces/default/tapps/example-tapp/scale"
#watch a single tapp
#"$APISERVER/apis/apps.tkestack.io/v1/watch/namespaces/default/tapps/example-tapp"
#watch all default namespace tapps
#"$APISERVER/apis/apps.tkestack.io/v1/watch/namespaces/default/tapps"

# Special verbs with subresources:
# /api/{version}/watch/{resource}
# /api/{version}/watch/namespaces/{namespace}/{resource}
#"$APISERVER/apis/apps.tkestack.io/v1/namespaces/default/tapps/example-tapp/scale"
