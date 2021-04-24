package configs

import "k8sapi/src/services"

//@Config
type ServiceConfig struct {}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}
func(*ServiceConfig) CommonService() *services.CommonService{
	return services.NewCommonService()
}
func(*ServiceConfig) DeploymentService() *services.DeploymentService{
	return services.NewDeploymentService()
}
func(*ServiceConfig) PodService() *services.PodService{
	return services.NewPodService()
}
func(*ServiceConfig) Helper() *services.Helper{
	return services.NewHelper()
}
func(*ServiceConfig) IngressService() *services.IngressService{
	return services.NewIngressService()
}
func(*ServiceConfig) SecretService() *services.SecretService{
	return services.NewSecretService()
}
func(*ServiceConfig) ConfigMapService() *services.ConfigMapService{
	return services.NewConfigMapService()
}
