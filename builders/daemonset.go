package builders

import (
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DaemonSetBuilder builds a DaemonSet resource using the common deployment state helpers.
type DaemonSetBuilder struct {
	*commonWorkloadBuilder[DaemonSetBuilder]
	spec appsv1.DaemonSetSpec
}

// DaemonSet starts a new DaemonSetBuilder.
func DaemonSet(name string) *DaemonSetBuilder {
	b := &DaemonSetBuilder{
		spec: appsv1.DaemonSetSpec{},
	}
	b.commonWorkloadBuilder = commonWorkloadState(name, b)
	return b
}

// UpdateStrategy sets the DaemonSet update strategy.
func (b *DaemonSetBuilder) UpdateStrategy(strategy appsv1.DaemonSetUpdateStrategy) *DaemonSetBuilder {
	b.spec.UpdateStrategy = strategy
	return b
}

// RevisionHistoryLimit sets the number of old history revisions to keep.
func (b *DaemonSetBuilder) RevisionHistoryLimit(limit int32) *DaemonSetBuilder {
	b.spec.RevisionHistoryLimit = &limit
	return b
}

func (b *DaemonSetBuilder) Build() (runtime.Object, error) {
	b.spec.Template = b.PodTemplateSpecBuilder.Build()
	b.spec.Selector = b.selector
	b.spec.MinReadySeconds = b.minReady

	return &appsv1.DaemonSet{
		TypeMeta:   daemonSetType,
		ObjectMeta: b.objectMetaBuilder.Build(),
		Spec:       b.spec,
	}, nil
}
