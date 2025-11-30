package builders

import (
	routev1 "github.com/openshift/api/route/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type RouteBuilder struct {
	*objectMetaBuilder[RouteBuilder]
	spec routev1.RouteSpec
}

func Route(name string) *RouteBuilder {
	rb := &RouteBuilder{}
	rb.objectMetaBuilder = ObjectMeta(rb).name(name)
	return rb
}

func (b *RouteBuilder) Host(host string) *RouteBuilder {
	b.spec.Host = host
	return b
}

func (b *RouteBuilder) Path(path string) *RouteBuilder {
	b.spec.Path = path
	return b
}

func (b *RouteBuilder) To(serviceName string) *RouteBuilder {
	b.spec.To = routev1.RouteTargetReference{
		Kind: "Service",
		Name: serviceName,
	}
	return b
}

func (b *RouteBuilder) ToWithWeight(serviceName string, weight int32) *RouteBuilder {
	b.spec.To = routev1.RouteTargetReference{
		Kind:   "Service",
		Name:   serviceName,
		Weight: &weight,
	}
	return b
}

func (b *RouteBuilder) port(port intstr.IntOrString) *RouteBuilder {
	b.spec.Port = &routev1.RoutePort{
		TargetPort: port,
	}
	return b
}

func (b *RouteBuilder) PortName(portName string) *RouteBuilder {
	return b.port(intstr.FromString(portName))
}

func (b *RouteBuilder) Port(number int32) *RouteBuilder {
	return b.port(intstr.FromInt32(number))
}

func (b *RouteBuilder) TLS(tls *routev1.TLSConfig) *RouteBuilder {
	b.spec.TLS = tls
	return b
}

func (b *RouteBuilder) TLSEdge() *RouteBuilder {
	b.spec.TLS = &routev1.TLSConfig{
		Termination: routev1.TLSTerminationEdge,
	}
	return b
}

func (b *RouteBuilder) TLSPassthrough() *RouteBuilder {
	b.spec.TLS = &routev1.TLSConfig{
		Termination: routev1.TLSTerminationPassthrough,
	}
	return b
}

func (b *RouteBuilder) TLSReencrypt() *RouteBuilder {
	b.spec.TLS = &routev1.TLSConfig{
		Termination: routev1.TLSTerminationReencrypt,
	}
	return b
}

func (b *RouteBuilder) InsecureEdgeTerminationPolicy(policy routev1.InsecureEdgeTerminationPolicyType) *RouteBuilder {
	if b.spec.TLS == nil {
		b.spec.TLS = &routev1.TLSConfig{}
	}
	b.spec.TLS.InsecureEdgeTerminationPolicy = policy
	return b
}

func (b *RouteBuilder) Certificate(cert string) *RouteBuilder {
	if b.spec.TLS == nil {
		b.spec.TLS = &routev1.TLSConfig{}
	}
	b.spec.TLS.Certificate = cert
	return b
}

func (b *RouteBuilder) Key(key string) *RouteBuilder {
	if b.spec.TLS == nil {
		b.spec.TLS = &routev1.TLSConfig{}
	}
	b.spec.TLS.Key = key
	return b
}

func (b *RouteBuilder) CACertificate(caCert string) *RouteBuilder {
	if b.spec.TLS == nil {
		b.spec.TLS = &routev1.TLSConfig{}
	}
	b.spec.TLS.CACertificate = caCert
	return b
}

func (b *RouteBuilder) DestinationCACertificate(destCACert string) *RouteBuilder {
	if b.spec.TLS == nil {
		b.spec.TLS = &routev1.TLSConfig{}
	}
	b.spec.TLS.DestinationCACertificate = destCACert
	return b
}

func (b *RouteBuilder) AlternateBackends(backends ...routev1.RouteTargetReference) *RouteBuilder {
	b.spec.AlternateBackends = append(b.spec.AlternateBackends, backends...)
	return b
}

func (b *RouteBuilder) WildcardPolicy(policy routev1.WildcardPolicyType) *RouteBuilder {
	b.spec.WildcardPolicy = policy
	return b
}

func (b *RouteBuilder) Build() (runtime.Object, error) {
	return &routev1.Route{
		TypeMeta:   routeType,
		ObjectMeta: b.objectMetaBuilder.Build(),
		Spec:       b.spec,
	}, nil
}
