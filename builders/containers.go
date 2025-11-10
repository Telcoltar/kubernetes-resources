package builders

import corev1 "k8s.io/api/core/v1"

type ContainersBuilder struct {
	containers []*ContainerBuilder
}

func Containers() *ContainersBuilder {
	return &ContainersBuilder{}
}

func (b *ContainersBuilder) Containers(containers ...*ContainerBuilder) {
	b.containers = append(b.containers, containers...)
}

func (b *ContainersBuilder) Build() []corev1.Container {
	return BuildAll(b.containers...)
}
