package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8sapi/src/models"
	"k8sapi/src/services"
)

//@restcontroller
type ConfigMapCtl struct {
	ConfigMap *services.ConfigMapStruct `inject:"-"`
	ConfigService *services.ConfigMapService `inject:"-"`
	Client *kubernetes.Clientset  `inject:"-"`
}
func NewConfigMapCtl() *ConfigMapCtl{
	return &ConfigMapCtl{}
}
func(*ConfigMapCtl)  Name() string{
	return "ConfigMapCtl"
}
//提交 config map
func(this *ConfigMapCtl) PostConfigmap(c *gin.Context) goft.Json{
	postModel:=&models.PostConfigMapModel{}
	err:=c.ShouldBindJSON(postModel)

	goft.Error(err)
	_,err=this.Client.CoreV1().ConfigMaps(postModel.NameSpace).Create(
		c,
		&corev1.ConfigMap{
			ObjectMeta:v1.ObjectMeta{
				Name:postModel.Name,
				Namespace:postModel.NameSpace,
			},
			Data:postModel.Data,
		},
		v1.CreateOptions{},
	)
	goft.Error(err)
	return gin.H{
		"code":20000,
		"data":"OK",
	}
}
// 列出config map
func(this *ConfigMapCtl) ListAll(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	return gin.H{
		"code":20000,
		"data":this.ConfigService.ListConfigMap(ns), //暂时 不分页
	}
}
//DELETE /configmaps?ns=xx&name=xx
func(this *ConfigMapCtl) RmCm(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")
	goft.Error(this.Client.CoreV1().ConfigMaps(ns).
		Delete(c,name,v1.DeleteOptions{}))
	return gin.H{
		"code":20000,
		"data":"OK",
	}
}
func(this *ConfigMapCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/configmaps",this.ListAll)
	//goft.Handle("GET","/configmaps/:ns/:name",this.Detail)
	goft.Handle("DELETE","/configmaps",this.RmCm)
	 goft.Handle("POST","/configmaps",this.PostConfigmap)
}