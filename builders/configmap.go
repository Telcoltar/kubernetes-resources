package builders

import (
	"maps"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type ConfigmapBuilder struct {
	*objectMetaBuilder[ConfigmapBuilder]
	data map[string]string
}

func ConfigMap(name string) *ConfigmapBuilder {
	b := &ConfigmapBuilder{
		data: map[string]string{},
	}
	b.objectMetaBuilder = ObjectMeta(b).name(name)
	return b
}

func (b *ConfigmapBuilder) Data(data map[string]string) *ConfigmapBuilder {
	maps.Copy(b.data, data)
	return b
}

func (b *ConfigmapBuilder) Datum(key, value string) *ConfigmapBuilder {
	b.data[key] = value
	return b
}

func (b *ConfigmapBuilder) Build() (runtime.Object, error) {
	return &corev1.ConfigMap{
		TypeMeta:   configMapType,
		ObjectMeta: b.objectMetaBuilder.Build(),
		Data:       b.data,
	}, nil
}

func (b *ConfigmapBuilder) Env(name string) *EnvKeyValuesBuilder {
	return Env(name).FromConfigmap(b.meta.Name)
}
