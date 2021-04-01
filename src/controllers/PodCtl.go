package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi/src/services"
)

type PodCtl struct {
	PodService *services.PodService `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}
func(this *PodCtl) GetAll(c *gin.Context) goft.Json{
	return this.PodService.ListByNs("default")
}
func(this *PodCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/pods",this.GetAll)
}
func(*PodCtl) Name() string{
	return "PodCtl"
}