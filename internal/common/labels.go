package common

import (
	alluxiov1alpha1 "github.com/zncdatadev/alluxio-operator/api/v1alpha1"
	"strings"
)

type RoleLabels struct {
	Cr   *alluxiov1alpha1.AlluxioCluster
	Name string
}

func (r *RoleLabels) GetLabels() map[string]string {
	return map[string]string{
		"app.kubernetes.io/Name":       strings.ToLower(r.Cr.Name),
		"app.kubernetes.io/component":  r.Name,
		"app.kubernetes.io/managed-by": "alluxio-operator",
	}
}
