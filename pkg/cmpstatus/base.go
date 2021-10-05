package cmpstatus

import (
	"context"
	"fmt"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	qv1 "github.com/quay/quay-operator/apis/quay/v1"
)

// Base checks a quay registry base component status. In order to evaluate the status for the
// base component we need to verify if quay and config-editor deployments succeed.
type Base struct {
	Client client.Client
	deploy deploy
}

// Name returns the component name this entity checks for health.
func (b *Base) Name() string {
	return "base"
}

// Check verifies if the quay and config-editor deployment associated with provided quay registry
// were created and rolled out as expected.
func (b *Base) Check(ctx context.Context, reg qv1.QuayRegistry) (qv1.Condition, error) {
	var zero qv1.Condition

	// we need to check two distinct deployments, the quay app and its config editor.
	for _, depsuffix := range []string{"quay-app", "quay-config-editor"} {
		depname := fmt.Sprintf("%s-%s", reg.Name, depsuffix)
		nsn := types.NamespacedName{
			Namespace: reg.Namespace,
			Name:      depname,
		}

		var dep appsv1.Deployment
		if err := b.Client.Get(ctx, nsn, &dep); err != nil {
			if errors.IsNotFound(err) {
				msg := fmt.Sprintf("Deployment %s not found", depname)
				return qv1.Condition{
					Type:           qv1.ComponentBaseReady,
					Status:         metav1.ConditionFalse,
					Reason:         qv1.ConditionReasonComponentNotReady,
					Message:        msg,
					LastUpdateTime: metav1.NewTime(time.Now()),
				}, nil
			}
			return zero, err
		}

		if !qv1.Owns(reg, &dep) {
			msg := fmt.Sprintf("Deployment %s not owned by QuayRegistry", depname)
			return qv1.Condition{
				Type:           qv1.ComponentBaseReady,
				Status:         metav1.ConditionFalse,
				Reason:         qv1.ConditionReasonComponentNotReady,
				Message:        msg,
				LastUpdateTime: metav1.NewTime(time.Now()),
			}, nil
		}

		cond := b.deploy.check(dep)
		if cond.Status != metav1.ConditionTrue {
			// if the deployment is in a faulty state bails out immediately.
			cond.Type = qv1.ComponentBaseReady
			return cond, nil
		}
	}

	return qv1.Condition{
		Type:           qv1.ComponentBaseReady,
		Reason:         qv1.ConditionReasonComponentReady,
		Status:         metav1.ConditionTrue,
		Message:        "Base component healthy",
		LastUpdateTime: metav1.NewTime(time.Now()),
	}, nil
}
