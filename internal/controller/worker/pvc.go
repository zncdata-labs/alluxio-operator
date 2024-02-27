package worker

import (
	stackv1alpha1 "github.com/zncdata-labs/alluxio-operator/api/v1alpha1"
	"github.com/zncdata-labs/alluxio-operator/internal/common"
	corev1 "k8s.io/api/core/v1"
	k8sResource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PvcReconciler struct {
	common.GeneralResourceStyleReconciler[*stackv1alpha1.Alluxio, *stackv1alpha1.WorkerRoleGroupSpec]
}

// NewService New a Service
func NewPvc(
	scheme *runtime.Scheme,
	instance *stackv1alpha1.Alluxio,
	client client.Client,
	mergedLabels map[string]string,
	mergedCfg *stackv1alpha1.WorkerRoleGroupSpec,
) *PvcReconciler {
	return &PvcReconciler{
		GeneralResourceStyleReconciler: *common.NewGeneraResourceStyleReconciler[*stackv1alpha1.Alluxio,
			*stackv1alpha1.WorkerRoleGroupSpec](
			scheme,
			instance,
			client,
			mergedLabels,
			mergedCfg,
		),
	}
}

func (s *PvcReconciler) Build(data common.ResourceBuilderData) (client.Object, error) {
	instance := s.Instance
	roleGroupName := data.GroupName
	shortCircuit := instance.Spec.ClusterConfig.GetShortCircuit()
	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      createPvcName(shortCircuit.PvcName, roleGroupName),
			Namespace: instance.Namespace,
			Labels:    s.MergedLabels,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			StorageClassName: &shortCircuit.StorageClass,
			AccessModes:      []corev1.PersistentVolumeAccessMode{corev1.PersistentVolumeAccessMode(shortCircuit.AccessMode)},
			Resources: corev1.VolumeResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: k8sResource.MustParse(shortCircuit.Size),
				},
			},
		},
	}
	return pvc, nil
}
