package configs

import (
	"k8s.io/client-go/kubernetes"
	"k8sapi/src/services"
)

type K8sMaps struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
}
func NewK8sMaps() *K8sMaps {
	return &K8sMaps{}
}
//初始化 deploymentmap
func(this *K8sMaps) InitDepMap() *services.DeploymentMap{
	return &services.DeploymentMap{}
}
//初始化 podmap
func(this *K8sMaps) InitPodMap() *services.PodMapStruct{
	return &services.PodMapStruct{}
}
//初始化 nsmap
func(this *K8sMaps) InitNsMap() *services.NsMapStruct{
	return &services.NsMapStruct{}
}
//初始化 eventmap
func(this *K8sMaps) InitEventMap() *services.EventMapStruct{
	return &services.EventMapStruct{}
}
//初始化 ingress map
func(this *K8sMaps) InitIngressMap() *services.IngressMapStruct{
	return &services.IngressMapStruct{}
}

//初始化 Service map
func(this *K8sMaps) InitServiceMap() *services.ServiceMapStruct{
	return &services.ServiceMapStruct{}
}

//初始化 Secret map
func(this *K8sMaps) InitSecretMap() *services.SecretMapStruct{
	return &services.SecretMapStruct{}
}

//初始化 ConfigMap map
func(this *K8sMaps) InitConfigMap() *services.ConfigMapStruct{
	return &services.ConfigMapStruct{}
}


