package main

import (
	"context"
	"flag"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
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
	discoveryClient := discovery.NewDiscoveryClientForConfigOrDie(config)
	//discoveryClient := client.DiscoveryClient
	resolver := scale.NewDiscoveryScaleKindResolver(discoveryClient)
	cachedDiscoveryClient := cacheddiscovery.NewMemCacheClient(discoveryClient)
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(cachedDiscoveryClient)
	stop := make(chan struct{})
	go wait.Until(func() {
		mapper.Reset()
	}, 5*time.Minute, stop)
	scaleGetter, err := scale.NewForConfig(config, mapper, dynamic.LegacyAPIPathResolverFunc, resolver)
	if err != nil {
		klog.Fatal(err)
	}
	tappGroupVersion := schema.GroupResource{Group: "apps.tkestack.io", Resource: "TApp"}
	scale, err := scaleGetter.Scales("default").Get(context.Background(), tappGroupVersion, "example-tapp", metav1.GetOptions{})
	if err != nil {
		klog.Fatal(err)
	}
	klog.Infof("scale %+v", *scale)
	// expect /apis/apps.tkestack.io/v1/namespaces/default/tapps//scale
	// got /apis/apps.tkestack.io/v1/namespaces/default/tapps/scale
	scales, err := scaleGetter.Scales("default").List(context.Background(), tappGroupVersion, metav1.ListOptions{})
	if err != nil {
		klog.Fatal(err)
	}
	klog.Infof("scale list: %+v", *scales)

	watcher, err := scaleGetter.Scales("default").Watch(context.Background(), tappGroupVersion, "", metav1.ListOptions{})
	if err != nil {
		klog.Fatal(err)
	}
	for {
		select {
		case event := <-watcher.ResultChan():
			klog.Infof("%s: %v", event.Type, event.Object)
			scale, ok := event.Object.(*autoscalingv1.Scale)
			if ok {
				klog.Infof("watch received obj is scale object: %v", *scale)
			}
		}
	}
}
