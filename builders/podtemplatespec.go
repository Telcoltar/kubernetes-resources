package builders

import corev1 "k8s.io/api/core/v1"

type PodTemplateSpecBuilder[P any] struct {
	*objectMetaBuilder[P]
	*PodSpecBuilder[P]
}

func PodTempalteSpec[P any](parent *P) *PodTemplateSpecBuilder[P] {
	b := &PodTemplateSpecBuilder[P]{
		objectMetaBuilder: ObjectMeta(parent),
		PodSpecBuilder:    PodSpec(parent),
	}
	return b
}

func (b *PodTemplateSpecBuilder[P]) PodLabel(key, value string) *P {
	return b.Label(key, value)
}

func (b *PodTemplateSpecBuilder[P]) PodLabels(labels map[string]string) *P {
	return b.Labels(labels)
}

func (b *PodTemplateSpecBuilder[P]) PodAnnotation(key, value string) *P {
	return b.Annotation(key, value)
}

func (b *PodTemplateSpecBuilder[P]) PodAnnotations(annotations map[string]string) *P {
	return b.Annotations(annotations)
}

func (b *PodTemplateSpecBuilder[P]) Build() corev1.PodTemplateSpec {
	return corev1.PodTemplateSpec{
		ObjectMeta: b.objectMetaBuilder.Build(),
		Spec:       b.PodSpecBuilder.Build(),
	}
}
