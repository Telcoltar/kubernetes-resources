package builders

import (
	"maps"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type SecretBuilder struct {
	*objectMetaBuilder[SecretBuilder]
	data map[string][]byte
}

func Secret(name string) *SecretBuilder {
	b := &SecretBuilder{
		data: map[string][]byte{},
	}
	b.objectMetaBuilder = ObjectMeta(b).name(name)
	return b
}

func (b *SecretBuilder) Data(data map[string][]byte) *SecretBuilder {
	maps.Copy(b.data, data)
	return b
}

func (b *SecretBuilder) Datum(key string, value []byte) *SecretBuilder {
	b.data[key] = value
	return b
}

func (b *SecretBuilder) StringData(data map[string]string) *SecretBuilder {
	for key, value := range data {
		b.data[key] = []byte(value)
	}
	return b
}

func (b *SecretBuilder) StringDatum(key, value string) *SecretBuilder {
	b.data[key] = []byte(value)
	return b
}

func (b *SecretBuilder) Build() (runtime.Object, error) {
	return &corev1.Secret{
		TypeMeta:   secretType,
		ObjectMeta: b.objectMetaBuilder.Build(),
		Data:       b.data,
	}, nil
}

func (b *SecretBuilder) Env(name string) *EnvKeyValuesBuilder {
	return Env(name).FromSecret(b.meta.Name)
}
