package configs

import "k8sapi/src/services"

//注入 回调handler
type K8sHandler struct {}

func NewK8sHandler() *K8sHandler {
	return &K8sHandler{}
}
// deployment handler
func(this *K8sHandler) DepHandlers() *services.DepHandler{
	return &services.DepHandler{}
}
// pod handler
func(this *K8sHandler) PodHandlers() *services.PodHandler{
	return &services.PodHandler{}
}
//ns handler
func(this *K8sHandler) NsHandlers() *services.NsHandler{
	return &services.NsHandler{}
}

// event handler
func(this *K8sHandler) EventHandlers() *services.EventHandler{
	return &services.EventHandler{}
}
// IngressHandler
func(this *K8sHandler) IngressHandler() *services.IngressHandler{
	return &services.IngressHandler{}
}

