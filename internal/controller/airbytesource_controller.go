/*
Copyright 2023.

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
	"fmt"

	"github.com/vSivarajah/airbyte-operator/airbyte"
	airbytev1 "github.com/vSivarajah/airbyte-operator/api/v1"
	v1 "github.com/vSivarajah/airbyte-operator/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// AirbyteSourceReconciler reconciles a AirbyteSource object
type AirbyteSourceReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	AirbyteClient *airbyte.AirbyteClient
}

//+kubebuilder:rbac:groups=airbyte.airbyte.poc,resources=airbytesources,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=airbyte.airbyte.poc,resources=airbytesources/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=airbyte.airbyte.poc,resources=airbytesources/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the AirbyteSource object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.0/pkg/reconcile
func (r *AirbyteSourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	// get config from yaml
	var airbyteSource v1.AirbyteSource

	if err := r.Client.Get(ctx, req.NamespacedName, &airbyteSource); err != nil {
		fmt.Println("couldn't fetch airbytesource")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	//create source
	r.AirbyteClient.CreateSource()
	
	//update source
	//if hasChanged(airbyteSource) {
		r.AirbyteClient.UpdateSource()

	//}

	//delete souce
	if !airbyteSource.DeletionTimestamp.IsZero() {
		//handle delete
		r.AirbyteClient.DeleteSource()

		//set finalizer on airbytesource object

		//update client
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AirbyteSourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&airbytev1.AirbyteSource{}).
		Complete(r)
}
