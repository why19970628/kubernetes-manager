package services

import "k8sapi/src/core"

//@Service
type PodService struct {
	PodMap *core.PodMapStruct `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}
func(this *PodService) ListByNs(ns string ) interface{}{
	return this.PodMap.ListByNs(ns)
}