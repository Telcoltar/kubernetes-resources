package builders

import (
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type DeploymentBuilder struct {
	*commonWorkloadBuilder[DeploymentBuilder]
	spec     appsv1.DeploymentSpec
	strategy DeploymentStrategyBuilderI
}

func Deployment(name string) *DeploymentBuilder {
	b := &DeploymentBuilder{
		spec: appsv1.DeploymentSpec{},
	}
	b.commonWorkloadBuilder = commonWorkloadState(name, b)
	return b
}

func (b *DeploymentBuilder) Build() (runtime.Object, error) {
	b.spec.Template = b.PodTemplateSpecBuilder.Build()
	b.spec.Replicas = b.replicas
	b.spec.Selector = b.selector
	b.spec.MinReadySeconds = b.minReady
	if b.strategy != nil {
		strategy := b.strategy.Build()
		b.spec.Strategy = strategy
	}
	return &appsv1.Deployment{
		TypeMeta:   deploymentType,
		ObjectMeta: b.objectMetaBuilder.Build(),
		Spec:       b.spec,
	}, nil
}

func (b *DeploymentBuilder) Strategy(strategy DeploymentStrategyBuilderI) *DeploymentBuilder {
	b.strategy = strategy
	return b
}
