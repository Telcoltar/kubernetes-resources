package builders

import (
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// IngressPathBuilderI is the interface for building HTTP ingress paths
type IngressPathBuilderI interface {
	Build() networkingv1.HTTPIngressPath
}

// IngressPathBuilder builds an HTTPIngressPath
type IngressPathBuilder struct {
	path networkingv1.HTTPIngressPath
}

// IngressPath creates a new IngressPathBuilder with the given path (defaults to PathTypePrefix)
func IngressPath(path string) *IngressPathBuilder {
	pathType := networkingv1.PathTypePrefix
	return &IngressPathBuilder{
		path: networkingv1.HTTPIngressPath{
			Path:     path,
			PathType: &pathType,
		},
	}
}

// PathType sets the path type for the ingress path
func (b *IngressPathBuilder) PathType(pathType networkingv1.PathType) *IngressPathBuilder {
	b.path.PathType = &pathType
	return b
}

// PathTypeExact sets the path type to Exact
func (b *IngressPathBuilder) PathTypeExact() *IngressPathBuilder {
	return b.PathType(networkingv1.PathTypeExact)
}

// PathTypePrefix sets the path type to Prefix
func (b *IngressPathBuilder) PathTypePrefix() *IngressPathBuilder {
	return b.PathType(networkingv1.PathTypePrefix)
}

// PathTypeImplementationSpecific sets the path type to ImplementationSpecific
func (b *IngressPathBuilder) PathTypeImplementationSpecific() *IngressPathBuilder {
	return b.PathType(networkingv1.PathTypeImplementationSpecific)
}

// Backend sets the backend service with port number
func (b *IngressPathBuilder) Backend(serviceName string, port int32) *IngressPathBuilder {
	b.path.Backend = networkingv1.IngressBackend{
		Service: &networkingv1.IngressServiceBackend{
			Name: serviceName,
			Port: networkingv1.ServiceBackendPort{
				Number: port,
			},
		},
	}
	return b
}

// BackendWithPortName sets the backend service with named port
func (b *IngressPathBuilder) BackendWithPortName(serviceName, portName string) *IngressPathBuilder {
	b.path.Backend = networkingv1.IngressBackend{
		Service: &networkingv1.IngressServiceBackend{
			Name: serviceName,
			Port: networkingv1.ServiceBackendPort{
				Name: portName,
			},
		},
	}
	return b
}

// Build returns the HTTPIngressPath
func (b *IngressPathBuilder) Build() networkingv1.HTTPIngressPath {
	return b.path
}

// IngressRuleBuilderI is the interface for building ingress rules
type IngressRuleBuilderI interface {
	Build() networkingv1.IngressRule
}

// IngressRuleBuilder builds an IngressRule
type IngressRuleBuilder struct {
	rule networkingv1.IngressRule
}

// IngressRule creates a new IngressRuleBuilder with the given host
func IngressRule(host string) *IngressRuleBuilder {
	return &IngressRuleBuilder{
		rule: networkingv1.IngressRule{
			Host: host,
		},
	}
}

// Paths adds HTTP paths to the ingress rule
func (b *IngressRuleBuilder) Paths(paths ...IngressPathBuilderI) *IngressRuleBuilder {
	if b.rule.HTTP == nil {
		b.rule.HTTP = &networkingv1.HTTPIngressRuleValue{}
	}
	b.rule.HTTP.Paths = append(b.rule.HTTP.Paths, BuildAll(paths...)...)
	return b
}

// Build returns the IngressRule
func (b *IngressRuleBuilder) Build() networkingv1.IngressRule {
	return b.rule
}

// IngressBuilder builds a Kubernetes Ingress resource
type IngressBuilder struct {
	*objectMetaBuilder[IngressBuilder]
	spec  networkingv1.IngressSpec
	rules []IngressRuleBuilderI
}

// Ingress creates a new IngressBuilder with the given name
func Ingress(name string) *IngressBuilder {
	ib := &IngressBuilder{}
	ib.objectMetaBuilder = ObjectMeta(ib).name(name)
	return ib
}

// IngressClassName sets the ingress class name
func (b *IngressBuilder) IngressClassName(className string) *IngressBuilder {
	b.spec.IngressClassName = &className
	return b
}

// DefaultBackend sets the default backend with port number
func (b *IngressBuilder) DefaultBackend(serviceName string, port int32) *IngressBuilder {
	b.spec.DefaultBackend = &networkingv1.IngressBackend{
		Service: &networkingv1.IngressServiceBackend{
			Name: serviceName,
			Port: networkingv1.ServiceBackendPort{
				Number: port,
			},
		},
	}
	return b
}

// DefaultBackendWithPortName sets the default backend with named port
func (b *IngressBuilder) DefaultBackendWithPortName(serviceName, portName string) *IngressBuilder {
	b.spec.DefaultBackend = &networkingv1.IngressBackend{
		Service: &networkingv1.IngressServiceBackend{
			Name: serviceName,
			Port: networkingv1.ServiceBackendPort{
				Name: portName,
			},
		},
	}
	return b
}

// Rules adds ingress rules
func (b *IngressBuilder) Rules(rules ...IngressRuleBuilderI) *IngressBuilder {
	b.rules = append(b.rules, rules...)
	return b
}

// TLS adds TLS configuration
func (b *IngressBuilder) TLS(hosts []string, secretName string) *IngressBuilder {
	b.spec.TLS = append(b.spec.TLS, networkingv1.IngressTLS{
		Hosts:      hosts,
		SecretName: secretName,
	})
	return b
}

// Build returns the Ingress object
func (b *IngressBuilder) Build() (runtime.Object, error) {
	b.spec.Rules = append(b.spec.Rules, BuildAll(b.rules...)...)
	return &networkingv1.Ingress{
		TypeMeta:   ingressType,
		ObjectMeta: b.objectMetaBuilder.Build(),
		Spec:       b.spec,
	}, nil
}
