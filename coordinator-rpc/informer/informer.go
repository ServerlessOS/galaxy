package informer

import (
	"coordinator_rpc/registerForK8s"
	"flag"
	"fmt"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

var Kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")

func init() {
	flag.Parse()
	if *Kubeconfig == "" {
		return
	}
	config, err := clientcmd.BuildConfigFromFlags("", *Kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	factory := informers.NewSharedInformerFactory(clientset, 1*time.Second)

	// 获取 Pod informer
	podInformer := factory.Core().V1().Pods().Informer()
	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			module := new(registerForK8s.RootModule)
			module.SetPod(obj.(*v1.Pod))
			switch {
			case strings.Contains(module.Pod.Name, "virtual-node-deployment"):
				fmt.Println("registerForK8s node")
				node := &registerForK8s.Node{*module}
				node.Register()
			case strings.Contains(module.Pod.Name, "scheduler-deployment"):
				fmt.Println("registerForK8s scheduler")
				scheduler := &registerForK8s.Scheduler{*module}
				scheduler.Register()
			case strings.Contains(module.Pod.Name, "dispatcher-deployment"):
				fmt.Println("registerForK8s dispatcher")
				dispatcher := &registerForK8s.Dispatcher{*module}
				dispatcher.Register()
			}
		},
	})
	fmt.Println("handle bind success")
	stopCh := make(chan struct{})
	defer close(stopCh)

	// 启动 informer
	factory.Start(stopCh)

	// 等待 cache 同步
	if !cache.WaitForCacheSync(stopCh, podInformer.HasSynced) {
		fmt.Println("Timed out waiting for caches to sync")
		return
	}
}
