package informer

import (
	"coordinator_rpc/register"
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

var Kubeconfig = flag.String("kubeconfig", "config", "absolute path to the kubeconfig file")

func init() {
	flag.Parse()
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
			module := new(register.RootModule)
			module.SetPod(obj.(*v1.Pod))
			switch {
			case strings.Contains(module.Pod.Name, "virtual-node-deployment"):
				fmt.Println("register node")
				node := &register.Node{*module}
				node.Register()
			case strings.Contains(module.Pod.Name, "scheduler-deployment"):
				fmt.Println("register scheduler")
				scheduler := &register.Scheduler{*module}
				scheduler.Register()
			case strings.Contains(module.Pod.Name, "dispatcher-deployment"):
				fmt.Println("register dispatcher")
				dispatcher := &register.Dispatcher{*module}
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
