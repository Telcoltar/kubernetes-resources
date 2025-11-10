package builders

import (
	"maps"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type objectMetaBuilder[T any] struct {
	meta   metav1.ObjectMeta
	parent *T
}

func ObjectMeta[T any](parent *T) *objectMetaBuilder[T] {
	return &objectMetaBuilder[T]{parent: parent}
}

func (b *objectMetaBuilder[T]) name(name string) *objectMetaBuilder[T] {
	b.meta.Name = name
	return b
}

func (b *objectMetaBuilder[T]) GetName() string {
	return b.meta.Name
}

func (b *objectMetaBuilder[T]) SetLabels(labels map[string]string) {
	if b.meta.Labels == nil {
		b.meta.Labels = make(map[string]string)
	}
	maps.Copy(b.meta.Labels, labels)
}

func (b *objectMetaBuilder[T]) Labels(labels map[string]string) *T {
	b.SetLabels(labels)
	return b.parent
}

func (b *objectMetaBuilder[T]) SetLabel(key, value string) {
	b.SetLabels(map[string]string{key: value})
}

func (b *objectMetaBuilder[T]) Label(key, value string) *T {
	return b.Labels(map[string]string{key: value})
}

func (b *objectMetaBuilder[T]) Annotations(annotations map[string]string) *T {
	if b.meta.Annotations == nil {
		b.meta.Annotations = make(map[string]string)
	}
	maps.Copy(b.meta.Annotations, annotations)
	return b.parent
}

func (b *objectMetaBuilder[T]) Annotation(key, value string) *T {
	return b.Annotations(map[string]string{key: value})
}

func (b *objectMetaBuilder[T]) Namespace(namespace string) *T {
	b.SetNamespace(namespace)
	return b.parent
}

func (b *objectMetaBuilder[T]) SetNamespace(namespace string) {
	b.meta.Namespace = namespace
}

func (b *objectMetaBuilder[T]) Build() metav1.ObjectMeta {
	return b.meta
}
