//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alluxio) DeepCopyInto(out *Alluxio) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alluxio.
func (in *Alluxio) DeepCopy() *Alluxio {
	if in == nil {
		return nil
	}
	out := new(Alluxio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Alluxio) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlluxioList) DeepCopyInto(out *AlluxioList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alluxio, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlluxioList.
func (in *AlluxioList) DeepCopy() *AlluxioList {
	if in == nil {
		return nil
	}
	out := new(AlluxioList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlluxioList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlluxioSpec) DeepCopyInto(out *AlluxioSpec) {
	*out = *in
	if in.ClusterConfig != nil {
		in, out := &in.ClusterConfig, &out.ClusterConfig
		*out = new(ClusterConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSpec)
		**out = **in
	}
	if in.Master != nil {
		in, out := &in.Master, &out.Master
		*out = new(MasterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Worker != nil {
		in, out := &in.Worker, &out.Worker
		*out = new(WorkerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlluxioSpec.
func (in *AlluxioSpec) DeepCopy() *AlluxioSpec {
	if in == nil {
		return nil
	}
	out := new(AlluxioSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlluxioStatus) DeepCopyInto(out *AlluxioStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.URLs != nil {
		in, out := &in.URLs, &out.URLs
		*out = make([]StatusURL, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlluxioStatus.
func (in *AlluxioStatus) DeepCopy() *AlluxioStatus {
	if in == nil {
		return nil
	}
	out := new(AlluxioStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseGroupSpec) DeepCopyInto(out *BaseGroupSpec) {
	*out = *in
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseGroupSpec.
func (in *BaseGroupSpec) DeepCopy() *BaseGroupSpec {
	if in == nil {
		return nil
	}
	out := new(BaseGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUResource) DeepCopyInto(out *CPUResource) {
	*out = *in
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUResource.
func (in *CPUResource) DeepCopy() *CPUResource {
	if in == nil {
		return nil
	}
	out := new(CPUResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigSpec) DeepCopyInto(out *ClusterConfigSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TieredStore != nil {
		in, out := &in.TieredStore, &out.TieredStore
		*out = make([]*TieredStore, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(TieredStore)
				**out = **in
			}
		}
	}
	if in.ShortCircuit != nil {
		in, out := &in.ShortCircuit, &out.ShortCircuit
		*out = new(ShortCircuitSpec)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.JvmOptions != nil {
		in, out := &in.JvmOptions, &out.JvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Journal != nil {
		in, out := &in.Journal, &out.Journal
		*out = new(JournalSpec)
		**out = **in
	}
	if in.Listener != nil {
		in, out := &in.Listener, &out.Listener
		*out = new(ListenerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigSpec.
func (in *ClusterConfigSpec) DeepCopy() *ClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigOverridesSpec) DeepCopyInto(out *ConfigOverridesSpec) {
	*out = *in
	if in.OverrideConfig != nil {
		in, out := &in.OverrideConfig, &out.OverrideConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigOverridesSpec.
func (in *ConfigOverridesSpec) DeepCopy() *ConfigOverridesSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigOverridesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseInlineSpec) DeepCopyInto(out *DatabaseInlineSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseInlineSpec.
func (in *DatabaseInlineSpec) DeepCopy() *DatabaseInlineSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseInlineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSpec) DeepCopyInto(out *DatabaseSpec) {
	*out = *in
	if in.Inline != nil {
		in, out := &in.Inline, &out.Inline
		*out = new(DatabaseInlineSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSpec.
func (in *DatabaseSpec) DeepCopy() *DatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobMasterPortsSpec) DeepCopyInto(out *JobMasterPortsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobMasterPortsSpec.
func (in *JobMasterPortsSpec) DeepCopy() *JobMasterPortsSpec {
	if in == nil {
		return nil
	}
	out := new(JobMasterPortsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobMasterSpec) DeepCopyInto(out *JobMasterSpec) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = new(JobMasterPortsSpec)
		**out = **in
	}
	if in.JvmOptions != nil {
		in, out := &in.JvmOptions, &out.JvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobMasterSpec.
func (in *JobMasterSpec) DeepCopy() *JobMasterSpec {
	if in == nil {
		return nil
	}
	out := new(JobMasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobWorkerPortsSpec) DeepCopyInto(out *JobWorkerPortsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobWorkerPortsSpec.
func (in *JobWorkerPortsSpec) DeepCopy() *JobWorkerPortsSpec {
	if in == nil {
		return nil
	}
	out := new(JobWorkerPortsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobWorkerSpec) DeepCopyInto(out *JobWorkerSpec) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = new(JobWorkerPortsSpec)
		**out = **in
	}
	if in.JvmOptions != nil {
		in, out := &in.JvmOptions, &out.JvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobWorkerSpec.
func (in *JobWorkerSpec) DeepCopy() *JobWorkerSpec {
	if in == nil {
		return nil
	}
	out := new(JobWorkerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JournalSpec) DeepCopyInto(out *JournalSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JournalSpec.
func (in *JournalSpec) DeepCopy() *JournalSpec {
	if in == nil {
		return nil
	}
	out := new(JournalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerSpec) DeepCopyInto(out *ListenerSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerSpec.
func (in *ListenerSpec) DeepCopy() *ListenerSpec {
	if in == nil {
		return nil
	}
	out := new(ListenerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterConfigSpec) DeepCopyInto(out *MasterConfigSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraContainers != nil {
		in, out := &in.ExtraContainers, &out.ExtraContainers
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumes != nil {
		in, out := &in.ExtraVolumes, &out.ExtraVolumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumeMounts != nil {
		in, out := &in.ExtraVolumeMounts, &out.ExtraVolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.JvmOptions != nil {
		in, out := &in.JvmOptions, &out.JvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HostPID != nil {
		in, out := &in.HostPID, &out.HostPID
		*out = new(bool)
		**out = **in
	}
	if in.HostNetwork != nil {
		in, out := &in.HostNetwork, &out.HostNetwork
		*out = new(bool)
		**out = **in
	}
	if in.ShareProcessNamespace != nil {
		in, out := &in.ShareProcessNamespace, &out.ShareProcessNamespace
		*out = new(bool)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = new(MasterPortsSpec)
		**out = **in
	}
	if in.JobMaster != nil {
		in, out := &in.JobMaster, &out.JobMaster
		*out = new(JobMasterSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterConfigSpec.
func (in *MasterConfigSpec) DeepCopy() *MasterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(MasterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterPortsSpec) DeepCopyInto(out *MasterPortsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterPortsSpec.
func (in *MasterPortsSpec) DeepCopy() *MasterPortsSpec {
	if in == nil {
		return nil
	}
	out := new(MasterPortsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterRoleGroupSpec) DeepCopyInto(out *MasterRoleGroupSpec) {
	*out = *in
	in.BaseGroupSpec.DeepCopyInto(&out.BaseGroupSpec)
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(MasterConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterRoleGroupSpec.
func (in *MasterRoleGroupSpec) DeepCopy() *MasterRoleGroupSpec {
	if in == nil {
		return nil
	}
	out := new(MasterRoleGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterSpec) DeepCopyInto(out *MasterSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(MasterConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleGroups != nil {
		in, out := &in.RoleGroups, &out.RoleGroups
		*out = make(map[string]*MasterRoleGroupSpec, len(*in))
		for key, val := range *in {
			var outVal *MasterRoleGroupSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(MasterRoleGroupSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterSpec.
func (in *MasterSpec) DeepCopy() *MasterSpec {
	if in == nil {
		return nil
	}
	out := new(MasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryResource) DeepCopyInto(out *MemoryResource) {
	*out = *in
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryResource.
func (in *MemoryResource) DeepCopy() *MemoryResource {
	if in == nil {
		return nil
	}
	out := new(MemoryResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetSpec) DeepCopyInto(out *PodDisruptionBudgetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetSpec.
func (in *PodDisruptionBudgetSpec) DeepCopy() *PodDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesSpec) DeepCopyInto(out *ResourcesSpec) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(CPUResource)
		(*in).DeepCopyInto(*out)
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(MemoryResource)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(StorageResource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesSpec.
func (in *ResourcesSpec) DeepCopy() *ResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(ResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleSpec) DeepCopyInto(out *RoleSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(MasterConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleGroups != nil {
		in, out := &in.RoleGroups, &out.RoleGroups
		*out = make(map[string]*MasterRoleGroupSpec, len(*in))
		for key, val := range *in {
			var outVal *MasterRoleGroupSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(MasterRoleGroupSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleSpec.
func (in *RoleSpec) DeepCopy() *RoleSpec {
	if in == nil {
		return nil
	}
	out := new(RoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketInlineSpec) DeepCopyInto(out *S3BucketInlineSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketInlineSpec.
func (in *S3BucketInlineSpec) DeepCopy() *S3BucketInlineSpec {
	if in == nil {
		return nil
	}
	out := new(S3BucketInlineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketSpec) DeepCopyInto(out *S3BucketSpec) {
	*out = *in
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(string)
		**out = **in
	}
	if in.Inline != nil {
		in, out := &in.Inline, &out.Inline
		*out = new(S3BucketInlineSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketSpec.
func (in *S3BucketSpec) DeepCopy() *S3BucketSpec {
	if in == nil {
		return nil
	}
	out := new(S3BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShortCircuitSpec) DeepCopyInto(out *ShortCircuitSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShortCircuitSpec.
func (in *ShortCircuitSpec) DeepCopy() *ShortCircuitSpec {
	if in == nil {
		return nil
	}
	out := new(ShortCircuitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusURL) DeepCopyInto(out *StatusURL) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusURL.
func (in *StatusURL) DeepCopy() *StatusURL {
	if in == nil {
		return nil
	}
	out := new(StatusURL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageResource) DeepCopyInto(out *StorageResource) {
	*out = *in
	out.Capacity = in.Capacity.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageResource.
func (in *StorageResource) DeepCopy() *StorageResource {
	if in == nil {
		return nil
	}
	out := new(StorageResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageResourceSpec) DeepCopyInto(out *StorageResourceSpec) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(StorageResource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageResourceSpec.
func (in *StorageResourceSpec) DeepCopy() *StorageResourceSpec {
	if in == nil {
		return nil
	}
	out := new(StorageResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TieredStore) DeepCopyInto(out *TieredStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TieredStore.
func (in *TieredStore) DeepCopy() *TieredStore {
	if in == nil {
		return nil
	}
	out := new(TieredStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerConfigSpec) DeepCopyInto(out *WorkerConfigSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraContainers != nil {
		in, out := &in.ExtraContainers, &out.ExtraContainers
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumes != nil {
		in, out := &in.ExtraVolumes, &out.ExtraVolumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumeMounts != nil {
		in, out := &in.ExtraVolumeMounts, &out.ExtraVolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.JvmOptions != nil {
		in, out := &in.JvmOptions, &out.JvmOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HostPID != nil {
		in, out := &in.HostPID, &out.HostPID
		*out = new(bool)
		**out = **in
	}
	if in.HostNetwork != nil {
		in, out := &in.HostNetwork, &out.HostNetwork
		*out = new(bool)
		**out = **in
	}
	if in.ShareProcessNamespace != nil {
		in, out := &in.ShareProcessNamespace, &out.ShareProcessNamespace
		*out = new(bool)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resource.DeepCopyInto(&out.Resource)
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = new(WorkerPortsSpec)
		**out = **in
	}
	if in.JobWorker != nil {
		in, out := &in.JobWorker, &out.JobWorker
		*out = new(JobWorkerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerConfigSpec.
func (in *WorkerConfigSpec) DeepCopy() *WorkerConfigSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerPortsSpec) DeepCopyInto(out *WorkerPortsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerPortsSpec.
func (in *WorkerPortsSpec) DeepCopy() *WorkerPortsSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerPortsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerRoleGroupSpec) DeepCopyInto(out *WorkerRoleGroupSpec) {
	*out = *in
	in.BaseGroupSpec.DeepCopyInto(&out.BaseGroupSpec)
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(WorkerConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerRoleGroupSpec.
func (in *WorkerRoleGroupSpec) DeepCopy() *WorkerRoleGroupSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerRoleGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSpec) DeepCopyInto(out *WorkerSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(WorkerConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleGroups != nil {
		in, out := &in.RoleGroups, &out.RoleGroups
		*out = make(map[string]*WorkerRoleGroupSpec, len(*in))
		for key, val := range *in {
			var outVal *WorkerRoleGroupSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(WorkerRoleGroupSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSpec.
func (in *WorkerSpec) DeepCopy() *WorkerSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerSpec)
	in.DeepCopyInto(out)
	return out
}
