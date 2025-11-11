package builders

import (
	"maps"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type ConfigMapBuilder struct {
	*objectMetaBuilder[ConfigMapBuilder]
	data map[string]string
}

func ConfigMap(name string) *ConfigMapBuilder {
	b := &ConfigMapBuilder{
		data: map[string]string{},
	}
	b.objectMetaBuilder = ObjectMeta(b).name(name)
	return b
}

func (b *ConfigMapBuilder) Data(data map[string]string) *ConfigMapBuilder {
	maps.Copy(b.data, data)
	return b
}

func (b *ConfigMapBuilder) Datum(key, value string) *ConfigMapBuilder {
	b.data[key] = value
	return b
}

func (b *ConfigMapBuilder) Build() (runtime.Object, error) {
	return &corev1.ConfigMap{
		TypeMeta:   configMapType,
		ObjectMeta: b.objectMetaBuilder.Build(),
		Data:       b.data,
	}, nil
}

func (b *ConfigMapBuilder) Env(name string) *EnvKeyValuesBuilder {
	return Env(name).FromConfigmap(b.meta.Name)
}

func (b *ConfigMapBuilder) Volume() *VolumeKeyPathBuilder {
	return Volume(b.meta.Name).ConfigMap(b.meta.Name)
}
