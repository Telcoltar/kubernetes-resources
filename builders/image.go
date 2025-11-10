package builders

import "strings"

type ImageBuilder struct {
	repository string
	registry   *string
	tag        *string
}

func Image(repository string) *ImageBuilder {
	return &ImageBuilder{repository: repository}
}

func (b *ImageBuilder) Registry(registry string) *ImageBuilder {
	b.registry = &registry
	return b
}

func (b *ImageBuilder) Tag(tag string) *ImageBuilder {
	b.tag = &tag
	return b
}

func (b *ImageBuilder) Build() string {
	var sb strings.Builder
	if b.registry != nil && *b.registry != "" {
		sb.WriteString(*b.registry)
		sb.WriteString("/")
	}
	sb.WriteString(b.repository)
	tag := "latest"
	if b.tag != nil {
		tag = *b.tag
	}
	if tag != "" {
		sb.WriteString(":")
		sb.WriteString(tag)
	}
	return sb.String()
}
