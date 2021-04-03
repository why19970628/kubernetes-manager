package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/networking/v1beta1"
	"k8sapi/src/wscore"
	"log"
)

//处理deployment 回调的handler
type DepHandler struct {
	DepMap *DeploymentMap `inject:"-"`
	DepService *DeploymentService `inject:"-"`
}
func(this *DepHandler) OnAdd(obj interface{}){
	this.DepMap.Add(obj.(*v1.Deployment))
	ns:=obj.(*v1.Deployment).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type":"deployments",
			"result":gin.H{"ns":ns,"data":this.DepService.ListAll(ns)},
		},
	)
  }
func(this *DepHandler) OnUpdate(oldObj, newObj interface{}){
	err:=this.DepMap.Update(newObj.(*v1.Deployment))
	if err!=nil{
		log.Println(err)
	}else{
		ns:=newObj.(*v1.Deployment).Namespace
		wscore.ClientMap.SendAll(
			gin.H{
				"type":"deployments",
				"result":gin.H{"ns":ns,"data":this.DepService.ListAll(ns)},
			},
		)
 	}
}
func(this *DepHandler)	OnDelete(obj interface{}){
	if d,ok:=obj.(*v1.Deployment);ok{
		this.DepMap.Delete(d)
		ns:=obj.(*v1.Deployment).Namespace
		wscore.ClientMap.SendAll(
			gin.H{
				"type":"deployments",
				"result":gin.H{"ns":ns,"data":this.DepService.ListAll(ns)},
			},
		)
	}
}

// pod相关的回调handler
type PodHandler struct {
	PodMap *PodMapStruct `inject:"-"`
	PodService *PodService `inject:"-"`
}
func(this *PodHandler) OnAdd(obj interface{}){
	this.PodMap.Add(obj.(*corev1.Pod))
    ns:=obj.(*corev1.Pod).Namespace
    wscore.ClientMap.SendAll(
		gin.H{
			"type":"pods",
			"result":gin.H{"ns":ns,
				"data":this.PodService.PagePods(ns,1,5)},
		},
	)

}
func(this *PodHandler) OnUpdate(oldObj, newObj interface{}){
	err:=this.PodMap.Update(newObj.(*corev1.Pod))
	if err!=nil{
		log.Println(err)
	}else{
		ns:=newObj.(*corev1.Pod).Namespace
		wscore.ClientMap.SendAll(
			gin.H{
				"type":"pods",
				"result":gin.H{"ns":ns,
					"data":this.PodService.PagePods(ns,1,5)},
			},
		)
	}
}
func(this *PodHandler)	OnDelete(obj interface{}){
	if d,ok:=obj.(*corev1.Pod);ok {
		this.PodMap.Delete(d)
		ns:=obj.(*corev1.Pod).Namespace
		wscore.ClientMap.SendAll(
			gin.H{
				"type":"pods",
				"result":gin.H{"ns":ns,
					"data":this.PodService.PagePods(ns,1,5)},
			},
		)
	}
}


// namespace 相关的回调handler
type NsHandler struct {
	NsMap *NsMapStruct `inject:"-"`
}
func(this *NsHandler) OnAdd(obj interface{}){
	this.NsMap.Add(obj.(*corev1.Namespace))
}
func(this *NsHandler) OnUpdate(oldObj, newObj interface{}){
	 this.NsMap.Update(newObj.(*corev1.Namespace))

}
func(this *NsHandler)	OnDelete(obj interface{}){
	if d,ok:=obj.(*corev1.Namespace);ok{
		this.NsMap.Delete(d)
	}
}

// event 事件相关的handler
type EventHandler struct {
	EventMap *EventMapStruct  `inject:"-"`
}
func(this *EventHandler) storeData(obj interface{},isdelete bool){
	if event,ok:=obj.(*corev1.Event);ok{
		key:=fmt.Sprintf("%s_%s_%s",event.Namespace,event.InvolvedObject.Kind,event.InvolvedObject.Name)
		if !isdelete{
			this.EventMap.data.Store(key,event)
		}else{
			this.EventMap.data.Delete(key)
		}
	}
}
func(this *EventHandler) OnAdd(obj interface{}){
	this.storeData(obj,false)
}
func(this *EventHandler) OnUpdate(oldObj, newObj interface{}){
	this.storeData(newObj,false)
}
func(this *EventHandler) OnDelete(obj interface{}){
	this.storeData(obj,true)
}


// ingress相关handler
type IngressHandler struct {
	IngressMap *IngressMapStruct  `inject:"-"`
}
func(this *IngressHandler) OnAdd(obj interface{}){
	this.IngressMap.Add(obj.(*v1beta1.Ingress))
	ns:=obj.(*v1beta1.Ingress).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type":"ingress",
			"result":gin.H{"ns":ns,
				"data":this.IngressMap.ListAll(ns)},
		},
	)
}
func(this *IngressHandler) OnUpdate(oldObj, newObj interface{}){
	err:=this.IngressMap.Update(newObj.(*v1beta1.Ingress))
	if err!=nil{
		log.Println(err)
		return
	}
	ns:=newObj.(*v1beta1.Ingress).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type":"ingress",
			"result":gin.H{"ns":ns,
				"data":this.IngressMap.ListAll(ns)},
		},
	)

}
func(this *IngressHandler) OnDelete(obj interface{}){
	this.IngressMap.Delete(obj.(*v1beta1.Ingress))
	ns:=obj.(*v1beta1.Ingress).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type":"ingress",
			"result":gin.H{"ns":ns,
				"data":this.IngressMap.ListAll(ns)},
		},
	)
}