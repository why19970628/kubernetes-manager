package configs

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"k8sapi/src/services"
	"log"
)

type K8sConfig struct {
	DepHandler *services.DepHandler `inject:"-"`
	PodHandler *services.PodHandler `inject:"-"`
	NsHandler *services.NsHandler `inject:"-"`
	EventHandler *services.EventHandler `inject:"-"`
	IngressHandler *services.IngressHandler `inject:"-"`
}
func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}
//初始化客户端
func(*K8sConfig) InitClient() *kubernetes.Clientset{
	config:=&rest.Config{
		Host:"http://124.70.204.12:8009",
 	}
	c,err:=kubernetes.NewForConfig(config)
	if err!=nil{
		log.Fatal(err)
	}
	return c
}
//初始化Informer
func(this *K8sConfig) InitInformer() informers.SharedInformerFactory{
	fact:=informers.NewSharedInformerFactory(this.InitClient(), 0)

	depInformer:=fact.Apps().V1().Deployments()
	depInformer.Informer().AddEventHandler(this.DepHandler)

	podInformer:=fact.Core().V1().Pods() //监听pod
	podInformer.Informer().AddEventHandler(this.PodHandler)

	nsInformer:=fact.Core().V1().Namespaces() //监听namespace
	nsInformer.Informer().AddEventHandler(this.NsHandler)

	eventInformer:=fact.Core().V1().Events() //监听event
	eventInformer.Informer().AddEventHandler(this.EventHandler)


	 IngressInformer:=fact.Networking().V1beta1().Ingresses() //监听Ingress
	 IngressInformer.Informer().AddEventHandler(this.IngressHandler)

	fact.Start(wait.NeverStop)
	return fact
}