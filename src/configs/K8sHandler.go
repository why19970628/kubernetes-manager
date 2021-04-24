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

// ServiceHandler
func(this *K8sHandler) ServiceHandler() *services.ServiceHandler{
	return &services.ServiceHandler{}
}

// SecretHandler
func(this *K8sHandler) SecretHandler() *services.SecretHandler{
	return &services.SecretHandler{}
}

// ConfigMapHandler
func(this *K8sHandler) ConfigMapHandler() *services.ConfigMapHandler{
	return &services.ConfigMapHandler{}
}
