package builders

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type commonDeployStateBuilder[P any] struct {
	*objectMetaBuilder[P]
	*PodTemplateSpecBuilder[P]

	replicas *int32
	selector *metav1.LabelSelector
	minReady int32

	parent *P
}

func commonDeployState[P any](name string, parent *P) *commonDeployStateBuilder[P] {
	b := &commonDeployStateBuilder[P]{
		parent: parent,
	}
	b.objectMetaBuilder = ObjectMeta(parent).name(name)
	b.PodTemplateSpecBuilder = PodTempalteSpec(parent)
	return b
}

func (b *commonDeployStateBuilder[P]) SetLabel(key, value string) {
	b.objectMetaBuilder.SetLabel(key, value)
	b.PodTemplateSpecBuilder.SetLabel(key, value)
}

func (b *commonDeployStateBuilder[P]) SetLabels(labels map[string]string) {
	b.objectMetaBuilder.SetLabels(labels)
	b.PodTemplateSpecBuilder.SetLabels(labels)
}

func (b *commonDeployStateBuilder[P]) Label(key, value string) *P {
	b.SetLabel(key, value)
	return b.parent
}

func (b *commonDeployStateBuilder[P]) Labels(labels map[string]string) *P {
	b.SetLabels(labels)
	return b.parent
}

func (b *commonDeployStateBuilder[P]) Replicas(replicas int32) *P {
	b.replicas = &replicas
	return b.parent
}

func (b *commonDeployStateBuilder[P]) Selector(selectors labels.Set) *P {
	b.selector = metav1.SetAsLabelSelector(selectors)
	return b.parent
}

func (b *commonDeployStateBuilder[P]) MinReady(seconds int32) *P {
	b.minReady = seconds
	return b.parent
}
