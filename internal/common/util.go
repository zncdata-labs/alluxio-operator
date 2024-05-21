package common

import (
	alluxiov1alpha1 "github.com/zncdatadev/alluxio-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

type ResourceNameGenerator struct {
	InstanceName string
	RoleName     string
	GroupName    string
}

// NewResourceNameGenerator new a ResourceNameGenerator
func NewResourceNameGenerator(instanceName, roleName, groupName string) *ResourceNameGenerator {
	return &ResourceNameGenerator{
		InstanceName: instanceName,
		RoleName:     roleName,
		GroupName:    groupName,
	}
}

// GenerateResourceName generate resource Name
func (r *ResourceNameGenerator) GenerateResourceName(extraSuffix string) string {
	var res string
	if r.InstanceName != "" {
		res = r.InstanceName + "-"
	}
	if r.GroupName != "" {
		res = res + r.GroupName + "-"
	}
	if r.RoleName != "" {
		res = res + r.RoleName
	} else {
		res = res[:len(res)-1]
	}
	if extraSuffix != "" {
		return res + "-" + extraSuffix
	}
	return res
}

// CreateMasterConfigMapName create configMap Name
func CreateMasterConfigMapName(instanceName string, groupName string) string {
	return NewResourceNameGenerator(instanceName, "", groupName).GenerateResourceName("config")
}

// CreateRoleGroupLoggingConfigMapName create role group logging config-map name
func CreateRoleGroupLoggingConfigMapName(instanceName string, role string, groupName string) string {
	return NewResourceNameGenerator(instanceName, role, groupName).GenerateResourceName("log4j")
}

func OverrideEnvVars(origin *[]corev1.EnvVar, override map[string]string) {
	for _, env := range *origin {
		// if env Name is in override, then override it
		if value, ok := override[env.Name]; ok {
			env.Value = value
		}
	}
}
func GetStorageClass(origin string) *string {
	if origin == "" {
		return nil
	}
	return &origin
}
func ConvertToResourceRequirements(resources *alluxiov1alpha1.ResourcesSpec) *corev1.ResourceRequirements {
	var (
		cpuMin      = resource.MustParse("100m")
		cpuMax      = resource.MustParse("500")
		memoryLimit = resource.MustParse("1Gi")
	)
	if resources != nil {
		if resources.CPU != nil && resources.CPU.Min != nil {
			cpuMin = *resources.CPU.Min
		}
		if resources.CPU != nil && resources.CPU.Max != nil {
			cpuMax = *resources.CPU.Max
		}
		if resources.Memory != nil && resources.Memory.Limit != nil {
			memoryLimit = *resources.Memory.Limit
		}
	}
	return &corev1.ResourceRequirements{
		Limits: corev1.ResourceList{
			corev1.ResourceCPU:    cpuMax,
			corev1.ResourceMemory: memoryLimit,
		},
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    cpuMin,
			corev1.ResourceMemory: memoryLimit,
		},
	}
}

// GetWorkerPorts get worker ports
func GetWorkerPorts(workerCfg *alluxiov1alpha1.WorkerRoleGroupSpec) *alluxiov1alpha1.WorkerPortsSpec {
	workerPorts := workerCfg.Config.Ports
	if workerPorts == nil {
		workerPorts = &alluxiov1alpha1.WorkerPortsSpec{
			Web: alluxiov1alpha1.WorkerWebPort,
			Rpc: alluxiov1alpha1.WorkerRpcPort,
		}
	}
	return workerPorts
}

// GetJobWorkerPorts get job worker ports
func GetJobWorkerPorts(workerCfg *alluxiov1alpha1.WorkerRoleGroupSpec) *alluxiov1alpha1.JobWorkerPortsSpec {
	jobWorkerPorts := workerCfg.Config.JobWorker.Ports
	if jobWorkerPorts == nil {
		jobWorkerPorts = &alluxiov1alpha1.JobWorkerPortsSpec{
			Web:  alluxiov1alpha1.JobWorkerWebPort,
			Rpc:  alluxiov1alpha1.JobWorkerRpcPort,
			Data: alluxiov1alpha1.JobWorkerDataPort,
		}
	}
	return jobWorkerPorts
}

func GetJournal(cluster *alluxiov1alpha1.ClusterConfigSpec) *alluxiov1alpha1.JournalSpec {
	if cluster.Journal == nil {
		defaultJournal := cluster.GetJournal()
		return &defaultJournal
	}
	return cluster.Journal
}

func CreateAlluxioLoggerVolumeMounts() corev1.VolumeMount {
	return corev1.VolumeMount{
		Name:      Log4jVolumeName(),
		MountPath: "/opt/alluxio-2.9.3/conf/log4j.properties",
		SubPath:   Log4jCfgName,
	}
}
