package registerForK8s

import v1 "k8s.io/api/core/v1"

// initModule 初始化时按规定数量启动各组件
func initModule() {

}

type RootModule struct {
	Pod *v1.Pod
}

func (r *RootModule) SetPod(pod *v1.Pod) {
	r.Pod = pod
}
func (r *RootModule) getPodIP() string {
	if r.Pod != nil && r.Pod.Status.PodIP != "" {
		return r.Pod.Status.PodIP
	}
	return "Pod IP not available yet"
}
