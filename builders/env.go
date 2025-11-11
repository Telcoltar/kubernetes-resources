package builders

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"
)

type EnvBuilderI interface {
	Build() corev1.EnvVar
}

type envBaseBuilder struct {
	name          string
	actualBuilder EnvBuilderI
}

func Env(name string) *envBaseBuilder {
	return &envBaseBuilder{
		name: name,
	}
}

func (b *envBaseBuilder) Build() corev1.EnvVar {
	if b.actualBuilder != nil {
		return b.actualBuilder.Build()
	}
	return corev1.EnvVar{
		Name: b.name,
	}
}

type EnvValueBuilder struct {
	name  string
	value string
}

func (b *envBaseBuilder) Value(value string) *EnvValueBuilder {
	eb := &EnvValueBuilder{
		name:  b.name,
		value: value,
	}
	b.actualBuilder = eb
	return eb
}

func (b *EnvValueBuilder) Build() corev1.EnvVar {
	return corev1.EnvVar{
		Name:  b.name,
		Value: b.value,
	}
}

type EnvKeyValuesBuilder struct {
	name          string
	key           string
	optional      *bool
	valueFromFunc func(key string, optional *bool) *corev1.EnvVarSource
}

func (b *envBaseBuilder) FromSecret(name string) *EnvKeyValuesBuilder {
	return &EnvKeyValuesBuilder{
		name: b.name,
		valueFromFunc: func(key string, optional *bool) *corev1.EnvVarSource {
			return &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					Key:      key,
					Optional: optional,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: name,
					},
				},
			}
		},
	}
}

func (b *envBaseBuilder) FromConfigmap(name string) *EnvKeyValuesBuilder {
	return &EnvKeyValuesBuilder{
		name: b.name,
		valueFromFunc: func(key string, optional *bool) *corev1.EnvVarSource {
			return &corev1.EnvVarSource{
				ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
					Key:      key,
					Optional: optional,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: name,
					},
				},
			}
		},
	}
}

func (b *EnvKeyValuesBuilder) Key(key string) *EnvKeyValuesBuilder {
	b.key = key
	return b
}

func (b *EnvKeyValuesBuilder) Optional() *EnvKeyValuesBuilder {
	b.optional = ptr.To(true)
	return b
}

func (b *EnvKeyValuesBuilder) Build() corev1.EnvVar {
	return corev1.EnvVar{
		Name:      b.name,
		ValueFrom: b.valueFromFunc(b.key, b.optional),
	}
}

type EnvFieldRefBuilder struct {
	name string
	path string
}

func (b *envBaseBuilder) NodeName() *EnvFieldRefBuilder {
	return &EnvFieldRefBuilder{
		name: b.name,
		path: "spec.nodeName",
	}
}

func (b *envBaseBuilder) Namespace() *EnvFieldRefBuilder {
	return &EnvFieldRefBuilder{
		name: b.name,
		path: "metadata.namespace",
	}
}

func (b *envBaseBuilder) FieldRef(path string) *EnvFieldRefBuilder {
	return &EnvFieldRefBuilder{
		name: b.name,
		path: path,
	}
}

func (b *EnvFieldRefBuilder) Build() corev1.EnvVar {
	return corev1.EnvVar{
		Name: b.name,
		ValueFrom: &corev1.EnvVarSource{
			FieldRef: &corev1.ObjectFieldSelector{
				FieldPath: b.path,
			},
		},
	}
}
