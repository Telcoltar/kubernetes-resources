package builders

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type ServicePortBuilderI interface {
	Build() corev1.ServicePort
}

type basePortBuilder[P any] struct {
	port corev1.ServicePort

	parent *P
}

func basePort[P any](parent *P) *basePortBuilder[P] {
	return &basePortBuilder[P]{
		port:   corev1.ServicePort{},
		parent: parent,
	}
}

func (b *basePortBuilder[P]) Name(name string) *P {
	b.port.Name = name
	return b.parent
}

func (b *basePortBuilder[P]) Port(port int32) *P {
	b.port.Port = port
	return b.parent
}

func (b *basePortBuilder[P]) NodePort(port int32) *P {
	b.port.NodePort = port
	return b.parent
}

func (b *basePortBuilder[P]) targetPort(port int32) *P {
	b.port.TargetPort = intstr.FromInt32(port)
	return b.parent
}

func (b *basePortBuilder[P]) targetPortName(name string) *P {
	b.port.TargetPort = intstr.FromString(name)
	return b.parent
}

func (b *basePortBuilder[P]) Build() corev1.ServicePort {
	return b.port
}

type ServicePortBuilder struct {
	*basePortBuilder[ServicePortBuilder]
}

func ServicePort() *ServicePortBuilder {
	b := &ServicePortBuilder{}
	b.basePortBuilder = basePort(b)
	return b
}

func (b *ServicePortBuilder) TargetPort(port int32) *ServicePortBuilder {
	return b.targetPort(port)
}

func (b *ServicePortBuilder) TargetPortName(name string) *ServicePortBuilder {
	return b.targetPortName(name)
}

type DerivedServicePortBuilder struct {
	containerPort *ContainerPortBuilder
	*basePortBuilder[DerivedServicePortBuilder]
}

func derivedServicePort(port *ContainerPortBuilder) *DerivedServicePortBuilder {
	b := &DerivedServicePortBuilder{
		containerPort: port,
	}
	b.basePortBuilder = basePort(b)
	return b
}

func (b *DerivedServicePortBuilder) Build() corev1.ServicePort {
	if b.containerPort.port.Name != "" {
		b.targetPortName(b.containerPort.port.Name)
	} else {
		b.targetPort(b.containerPort.port.ContainerPort)
	}
	if b.port.Port == 0 {
		b.Port(b.containerPort.port.ContainerPort)
	}
	return b.basePortBuilder.Build()
}
