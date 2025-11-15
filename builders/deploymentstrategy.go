package builders

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type DeploymentStrategyBuilderI interface {
	Build() appsv1.DeploymentStrategy
}

type DeploymentStrategyBuilder struct{}

func DeploymentStrategy() *DeploymentStrategyBuilder {
	return &DeploymentStrategyBuilder{}
}

type DeploymentStrategyRecreateBuilder struct{}

// Recreate sets the deployment strategy to Recreate
func (b *DeploymentStrategyBuilder) Recreate() *DeploymentStrategyRecreateBuilder {
	return &DeploymentStrategyRecreateBuilder{}
}

func (b *DeploymentStrategyRecreateBuilder) Build() appsv1.DeploymentStrategy {
	return appsv1.DeploymentStrategy{
		Type: appsv1.RecreateDeploymentStrategyType,
	}
}

// RollingUpdate returns a RollingUpdateStrategyBuilder for configuring rolling update strategy
func (b *DeploymentStrategyBuilder) RollingUpdate() *RollingUpdateStrategyBuilder {
	return &RollingUpdateStrategyBuilder{
		strategy: appsv1.DeploymentStrategy{
			Type:          appsv1.RollingUpdateDeploymentStrategyType,
			RollingUpdate: &appsv1.RollingUpdateDeployment{},
		},
	}
}

type RollingUpdateStrategyBuilder struct {
	strategy appsv1.DeploymentStrategy
}

// MaxUnavailable sets the maximum number of pods that can be unavailable during the update (absolute number)
func (b *RollingUpdateStrategyBuilder) MaxUnavailable(value int32) *RollingUpdateStrategyBuilder {
	intOrStr := intstr.FromInt32(value)
	b.strategy.RollingUpdate.MaxUnavailable = &intOrStr
	return b
}

// MaxUnavailablePercent sets the maximum number of pods that can be unavailable during the update (percentage)
func (b *RollingUpdateStrategyBuilder) MaxUnavailablePercent(percent int32) *RollingUpdateStrategyBuilder {
	intOrStr := intstr.FromString(fmt.Sprintf("%d%%", percent))
	b.strategy.RollingUpdate.MaxUnavailable = &intOrStr
	return b
}

// MaxSurge sets the maximum number of pods that can be scheduled above the desired number (absolute number)
func (b *RollingUpdateStrategyBuilder) MaxSurge(value int32) *RollingUpdateStrategyBuilder {
	intOrStr := intstr.FromInt32(value)
	b.strategy.RollingUpdate.MaxSurge = &intOrStr
	return b
}

// MaxSurgePercent sets the maximum number of pods that can be scheduled above the desired number (percentage)
func (b *RollingUpdateStrategyBuilder) MaxSurgePercent(percent int32) *RollingUpdateStrategyBuilder {
	intOrStr := intstr.FromString(fmt.Sprintf("%d%%", percent))
	b.strategy.RollingUpdate.MaxSurge = &intOrStr
	return b
}

func (b *RollingUpdateStrategyBuilder) Build() appsv1.DeploymentStrategy {
	return b.strategy
}
