package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi/src/services"
)

type PodCtl struct {
	PodService *services.PodService `inject:"-"`
	Helper *services.Helper `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}
func(this *PodCtl) GetAll(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	page:=c.DefaultQuery("current","1") //当前页
	size:=c.DefaultQuery("size","5")
	//pods:=this.PodService.ListByNs(ns).([]*models.Pod) //获取到最终排序 过后的pods列表
	//
	//readyCount:=0 //就绪的pod数量
	//allCount:=0 //总数量
	//ipods:=make([]interface{},len(pods))
	//for i,pod:=range pods{
	//	allCount++
	//	ipods[i]=pod
	//	if pod.IsReady{
	//		readyCount++
	//	}
	//}
	//return gin.H{
	//	"code":20000,
	//	"data":this.Helper.PageResource(
	//		this.Helper.StrToInt(page,1),
	//		this.Helper.StrToInt(size,5),
	//		ipods).SetExt(gin.H{"ReadyNum":readyCount,"AllNum":allCount}),
	//}
	return gin.H{
		"code":20000,
		"data":this.PodService.PagePods(ns,this.Helper.StrToInt(page,1),
			this.Helper.StrToInt(size,5)),
	}

}
func(this *PodCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/pods",this.GetAll)
}
func(*PodCtl) Name() string{
	return "PodCtl"
}