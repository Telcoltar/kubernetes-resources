package builders

import (
	"maps"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type ServiceBuilder struct {
	*objectMetaBuilder[ServiceBuilder]
	spec  corev1.ServiceSpec
	ports []ServicePortBuilderI
}

func Service(name string) *ServiceBuilder {
	sb := &ServiceBuilder{}
	sb.objectMetaBuilder = ObjectMeta(sb).name(name)
	return sb
}

func (b *ServiceBuilder) ClusterIP(clusterIP string) *ServiceBuilder {
	b.spec.ClusterIP = clusterIP
	return b
}

func (b *ServiceBuilder) PublishNotReadyAdress() *ServiceBuilder {
	b.spec.PublishNotReadyAddresses = true
	return b
}

func (b *ServiceBuilder) Selectors(labels map[string]string) *ServiceBuilder {
	if b.spec.Selector == nil {
		b.spec.Selector = make(map[string]string)
	}
	maps.Copy(b.spec.Selector, labels)
	return b
}

func (b *ServiceBuilder) Selector(key, value string) *ServiceBuilder {
	return b.Selectors(map[string]string{key: value})
}

func (b *ServiceBuilder) Ports(ports ...ServicePortBuilderI) *ServiceBuilder {
	b.ports = append(b.ports, ports...)
	return b
}

func (b *ServiceBuilder) Type(serviceType corev1.ServiceType) *ServiceBuilder {
	b.spec.Type = serviceType
	return b
}

func (b *ServiceBuilder) LoadbalancerIP(ip string) *ServiceBuilder {
	b.spec.LoadBalancerIP = ip
	return b
}

func (b *ServiceBuilder) Build() (runtime.Object, error) {
	b.spec.Ports = append(b.spec.Ports, BuildAll(b.ports...)...)
	return &corev1.Service{
		TypeMeta:   serviceType,
		ObjectMeta: b.objectMetaBuilder.Build(),
		Spec:       b.spec,
	}, nil
}
