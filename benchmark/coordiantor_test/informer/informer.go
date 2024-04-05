package informer

import (
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"strings"
	"time"
)

var (
	Dispatcher_map  = make(map[string]string, 0)
	Scheduler_map   = make(map[string]string, 0)
	Virtualnode_map = make(map[string]string, 0)
)
var Kubeconfig = flag.String("kubeconfig", "config", "absolute path to the kubeconfig file")

func init() {
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *Kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	factory := informers.NewSharedInformerFactory(clientset, 30*time.Second)

	// 获取 Pod informer
	podInformer := factory.Core().V1().Pods().Informer()

	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {

			newPod := obj.(*v1.Pod)
			if strings.Contains(newPod.Name, "virtual-node-deployment") {
				Virtualnode_map[newPod.Name] = getPodIP(newPod)
			} else if strings.Contains(newPod.Name, "scheduler-deployment") {
				Scheduler_map[newPod.Name] = getPodIP(newPod)
			} else if strings.Contains(newPod.Name, "dispatcher-deployment") {
				Dispatcher_map[newPod.Name] = getPodIP(newPod)
			}

		},
	})

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
func getPodIP(pod *v1.Pod) string {
	if pod.Status.PodIP != "" {
		//fmt.Println(pod.Name, pod.Status.PodIP)
		return pod.Status.PodIP
	}

	return "Pod IP not available yet"
}
