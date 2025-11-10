package builders

import (
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type StatefulSetBuilder struct {
	*commonDeployStateBuilder[StatefulSetBuilder]
	pvcTemplates []*PVCTemplateBuilder
	spec         appsv1.StatefulSetSpec
}

func StatefulSet(name string) *StatefulSetBuilder {
	b := &StatefulSetBuilder{
		spec: appsv1.StatefulSetSpec{},
	}
	b.commonDeployStateBuilder = commonDeployState(name, b)
	return b
}

func (b *StatefulSetBuilder) ServiceName(name string) *StatefulSetBuilder {
	b.spec.ServiceName = name
	return b
}

func (b *StatefulSetBuilder) PodManagementPolicy(policy appsv1.PodManagementPolicyType) *StatefulSetBuilder {
	b.spec.PodManagementPolicy = policy
	return b
}

func (b *StatefulSetBuilder) ClaimTepmplates(claims ...*PVCTemplateBuilder) *StatefulSetBuilder {
	b.pvcTemplates = append(b.pvcTemplates, claims...)
	return b
}

func (b *StatefulSetBuilder) SetStorageClass(name string) {
	for _, pvcTemplate := range b.pvcTemplates {
		pvcTemplate.SetStorageClass(name)
	}
}

func (b *StatefulSetBuilder) Build() (runtime.Object, error) {
	b.spec.Template = b.PodTemplateSpecBuilder.Build()
	b.spec.Replicas = b.replicas
	b.spec.Selector = b.selector
	b.spec.MinReadySeconds = b.minReady
	for _, pvcTemplate := range b.pvcTemplates {
		pvcTemplate.Labels(b.meta.Labels)
	}
	b.spec.VolumeClaimTemplates = append(b.spec.VolumeClaimTemplates, BuildAll(b.pvcTemplates...)...)
	return &appsv1.StatefulSet{
		TypeMeta:   statefulSetType,
		ObjectMeta: b.objectMetaBuilder.Build(),
		Spec:       b.spec,
	}, nil
}
