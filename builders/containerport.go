package builders

import corev1 "k8s.io/api/core/v1"

type ContainerPortBuilderI interface {
	Build() corev1.ContainerPort
}

type baseContainerPortBuilder[P any] struct {
	port corev1.ContainerPort

	parent *P
}

func baseContainerPort[P any](parent *P) *baseContainerPortBuilder[P] {
	return &baseContainerPortBuilder[P]{
		port:   corev1.ContainerPort{},
		parent: parent,
	}
}

func (b *baseContainerPortBuilder[P]) Name(name string) *P {
	b.port.Name = name
	return b.parent
}

func (b *baseContainerPortBuilder[P]) Build() corev1.ContainerPort {
	return b.port
}

type ContainerPortBuilder struct {
	*baseContainerPortBuilder[ContainerPortBuilder]
}

func ContainerPort() *ContainerPortBuilder {
	b := &ContainerPortBuilder{}
	b.baseContainerPortBuilder = baseContainerPort(b)
	return b
}

func (b *ContainerPortBuilder) Port(port int32) *ContainerPortBuilder {
	b.port.ContainerPort = port
	return b
}

func (b *ContainerPortBuilder) Protocol(protocol corev1.Protocol) *ContainerPortBuilder {
	b.port.Protocol = protocol
	return b
}

func (b *ContainerPortBuilder) HostPort(port int32) *ContainerPortBuilder {
	b.port.HostPort = port
	return b
}

func (b *ContainerPortBuilder) HostIP(ip string) *ContainerPortBuilder {
	b.port.HostIP = ip
	return b
}

func (b *ContainerPortBuilder) Service() *DerivedServicePortBuilder {
	return derivedServicePort(b)
}
