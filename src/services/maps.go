package services

import (
	"fmt"
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/networking/v1beta1"
	"k8sapi/src/models"
	"reflect"
	"sort"
	"sync"
)


type MapItems []*MapItem
type MapItem struct {
	key string
	value interface{}
}
func(this *MapItem) String() string{
	return this.key
}
//把sync.map  转为 自定义切片
func convertToMapItems(m sync.Map) MapItems{
	items:=make(MapItems,0)
	m.Range(func(key, value interface{}) bool {
		items=append(items,&MapItem{key:key.(string),value:value})
		return true
	})
	return items
}
func(this MapItems) Len() int{
	return len(this)
}
func(this MapItems) Less(i, j int) bool{
	return this[i].key<this[j].key
}
func(this MapItems) Swap(i, j int){
	this[i],this[j]=this[j],this[i]
}



//对deployments的集合进行定义
type DeploymentMap struct {
	data sync.Map  // [key string] []*v1.Deployment    key=>namespace
}
//添加
func(this *DeploymentMap) Add(dep *v1.Deployment){

	if list,ok:=this.data.Load(dep.Namespace);ok{
		list=append(list.([]*v1.Deployment),dep)
		this.data.Store(dep.Namespace,list)
	}else{
		this.data.Store(dep.Namespace,[]*v1.Deployment{dep})
	}
}
//更新
func(this *DeploymentMap) Update(dep *v1.Deployment) error {
	if list,ok:=this.data.Load(dep.Namespace);ok{
		for i,range_dep:=range list.([]*v1.Deployment){
			if range_dep.Name==dep.Name{
				list.([]*v1.Deployment)[i]=dep
			}
		}
		return nil
	}
	return fmt.Errorf("deployment-%s not found",dep.Name)
}
// 删除
func(this *DeploymentMap) Delete(dep *v1.Deployment){
	if list,ok:=this.data.Load(dep.Namespace);ok{
		for i,range_dep:=range list.([]*v1.Deployment){
			if range_dep.Name==dep.Name{
				newList:= append(list.([]*v1.Deployment)[:i], list.([]*v1.Deployment)[i+1:]...)
				this.data.Store(dep.Namespace,newList)
				break
			}
		}
	}
}
func(this *DeploymentMap) ListByNS(ns string) ([]*v1.Deployment,error){
	if list,ok:=this.data.Load(ns);ok {
		 return  list.([]*v1.Deployment),nil
	}
	return nil,fmt.Errorf("record not found")
}
func(this *DeploymentMap) GetDeployment(ns string,depname string) (*v1.Deployment,error){
	if list,ok:=this.data.Load(ns);ok {
		for _,item:=range list.([]*v1.Deployment){
			if item.Name==depname{
				return item,nil
			}
		}
	}
	return nil,fmt.Errorf("record not found")
}


type CoreV1Pods []*corev1.Pod
func(this CoreV1Pods) Len() int{
	return len(this)
}
func(this CoreV1Pods) Less(i, j int) bool{
	//根据时间排序    正排序
	return this[i].CreationTimestamp.Time.Before(this[j].CreationTimestamp.Time)
}
func(this CoreV1Pods) Swap(i, j int){
	this[i],this[j]=this[j],this[i]
}
// 保存Pod集合
type PodMapStruct struct {
	data sync.Map  // [key string] []*v1.Pod    key=>namespace
}
func(this *PodMapStruct) ListByNs(ns string) []*corev1.Pod{
	if list,ok:=this.data.Load(ns);ok{
		 ret:=list.([]*corev1.Pod)
		 sort.Sort(CoreV1Pods(ret))//排序
		 return ret
	}
	return nil
}
func(this *PodMapStruct) Get(ns string,podName string) *corev1.Pod{
	if list,ok:=this.data.Load(ns);ok{
		for _,pod:=range list.([]*corev1.Pod){
			if pod.Name==podName{
				return pod
			}
		}
	}
	return nil
}
func(this *PodMapStruct) Add(pod *corev1.Pod){
	if list,ok:=this.data.Load(pod.Namespace);ok{
		list=append(list.([]*corev1.Pod),pod)
		this.data.Store(pod.Namespace,list)
	}else{
		this.data.Store(pod.Namespace,[]*corev1.Pod{pod})
	}
}
func(this *PodMapStruct) Update(pod *corev1.Pod) error {
	if list,ok:=this.data.Load(pod.Namespace);ok{
		for i,range_pod:=range list.([]*corev1.Pod){
			if range_pod.Name==pod.Name{
				list.([]*corev1.Pod)[i]=pod
			}
		}
		return nil
	}
	return fmt.Errorf("Pod-%s not found",pod.Name)
}
func(this *PodMapStruct) Delete(pod *corev1.Pod){
	if list,ok:=this.data.Load(pod.Namespace);ok{
		for i,range_pod:=range list.([]*corev1.Pod){
			if range_pod.Name==pod.Name{
				newList:= append(list.([]*corev1.Pod)[:i], list.([]*corev1.Pod)[i+1:]...)
				this.data.Store(pod.Namespace,newList)
				break
			}
		}
	}
}
//根据标签获取 POD列表
func(this *PodMapStruct) ListByLabels(ns string,labels []map[string]string) ([]*corev1.Pod,error){
	ret:=make([]*corev1.Pod,0)
	if list,ok:=this.data.Load(ns);ok {
		for _,pod:=range list.([]*corev1.Pod){
			for _,label:=range labels{
				if reflect.DeepEqual(pod.Labels,label){  //标签完全匹配
					ret=append(ret,pod)
				}
			}
		}
		return ret,nil
	}
	return nil,fmt.Errorf("pods not found ")
}
func(this *PodMapStruct) DEBUG_ListByNS(ns string) ([]*corev1.Pod){
	ret:=make([]*corev1.Pod,0)
	if list,ok:=this.data.Load(ns);ok {
		for _,pod:=range list.([]*corev1.Pod){
			ret=append(ret,pod)
		}

	}
	return ret
}


// namespace相关
type NsMapStruct struct {
	data sync.Map  // [key string] []*corev1.Namespace    key=>namespace的名称
}
func(this *NsMapStruct) Get(ns string) *corev1.Namespace{
	if item,ok:=this.data.Load(ns);ok{
		 return item.(*corev1.Namespace)
	}
	return nil
}
func(this *NsMapStruct) Add(ns *corev1.Namespace){
	this.data.Store(ns.Name,ns)
}
func(this *NsMapStruct) Update(ns *corev1.Namespace) {
	this.data.Store(ns.Name,ns)
}
func(this *NsMapStruct) Delete(ns *corev1.Namespace){
	this.data.Delete(ns.Name)
}
//显示所有的 namespace
func(this *NsMapStruct) ListAll()[]*models.NsModel{

	 //this.data.Range(func(key, value interface{}) bool {
		//ret=append(ret,&models.NsModel{Name:key.(string)})
	 //	return true
	 //})
	 items:=convertToMapItems(this.data)
	 sort.Sort(items)
	  ret:=make([]*models.NsModel,len(items))
	 for index,item:=range items{
		 ret[index]=&models.NsModel{Name:item.key}
	 }

	return ret
}

// event 事件map 相关
// EventSet 集合 用来保存事件, 只保存最新的一条
type EventMapStruct struct {
	data sync.Map   // [key string] *v1.Event
	// key=>namespace+"_"+kind+"_"+name 这里的name 不一定是pod ,这样确保唯一
}
func(this *EventMapStruct) GetMessage(ns string,kind string,name string) string{
	key:=fmt.Sprintf("%s_%s_%s",ns,kind,name)
	if v,ok:=this.data.Load(key);ok{
		return v.(*corev1.Event).Message
	}
	return ""
}


type V1Beta1Ingress []*v1beta1.Ingress
func(this V1Beta1Ingress) Len() int{
	return len(this)
}
func(this V1Beta1Ingress) Less(i, j int) bool{
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func(this V1Beta1Ingress) Swap(i, j int){
	this[i],this[j]=this[j],this[i]
}
type IngressMapStruct struct {
	data sync.Map   // [ns string] []*v1beta1.Ingress
}
//获取单个Ingress
func(this *IngressMapStruct) Get(ns string,name string) *v1beta1.Ingress{
	if items,ok:=this.data.Load(ns);ok{
		for _,item:=range items.([]*v1beta1.Ingress){
			if item.Name==name{
				return item
			}
		}
	}
	return nil
}
func(this *IngressMapStruct) Add(ingress *v1beta1.Ingress){
	if list,ok:=this.data.Load(ingress.Namespace);ok{
		list=append(list.([]*v1beta1.Ingress),ingress)
		this.data.Store(ingress.Namespace,list)
	}else{
		this.data.Store(ingress.Namespace,[]*v1beta1.Ingress{ingress})
	}
}
func(this *IngressMapStruct) Update(ingress *v1beta1.Ingress) error {
	if list,ok:=this.data.Load(ingress.Namespace);ok{
		for i,range_pod:=range list.([]*v1beta1.Ingress){
			if range_pod.Name==ingress.Name{
				list.([]*v1beta1.Ingress)[i]=ingress
			}
		}
	 return nil
	}
	return fmt.Errorf("ingress-%s not found",ingress.Name)
}
func(this *IngressMapStruct) Delete(ingress *v1beta1.Ingress){
	if list,ok:=this.data.Load(ingress.Namespace);ok{
		for i,range_ingress:=range list.([]*v1beta1.Ingress){
			if range_ingress.Name==ingress.Name{
				newList:= append(list.([]*v1beta1.Ingress)[:i], list.([]*v1beta1.Ingress)[i+1:]...)
				this.data.Store(ingress.Namespace,newList)
				break
			}
		}
	}
}
func(this *IngressMapStruct) ListAll(ns string)[]*models.IngressModel{
	     if list,ok:=this.data.Load(ns);ok{
	     	 newList:=list.([]*v1beta1.Ingress)
			 sort.Sort(V1Beta1Ingress(newList))//  按时间倒排序
			 ret:=make([]*models.IngressModel,len(newList))
			 for i,item:=range newList{
				 ret[i]=&models.IngressModel{
				 	Name:item.Name,
				 	CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
				 	NameSpace:item.Namespace,
				 }
			 }
			 return ret
		 }
	     return []*models.IngressModel{} //返回空列表
}