package configs

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

type K8sConfig struct {

}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}


func(*K8sConfig) InitClient() *kubernetes.Clientset{
	config:=&rest.Config{
		Host:"http://121.36.226.197:9090",
		BearerToken: "9e92eb8825d529819191f4eb9de32ab61",
 	}
	c,err:=kubernetes.NewForConfig(config)
	if err!=nil{
		log.Fatal(err)
	}
	return c
}