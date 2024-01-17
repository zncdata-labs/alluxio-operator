/*
Copyright 2023 zncdata-labs.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"github.com/zncdata-labs/operator-go/pkg/status"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

var (
	MasterEmbedded    = 19200
	MasterRpcPort     = 19998
	MasterWebPort     = 19999
	JobMasterRpcPort  = 20001
	JobMasterWebPort  = 20002
	JobMasterEmbedded = 20003
	WorkerRpcPort     = 29999
	WorkerWebPort     = 30000
	JobWorkerRpcPort  = 30001
	JobWorkerDataPort = 30002
	JobWorkerWebPort  = 30003
)

// AlluxioSpec defines the desired state of Alluxio
type AlluxioSpec struct {
	// +kubebuilder:validation:Required
	ClusterConfig *ClusterConfigSpec `json:"clusterConfig,omitempty"`

	// +kubebuilder:validation:Required
	Master *MasterSpec `json:"master,omitempty"`

	// +kubebuilder:validation:Required
	Worker *WorkerSpec `json:"worker,omitempty"`
}

func (r *Alluxio) GetNameWithSuffix(suffix string) string {
	// return sparkHistory.GetName() + rand.String(5) + suffix
	return r.GetName() + "-" + suffix
}

type ClusterConfigSpec struct {
	// +kubebuilder:validation:Optional
	Image *ImageSpec `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=1
	Replicas *int32 `json:"replicas,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`

	// +kubebuilder:validation:Optional
	MatchLabels map[string]string `json:"matchLabels,omitempty"`

	// +kubebuilder:validation:Optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`

	// +kubebuilder:validation:Optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// +kubebuilder:validation:Optional
	Tolerations *corev1.Toleration `json:"tolerations,omitempty"`

	// +kubebuilder:validation:Optional
	Resources *corev1.ResourceRequirements `json:"resources"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	EnvVars map[string]string `json:"envVars,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	Args []string `json:"args,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraContainers []corev1.Container `json:"extraContainers,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraVolumes []corev1.Volume `json:"extraVolumes,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraVolumeMounts []corev1.VolumeMount `json:"extraVolumeMounts,omitempty"`

	// +kubebuilder:validation:Optional
	JobMaster *JobMasterSpec `json:"jobMaster,omitempty"`

	// +kubebuilder:validation:Optional
	JobWorker *JobWorkerSpec `json:"jobWorker,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	HostPID bool `json:"hostPID,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	HostNetwork bool `json:"hostNetwork,omitempty"`

	// +kubebuilder:validation:Optional
	DnsPolicy corev1.DNSPolicy `json:"dnsPolicy,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	ShareProcessNamespace bool `json:"shareProcessNamespace,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	TieredStore []*TieredStore `json:"tieredStore,omitempty"`

	// +kubebuilder:validation:Optional
	ShortCircuit *ShortCircuitSpec `json:"shortCircuit,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={"alluxio.security.stale.channel.purge.interval": "365d"}
	Properties map[string]string `json:"properties,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={"-XX:+UseContainerSupport"}
	JvmOptions []string `json:"jvmOptions,omitempty"`

	// +kubebuilder:validation:Optional
	Journal *JournalSpec `json:"journal,omitempty"`
}

func (clusterConfig *ClusterConfigSpec) GetSecurityContext() *corev1.PodSecurityContext {
	if clusterConfig != nil && clusterConfig.SecurityContext != nil {
		return clusterConfig.SecurityContext
	}
	runAsUser := int64(1000)
	runAsGroup := int64(1000)
	fsGroup := int64(1000)
	return &corev1.PodSecurityContext{
		RunAsUser:  &runAsUser,
		RunAsGroup: &runAsGroup,
		FSGroup:    &fsGroup,
	}
}

type MasterSpec struct {
	// +kubebuilder:validation:Optional
	RoleConfig *RoleMasterSpec `json:"roleConfig,omitempty"`

	// +kubebuilder:validation:Optional
	RoleGroups map[string]*RoleMasterSpec `json:"roleGroups,omitempty"`
}

func (master *MasterSpec) GetRoleGroup(instance *Alluxio, name string) *RoleMasterSpec {
	if master.RoleGroups == nil {
		return nil
	}

	clusterConfig := instance.Spec.ClusterConfig
	roleGroup := master.RoleGroups[name]
	roleConfig := master.RoleConfig

	var image ImageSpec
	var podSecurityContext *corev1.PodSecurityContext
	var jobResources corev1.ResourceRequirements
	var envVars map[string]string

	var extraContainers []corev1.Container
	var resources corev1.ResourceRequirements
	var replica int32

	hostPID := master.GetHostPID(roleGroup, roleConfig)
	hostNetwork := master.GetHostNetwork(roleGroup, roleConfig)
	dnsPolicy := master.GetDnsPolicy(roleGroup, roleConfig)
	shareProcessNamespace := master.GetShareProcessNamespace(roleGroup, roleConfig)
	args := master.GetMasterArgs(roleGroup, roleConfig)
	jobArgs := master.GetJobMasterArgs(roleGroup, roleConfig)

	if roleGroup != nil {
		if roleGroup.Image != nil {
			image = *roleGroup.Image
		} else if roleConfig.Image != nil {
			image = *roleConfig.Image
		} else if clusterConfig.Image != nil {
			image = *clusterConfig.Image
		}

		if roleGroup.SecurityContext != nil {
			podSecurityContext = roleGroup.SecurityContext
		} else {
			podSecurityContext = &corev1.PodSecurityContext{
				RunAsUser:  instance.Spec.ClusterConfig.GetSecurityContext().RunAsUser,
				RunAsGroup: instance.Spec.ClusterConfig.GetSecurityContext().RunAsGroup,
				FSGroup:    instance.Spec.ClusterConfig.GetSecurityContext().FSGroup,
			}
		}

		if roleGroup.Replicas != nil {
			replica = *roleGroup.Replicas
		} else if roleConfig.Replicas != nil {
			replica = *roleConfig.Replicas
		} else if clusterConfig.Replicas != nil {
			replica = *clusterConfig.Replicas
		} else {
			replica = 1
		}

		if roleGroup.Resources != nil {
			resources = *roleGroup.Resources
		} else if roleConfig.Resources != nil {
			resources = *roleConfig.Resources
		}

		if roleGroup.EnvVars != nil {
			envVars = roleGroup.EnvVars
		} else if instance.Spec.Worker.RoleConfig.EnvVars != nil {
			envVars = instance.Spec.Worker.RoleConfig.EnvVars
		}

		if roleGroup.JobMaster != nil {
			if roleGroup.JobMaster.Args != nil {
				jobArgs = roleGroup.JobMaster.Args
			}

			if roleGroup.JobMaster.Resources != nil {
				jobResources = *roleGroup.JobMaster.Resources
			}
		}
	}

	if roleGroup != nil && roleGroup.ExtraContainers != nil {
		extraContainers = roleGroup.ExtraContainers
	} else if instance.Spec.Worker.RoleConfig.ExtraContainers != nil {
		extraContainers = instance.Spec.Worker.RoleConfig.ExtraContainers
	}

	masterPorts := master.GetMasterPorts(roleGroup, roleConfig)
	masterRpcPort := masterPorts.Rpc
	masterWebPort := masterPorts.Web
	masterEmbedded := masterPorts.Embedded

	jobMasterPorts := master.GetJobMasterPorts(clusterConfig, roleGroup, roleConfig)
	jobMasterRpcPort := jobMasterPorts.Rpc
	jobMasterWebPort := jobMasterPorts.Web
	jobMasterEmbedded := jobMasterPorts.Embedded

	mergedRoleGroup := &RoleMasterSpec{
		HostPID:               &hostPID,
		HostNetwork:           &hostNetwork,
		DnsPolicy:             dnsPolicy,
		ShareProcessNamespace: &shareProcessNamespace,
		Image:                 &image,
		Replicas:              &replica,
		Ports: &MasterPortsSpec{
			Rpc:      masterRpcPort,
			Web:      masterWebPort,
			Embedded: masterEmbedded,
		},
		Resources:       &resources,
		SecurityContext: podSecurityContext,
		Args:            args,
		EnvVars:         envVars,
		ExtraContainers: extraContainers,
		JobMaster: &JobMasterSpec{
			Args:      jobArgs,
			Resources: &jobResources,
			Ports: &JobMasterPortsSpec{
				Rpc:      jobMasterRpcPort,
				Web:      jobMasterWebPort,
				Embedded: jobMasterEmbedded,
			},
		},
	}

	return mergedRoleGroup
}

func (master *MasterSpec) GetHostPID(RoleGroup *RoleMasterSpec, RoleConfig *RoleMasterSpec) bool {
	if RoleGroup != nil && RoleGroup.HostPID != nil {
		return *RoleGroup.HostPID
	} else {
		return *RoleConfig.HostPID
	}
}

func (master *MasterSpec) GetHostNetwork(RoleGroup *RoleMasterSpec, RoleConfig *RoleMasterSpec) bool {
	if RoleGroup != nil && RoleGroup.HostNetwork != nil {
		return *RoleGroup.HostNetwork
	} else {
		return *RoleConfig.HostNetwork
	}
}

func (master *MasterSpec) GetDnsPolicy(RoleGroup *RoleMasterSpec, RoleConfig *RoleMasterSpec) corev1.DNSPolicy {
	if master.GetHostNetwork(RoleGroup, RoleConfig) {
		return corev1.DNSClusterFirstWithHostNet
	}
	return corev1.DNSClusterFirst
}

func (master *MasterSpec) GetShareProcessNamespace(RoleGroup *RoleMasterSpec, RoleConfig *RoleMasterSpec) bool {
	if RoleGroup != nil && RoleGroup.ShareProcessNamespace != nil {
		return *RoleGroup.ShareProcessNamespace
	} else {
		return *RoleConfig.ShareProcessNamespace
	}
}

type RoleMasterSpec struct {
	// +kubebuilder:validation:Optional
	Image *ImageSpec `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=1
	Replicas *int32 `json:"replicas,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`

	// +kubebuilder:validation:Optional
	MatchLabels map[string]string `json:"matchLabels,omitempty"`

	// +kubebuilder:validation:Optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`

	// +kubebuilder:validation:Optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// +kubebuilder:validation:Optional
	Tolerations *corev1.Toleration `json:"tolerations,omitempty"`

	// +kubebuilder:validation:Optional
	Resources *corev1.ResourceRequirements `json:"resources"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	EnvVars map[string]string `json:"envVars,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	Args []string `json:"args,omitempty"`

	// +kubebuilder:validation:Optional
	Ports *MasterPortsSpec `json:"ports,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraContainers []corev1.Container `json:"extraContainers,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraVolumes []corev1.Volume `json:"extraVolumes,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraVolumeMounts []corev1.VolumeMount `json:"extraVolumeMounts,omitempty"`

	// +kubebuilder:validation:Optional
	JobMaster *JobMasterSpec `json:"jobMaster,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	JvmOptions []string `json:"jvmOptions,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	HostPID *bool `json:"hostPID,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	HostNetwork *bool `json:"hostNetwork,omitempty"`

	// +kubebuilder:validation:Optional
	DnsPolicy corev1.DNSPolicy `json:"dnsPolicy,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	ShareProcessNamespace *bool `json:"shareProcessNamespace,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	Properties map[string]string `json:"properties,omitempty"`
}

func (master *MasterSpec) GetMasterArgs(RoleGroup *RoleMasterSpec, RoleConfig *RoleMasterSpec) []string {
	if RoleGroup != nil && RoleGroup.Args != nil {
		return RoleGroup.Args
	} else if RoleConfig != nil && RoleConfig.Args != nil {
		return RoleConfig.Args
	} else {
		return []string{"master-only", "--no-format"}
	}
}

type MasterPortsSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=19200
	Embedded int32 `json:"embedded,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=19998
	Rpc int32 `json:"rpc,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=19999
	Web int32 `json:"web,omitempty"`
}

func (master *MasterSpec) GetMasterPorts(RoleGroup *RoleMasterSpec, RoleConfig *RoleMasterSpec) *MasterPortsSpec {
	if RoleGroup != nil && RoleGroup.Ports != nil {
		return RoleGroup.Ports
	} else if RoleConfig != nil && RoleConfig.Ports != nil {
		return RoleConfig.Ports
	} else {
		return &MasterPortsSpec{
			Embedded: int32(MasterEmbedded),
			Rpc:      int32(MasterRpcPort),
			Web:      int32(MasterWebPort),
		}
	}
}

type WorkerSpec struct {
	// +kubebuilder:validation:Optional
	RoleConfig *RoleWorkerSpec `json:"roleConfig,omitempty"`

	// +kubebuilder:validation:Optional
	RoleGroups map[string]*RoleWorkerSpec `json:"roleGroups,omitempty"`
}

func (worker *WorkerSpec) GetRoleGroup(instance *Alluxio, name string) *RoleWorkerSpec {
	if worker.RoleGroups == nil {
		return nil
	}

	clusterConfig := instance.Spec.ClusterConfig
	roleGroup := worker.RoleGroups[name]
	roleConfig := worker.RoleConfig

	var image ImageSpec
	var podSecurityContext *corev1.PodSecurityContext
	var jobResources corev1.ResourceRequirements
	var envVars map[string]string
	var extraContainers []corev1.Container
	var resources corev1.ResourceRequirements
	var replica int32

	hostPID := worker.GetHostPID(roleGroup, roleConfig)
	hostNetwork := worker.GetHostNetwork(roleGroup, roleConfig)
	dnsPolicy := worker.GetDnsPolicy(roleGroup, roleConfig)
	shareProcessNamespace := worker.GetShareProcessNamespace(roleGroup, roleConfig)
	args := worker.GetWorkerArgs(roleGroup, roleConfig)
	jobArgs := worker.GetJobWorkerArgs(roleGroup, roleConfig)

	if roleGroup != nil {
		if roleGroup.Image != nil {
			image = *roleGroup.Image
		} else if roleConfig.Image != nil {
			image = *roleConfig.Image
		} else if clusterConfig.Image != nil {
			image = *clusterConfig.Image
		}

		if roleGroup.SecurityContext != nil {
			podSecurityContext = roleGroup.SecurityContext
		} else {
			podSecurityContext = &corev1.PodSecurityContext{
				RunAsUser:  instance.Spec.ClusterConfig.GetSecurityContext().RunAsUser,
				RunAsGroup: instance.Spec.ClusterConfig.GetSecurityContext().RunAsGroup,
				FSGroup:    instance.Spec.ClusterConfig.GetSecurityContext().FSGroup,
			}
		}

		if roleGroup.Replicas != nil {
			replica = *roleGroup.Replicas
		} else if roleConfig.Replicas != nil {
			replica = *roleConfig.Replicas
		} else if clusterConfig.Replicas != nil {
			replica = *clusterConfig.Replicas
		} else {
			replica = 1
		}

		if roleGroup.Resources != nil {
			resources = *roleGroup.Resources
		} else if roleConfig.Resources != nil {
			resources = *roleConfig.Resources
		}

		if roleGroup.EnvVars != nil {
			envVars = roleGroup.EnvVars
		} else {
			envVars = instance.Spec.Worker.RoleConfig.EnvVars
		}

		if roleGroup.JobWorker != nil {
			if roleGroup.JobWorker.Args != nil {
				jobArgs = roleGroup.JobWorker.Args
			}

			if roleGroup.JobWorker.Resources != nil {
				jobResources = *roleGroup.JobWorker.Resources
			}
		}
	}

	if roleGroup != nil && roleGroup.ExtraContainers != nil {
		extraContainers = roleGroup.ExtraContainers
	} else if instance.Spec.Worker.RoleConfig.ExtraContainers != nil {
		extraContainers = instance.Spec.Worker.RoleConfig.ExtraContainers
	}

	workerPorts := worker.GetWorkerPorts(roleGroup, roleConfig)
	workerRpcPort := workerPorts.Rpc
	workerWebPort := workerPorts.Web

	jobWorkerPorts := worker.GetJobWorkerPorts(clusterConfig, roleGroup, roleConfig)
	jobWorkerRpcPort := jobWorkerPorts.Rpc
	jobWorkerDataPort := jobWorkerPorts.Data
	jobWorkerWebPort := jobWorkerPorts.Web

	mergedRoleGroup := &RoleWorkerSpec{
		HostPID:               &hostPID,
		HostNetwork:           &hostNetwork,
		DnsPolicy:             dnsPolicy,
		ShareProcessNamespace: &shareProcessNamespace,
		Image:                 &image,
		Replicas:              &replica,
		Ports: &WorkerPortsSpec{
			Rpc: workerRpcPort,
			Web: workerWebPort,
		},
		Resources:       &resources,
		SecurityContext: podSecurityContext,
		Args:            args,
		EnvVars:         envVars,
		ExtraContainers: extraContainers,
		JobWorker: &JobWorkerSpec{
			Args:      jobArgs,
			Resources: &jobResources,
			Ports: &JobWorkerPortsSpec{
				Data: jobWorkerDataPort,
				Rpc:  jobWorkerRpcPort,
				Web:  jobWorkerWebPort,
			},
		},
	}

	return mergedRoleGroup
}

func (worker *WorkerSpec) GetHostPID(RoleGroup *RoleWorkerSpec, RoleConfig *RoleWorkerSpec) bool {
	if RoleGroup != nil && RoleGroup.HostPID != nil {
		return *RoleGroup.HostPID
	} else {
		return *RoleConfig.HostPID
	}
}

func (worker *WorkerSpec) GetHostNetwork(RoleGroup *RoleWorkerSpec, RoleConfig *RoleWorkerSpec) bool {
	if RoleGroup != nil && RoleGroup.HostNetwork != nil {
		return *RoleGroup.HostNetwork
	} else {
		return *RoleConfig.HostNetwork
	}
}

func (worker *WorkerSpec) GetDnsPolicy(RoleGroup *RoleWorkerSpec, RoleConfig *RoleWorkerSpec) corev1.DNSPolicy {
	if worker.GetHostNetwork(RoleGroup, RoleConfig) {
		return corev1.DNSClusterFirstWithHostNet
	}
	return corev1.DNSClusterFirst
}

func (worker *WorkerSpec) GetShareProcessNamespace(RoleGroup *RoleWorkerSpec, RoleConfig *RoleWorkerSpec) bool {
	if RoleGroup != nil && RoleGroup.ShareProcessNamespace != nil {
		return *RoleGroup.ShareProcessNamespace
	} else {
		return *RoleConfig.ShareProcessNamespace
	}
}

type RoleWorkerSpec struct {
	// +kubebuilder:validation:Optional
	Image *ImageSpec `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=1
	Replicas *int32 `json:"replicas,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`

	// +kubebuilder:validation:Optional
	MatchLabels map[string]string `json:"matchLabels,omitempty"`

	// +kubebuilder:validation:Optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`

	// +kubebuilder:validation:Optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// +kubebuilder:validation:Optional
	Tolerations *corev1.Toleration `json:"tolerations,omitempty"`

	// +kubebuilder:validation:Optional
	Resources *corev1.ResourceRequirements `json:"resources"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	EnvVars map[string]string `json:"envVars,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	Args []string `json:"args,omitempty"`

	// +kube:validation:Optional
	Ports *WorkerPortsSpec `json:"ports,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraContainers []corev1.Container `json:"extraContainers,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraVolumes []corev1.Volume `json:"extraVolumes,omitempty"`

	// +kubebuilder:validation:Optional
	ExtraVolumeMounts []corev1.VolumeMount `json:"extraVolumeMounts,omitempty"`

	// +kubebuilder:validation:Optional
	JobWorker *JobWorkerSpec `json:"jobWorker,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	JvmOptions []string `json:"jvmOptions,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	HostPID *bool `json:"hostPID,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	HostNetwork *bool `json:"hostNetwork,omitempty"`

	// +kubebuilder:validation:Optional
	DnsPolicy corev1.DNSPolicy `json:"dnsPolicy,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	ShareProcessNamespace *bool `json:"shareProcessNamespace,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	Properties map[string]string `json:"properties,omitempty"`
}

func (worker *WorkerSpec) GetWorkerArgs(RoleGroup *RoleWorkerSpec, RoleConfig *RoleWorkerSpec) []string {
	if RoleGroup != nil && RoleGroup.Args != nil {
		return RoleGroup.Args
	} else if RoleConfig != nil && RoleConfig.Args != nil {
		return RoleConfig.Args
	} else {
		return []string{"worker-only", "--no-format"}
	}
}

type WorkerPortsSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=29999
	Rpc int32 `json:"rpc,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=30000
	Web int32 `json:"web,omitempty"`
}

func (worker *WorkerSpec) GetWorkerPorts(RoleGroup *RoleWorkerSpec, RoleConfig *RoleWorkerSpec) *WorkerPortsSpec {
	if RoleGroup != nil && RoleGroup.Ports != nil {
		return RoleGroup.Ports
	} else if RoleConfig != nil && RoleConfig.Ports != nil {
		return RoleConfig.Ports
	} else {
		return &WorkerPortsSpec{
			Rpc: int32(WorkerRpcPort),
			Web: int32(WorkerWebPort),
		}
	}
}

type JobMasterSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default={"job-master"}
	Args []string `json:"args,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	Properties map[string]string `json:"properties,omitempty"`

	// +kubebuilder:validation:Optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	Ports *JobMasterPortsSpec `json:"ports,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	JvmOptions []string `json:"jvmOptions,omitempty"`
}

func (master *MasterSpec) GetJobMasterArgs(RoleGroup *RoleMasterSpec, RoleConfig *RoleMasterSpec) []string {
	if RoleGroup != nil && RoleGroup.JobMaster != nil && RoleGroup.JobMaster.Args != nil {
		return RoleGroup.JobMaster.Args
	} else if RoleConfig != nil && RoleConfig.JobMaster != nil && RoleConfig.JobMaster.Args != nil {
		return RoleConfig.JobMaster.Args
	} else {
		return []string{"job-master"}
	}
}

type JobMasterPortsSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=20001
	Rpc int32 `json:"rpc,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=20002
	Web int32 `json:"web,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=20003
	Embedded int32 `json:"embedded,omitempty"`
}

func (master *MasterSpec) GetJobMasterPorts(ClusterConfig *ClusterConfigSpec, RoleGroup *RoleMasterSpec, RoleConfig *RoleMasterSpec) *JobMasterPortsSpec {
	if RoleGroup != nil && RoleGroup.JobMaster != nil && RoleGroup.JobMaster.Ports != nil {
		return RoleGroup.JobMaster.Ports
	} else if RoleConfig != nil && RoleConfig.JobMaster != nil && RoleConfig.JobMaster.Ports != nil {
		return RoleConfig.JobMaster.Ports
	} else if ClusterConfig != nil && ClusterConfig.JobWorker != nil && ClusterConfig.JobWorker.Ports != nil {
		return ClusterConfig.JobMaster.Ports
	} else {
		return &JobMasterPortsSpec{
			Rpc:      int32(JobMasterRpcPort),
			Web:      int32(JobMasterWebPort),
			Embedded: int32(JobMasterEmbedded),
		}
	}
}

type JobWorkerSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default={"job-worker"}
	Args []string `json:"args,omitempty"`

	// +kubebuilder:validation:Optional
	Properties map[string]string `json:"properties,omitempty"`

	// +kubebuilder:validation:Optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	Ports *JobWorkerPortsSpec `json:"ports,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={}
	JvmOptions []string `json:"jvmOptions,omitempty"`
}

func (worker *WorkerSpec) GetJobWorkerArgs(RoleGroup *RoleWorkerSpec, RoleConfig *RoleWorkerSpec) []string {
	if RoleGroup != nil && RoleGroup.JobWorker != nil && RoleGroup.JobWorker.Args != nil {
		return RoleGroup.JobWorker.Args
	} else if RoleConfig != nil && RoleConfig.JobWorker != nil && RoleConfig.JobWorker.Args != nil {
		return RoleConfig.JobWorker.Args
	} else {
		return []string{"job-worker"}
	}
}

type JobWorkerPortsSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=30001
	Rpc int32 `json:"rpc,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=30002
	Data int32 `json:"data,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=30003
	Web int32 `json:"web,omitempty"`
}

func (worker *WorkerSpec) GetJobWorkerPorts(ClusterConfig *ClusterConfigSpec, RoleGroup *RoleWorkerSpec, RoleConfig *RoleWorkerSpec) *JobWorkerPortsSpec {
	if RoleGroup != nil && RoleGroup.JobWorker != nil && RoleGroup.JobWorker.Ports != nil {
		return RoleGroup.JobWorker.Ports
	} else if RoleConfig != nil && RoleConfig.JobWorker != nil && RoleConfig.JobWorker.Ports != nil {
		return RoleConfig.JobWorker.Ports
	} else if ClusterConfig != nil && ClusterConfig.JobWorker != nil && ClusterConfig.JobWorker.Ports != nil {
		return ClusterConfig.JobWorker.Ports
	} else {
		return &JobWorkerPortsSpec{
			Rpc:  int32(JobWorkerRpcPort),
			Data: int32(JobWorkerDataPort),
			Web:  int32(JobWorkerWebPort),
		}
	}
}

type ImageSpec struct {

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=alluxio/alluxio
	Repository string `json:"repository"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=latest
	Tag string `json:"tag"`

	// +kubebuilder:validation:enum=Always;Never;IfNotPresent
	// +kubebuilder:default=IfNotPresent
	PullPolicy corev1.PullPolicy `json:"pullPolicy"`
}

type TieredStore struct {
	Level      int32   `json:"level"`
	Alias      string  `json:"alias"`
	MediumType string  `json:"mediumType"`
	Path       string  `json:"path"`
	Type       string  `json:"type"`
	Quota      string  `json:"quota"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
}

type ShortCircuitSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="uuid"
	Policy string `json:"policy,omitempty"`

	// +kubebuilder:validation:hostPath,persistentVolumeClaim
	// +kubebuilder:default="hostPath"
	VolumeType string `json:"volumeType,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="100Mi"
	Size string `json:"size,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="alluxio-worker-domain-socket"
	PvcName string `json:"pvcName,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="standard"
	StorageClass string `json:"storageClass,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="ReadWriteOnce"
	AccessMode string `json:"accessMode,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="/tmp/"
	HosePath string `json:"path,omitempty"`
}

type JournalSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="UFS"
	Type string `json:"type,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="/journal"
	Folder string `json:"folder,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="local"
	UfsType string `json:"ufsType,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="persistentVolumeClaim"
	VolumeType string `json:"volumeType,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="1Gi"
	Size string `json:"size,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="standard"
	StorageClass string `json:"storageClass,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="ReadWriteOnce"
	AccessMode string `json:"accessMode,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	Medium string `json:"medium,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	RunFormat bool `json:"runFormat,omitempty"`
}

func (clusterConfig *ClusterConfigSpec) GetShortCircuit() ShortCircuitSpec {
	return ShortCircuitSpec{
		Enabled:      true,
		Policy:       "uuid",
		VolumeType:   "persistentVolumeClaim",
		Size:         "100Mi",
		PvcName:      "alluxio-worker-domain-socket",
		StorageClass: "local-path",
		AccessMode:   "ReadWriteOnce",
		HosePath:     "/tmp/alluxio-domain",
	}
}

func (clusterConfig *ClusterConfigSpec) GetJournal() JournalSpec {
	return JournalSpec{
		Type:         "UFS",
		Folder:       "/journal",
		UfsType:      "local",
		VolumeType:   "persistentVolumeClaim",
		Size:         "1Gi",
		StorageClass: "local-path",
		AccessMode:   "ReadWriteOnce",
		Medium:       "",
		RunFormat:    false,
	}
}

// AlluxioStatus defines the observed state of Alluxio
type AlluxioStatus struct {
	// +kubebuilder:validation:Optional
	Conditions []metav1.Condition `json:"condition,omitempty"`
	// +kubebuilder:validation:Optional
	URLs []StatusURL `json:"urls,omitempty"`
}

type StatusURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// SetStatusCondition updates the status condition using the provided arguments.
// If the condition already exists, it updates the condition; otherwise, it appends the condition.
// If the condition status has changed, it updates the condition's LastTransitionTime.
func (r *Alluxio) SetStatusCondition(condition metav1.Condition) {
	r.Status.SetStatusCondition(condition)
}

// InitStatusConditions initializes the status conditions to the provided conditions.
func (r *Alluxio) InitStatusConditions() {
	r.Status.InitStatus(r)
	r.Status.InitStatusConditions()
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Alluxio is the Schema for the alluxios API
type Alluxio struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlluxioSpec   `json:"spec,omitempty"`
	Status status.Status `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AlluxioList contains a list of Alluxio
type AlluxioList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alluxio `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Alluxio{}, &AlluxioList{})
}
