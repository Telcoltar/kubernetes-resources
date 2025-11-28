package builders

import (
	"bytes"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
)

var s = json.NewSerializerWithOptions(
	json.DefaultMetaFactory, // Used to create objects during decoding
	scheme.Scheme,           // The scheme for type registration (The key part!)
	scheme.Scheme,           // The scheme for type creation
	json.SerializerOptions{
		Yaml: true, // We want JSON output, not YAML
		Pretty: true,
        Strict: true,
	},
)

type NamespaceBuilderI interface {
	SetNamespace(namespace string)
}

type LabelBuilderI interface {
	SetLabels(labels map[string]string)
	SetLabel(key, value string)
}

type StorageClassI interface {
	SetStorageClass(name string)
}

type Builder interface {
	Build() (runtime.Object, error)
}

func BuildYaml(builder Builder) ([]byte, error) {
	resource, err := builder.Build()
	if err != nil {
		return nil, err
	}
	var buffer bytes.Buffer
	err = s.Encode(resource, &buffer)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

type genericBuilder[T any] interface {
	Build() T
}

func BuildAll[B genericBuilder[T], T any](builders ...B) []T {
	result := make([]T, len(builders))
	for i, builder := range builders {
		result[i] = builder.Build()
	}
	return result
}
