package configs

import "k8sapi/src/core"

//注入 回调handler
type K8sHandler struct {}

func NewK8sHandler() *K8sHandler {
	return &K8sHandler{}
}
// deployment handler
func(this *K8sHandler) DepHandlers() *core.DepHandler{
	return &core.DepHandler{}
}
