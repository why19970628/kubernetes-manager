package configs

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8sapi/src/core"
	"log"
)

type K8sConfig struct {
	DepHandler *core.DepHandler `inject:"-"`
}


func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}


func(*K8sConfig) InitClient() *kubernetes.Clientset{
	config:=&rest.Config{
		Host:"http://121.36.226.197:9090",
		BearerToken: "9e92eb8825d529819191f4eb9de32ab6",
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

	fact.Start(wait.NeverStop)

	return fact
}