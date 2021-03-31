package services

import (
	"fmt"
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1")


//@Service
type CommonService struct {

}
func NewCommonService() *CommonService {
	return &CommonService{}
}
func(this *CommonService)  GetImages(dep v1.Deployment) string   {
	return this.GetImagesByPod(dep.Spec.Template.Spec.Containers)
}
func(*CommonService) GetImagesByPod(containers []corev1.Container) string{
	images:=containers[0].Image
	if imgLen:=len(containers);imgLen>1{
		images+=fmt.Sprintf("+其他%d个镜像",imgLen-1)
	}
	return images
}

