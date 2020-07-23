package main

import (
	"flag"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/discovery"
	cacheddiscovery "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/scale"
	"k8s.io/client-go/scale/informers"
	"k8s.io/client-go/scale/listers"
	"k8s.io/client-go/tools/cache"
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
	informer := informers.NewScaleInformer(scaleGetter, metav1.NamespaceAll, tappGroupVersion, time.Minute)
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			scale, ok := obj.(*autoscalingv1.Scale)
			if ok {
				klog.Infof("add scale object: %v", *scale)
			}
		},
		DeleteFunc: func(obj interface{}) {
			scale, ok := obj.(*autoscalingv1.Scale)
			if ok {
				klog.Infof("delete scale object: %v", *scale)
			}
		},
	})
	go informer.Run(make(chan struct{}))
	if err := wait.PollInfinite(time.Second, func() (done bool, err error) {
		return informer.HasSynced(), nil
	}); err != nil {
		klog.Fatal(err)
	}
	scaleLister := listers.NewScaleLister(informer.GetIndexer())
	scales, err := scaleLister.Scales("default").List(labels.Everything())
	if err != nil {
		klog.Fatal(err)
	}
	for _, scale := range scales {
		klog.Infof("scale object: %v", *scale)
	}
	time.Sleep(time.Minute*10)
}
