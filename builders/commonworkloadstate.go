package builders

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type commonWorkloadBuilder[P any] struct {
	*objectMetaBuilder[P]
	*PodTemplateSpecBuilder[P]

	replicas *int32
	selector *metav1.LabelSelector
	minReady int32

	parent *P
}

func commonWorkloadState[P any](name string, parent *P) *commonWorkloadBuilder[P] {
	b := &commonWorkloadBuilder[P]{
		parent: parent,
	}
	b.objectMetaBuilder = ObjectMeta(parent).name(name)
	b.PodTemplateSpecBuilder = PodTempalteSpec(parent)
	return b
}

func (b *commonWorkloadBuilder[P]) SetLabel(key, value string) {
	b.objectMetaBuilder.SetLabel(key, value)
	b.PodTemplateSpecBuilder.SetLabel(key, value)
}

func (b *commonWorkloadBuilder[P]) SetLabels(labels map[string]string) {
	b.objectMetaBuilder.SetLabels(labels)
	b.PodTemplateSpecBuilder.SetLabels(labels)
}

func (b *commonWorkloadBuilder[P]) Label(key, value string) *P {
	b.SetLabel(key, value)
	return b.parent
}

func (b *commonWorkloadBuilder[P]) Labels(labels map[string]string) *P {
	b.SetLabels(labels)
	return b.parent
}

func (b *commonWorkloadBuilder[P]) Replicas(replicas int32) *P {
	b.replicas = &replicas
	return b.parent
}

func (b *commonWorkloadBuilder[P]) Selector(selectors labels.Set) *P {
	b.selector = metav1.SetAsLabelSelector(selectors)
	return b.parent
}

func (b *commonWorkloadBuilder[P]) MinReady(seconds int32) *P {
	b.minReady = seconds
	return b.parent
}
