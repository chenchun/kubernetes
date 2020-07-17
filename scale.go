package main

import (
	"context"
	"flag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/deprecated"
	"k8s.io/client-go/discovery"
	cacheddiscovery "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/scale"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"time"
)

func main() {
	flag.Parse()
	kubeconfig := "/var/run/kubernetes/admin.kubeconfig"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		klog.Fatal(err)
	}
	client := clientset.NewForConfigOrDie(config)
	//client.
	discoveryClient := discovery.NewDiscoveryClientForConfigOrDie(config)
	resolver := scale.NewDiscoveryScaleKindResolver(discoveryClient)
	cachedDiscoveryClient := cacheddiscovery.NewMemCacheClient(discoveryClient)
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(cachedDiscoveryClient)
	stop := make(chan struct{})
	go wait.Until(func() {
		mapper.Reset()
	}, 5*time.Minute, stop)
	scaleGetter := scale.New(client.RESTClient(), mapper, dynamic.LegacyAPIPathResolverFunc, resolver)
	scale, err := scaleGetter.Scales("default").Get(context.Background(), schema.GroupResource{
		Group:    "apps.tkestack.io",
		Resource: "TApp",
	}, "example-tapp", metav1.GetOptions{})
	if err != nil {
		klog.Fatal(err)
	}
	klog.Infof("%+v", *scale)
}
