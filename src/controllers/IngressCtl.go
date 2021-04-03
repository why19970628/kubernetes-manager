package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi/src/services"
)

type IngressCtl struct{
	IngressMap *services.IngressMapStruct `inject:"-"`
}
func NewIngressCtl() *IngressCtl{
  return &IngressCtl{}
}
func(*IngressCtl)  Name() string{
	 return "IngressCtl"
}
func(this *IngressCtl) ListAll(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	return gin.H{
		"code":20000,
		"data":this.IngressMap.ListAll(ns), //暂时 不分页
	}
}
func(this *IngressCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/ingress",this.ListAll)
}
