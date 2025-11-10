package builders

import (
	"bytes"
)

type Builders []Builder

func (b Builders) Namespace(namespace string) Builders {
	for _, builder := range b {
		if namespaceBuilder, ok := builder.(NamespaceBuilderI); ok {
			namespaceBuilder.SetNamespace(namespace)
		}
	}
	return b
}

func (b Builders) Labels(labels map[string]string) Builders {
	for _, builder := range b {
		if labelBuilder, ok := builder.(LabelBuilderI); ok {
			labelBuilder.SetLabels(labels)
		}
	}
	return b
}

func (b Builders) Label(key, value string) Builders {
	for _, builder := range b {
		if labelBuilder, ok := builder.(LabelBuilderI); ok {
			labelBuilder.SetLabel(key, value)
		}
	}
	return b
}

func (b Builders) StorageClass(name string) Builders {
	for _, builder := range b {
		if storageClassBuilder, ok := builder.(StorageClassI); ok {
			storageClassBuilder.SetStorageClass(name)
		}
	}
	return b
}

func (b Builders) AsYaml() ([]byte, error) {
	yamls := make([][]byte, len(b))
	for i, builder := range b {
		if resource, err := BuildYaml(builder); err != nil {
			return nil, err
		} else {
			yamls[i] = resource
		}
	}

	return bytes.Join(yamls, []byte("\n---\n")), nil
}
