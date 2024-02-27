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

package controller

import (
	"context"
	"github.com/go-logr/logr"
	stackv1alpha1 "github.com/zncdata-labs/alluxio-operator/api/v1alpha1"
	"github.com/zncdata-labs/operator-go/pkg/status"
	utils "github.com/zncdata-labs/operator-go/pkg/util"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AlluxioReconciler reconciles a Alluxio object
type AlluxioReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

//+kubebuilder:rbac:groups=stack.zncdata.net,resources=alluxios,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=stack.zncdata.net,resources=alluxios/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=stack.zncdata.net,resources=alluxios/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Alluxio object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *AlluxioReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	r.Log.Info("Reconciling Alluxio")

	alluxio := &stackv1alpha1.Alluxio{}

	if err := r.Get(ctx, req.NamespacedName, alluxio); err != nil {
		if client.IgnoreNotFound(err) != nil {
			r.Log.Error(err, "unable to fetch Alluxio")
			return ctrl.Result{}, err
		}
		r.Log.Error(err, "Alluxio resource not found. Ignoring since object must be deleted")
		return ctrl.Result{}, err
	}

	// Get the status condition, if it exists and its generation is not the
	// same as the Alluxio's, then we need to update the status condition

	readCondition := apimeta.FindStatusCondition(alluxio.Status.Conditions, status.ConditionTypeProgressing)
	if readCondition == nil || readCondition.ObservedGeneration != alluxio.GetGeneration() {
		alluxio.InitStatusConditions()

		if err := utils.UpdateStatus(ctx, r.Client, alluxio); err != nil {
			r.Log.Error(err, "unable to update Alluxio status")
			return ctrl.Result{}, err
		}
	}

	r.Log.Info("Alluxio found", "Name", alluxio.Name)

	// Reconcile the Alluxio cluster
	clusterReconciler := NewClusterReconciler(r.Client, r.Scheme, alluxio)
	result, err := clusterReconciler.ReconcileCluster(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}
	r.Log.Info("Successfully reconciled alluxio")
	return result, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AlluxioReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&stackv1alpha1.Alluxio{}).
		Complete(r)
}
