package builders

import (
	"testing"

	networkingv1 "k8s.io/api/networking/v1"
)

func TestIngressPath(t *testing.T) {
	path := IngressPath("/api").
		PathTypePrefix().
		Backend("api-service", 8080).
		Build()

	if path.Path != "/api" {
		t.Errorf("expected path /api, got %s", path.Path)
	}
	if *path.PathType != networkingv1.PathTypePrefix {
		t.Errorf("expected PathTypePrefix, got %v", *path.PathType)
	}
	if path.Backend.Service.Name != "api-service" {
		t.Errorf("expected service name api-service, got %s", path.Backend.Service.Name)
	}
	if path.Backend.Service.Port.Number != 8080 {
		t.Errorf("expected port 8080, got %d", path.Backend.Service.Port.Number)
	}
}

func TestIngressPathWithPortName(t *testing.T) {
	path := IngressPath("/web").
		PathTypeExact().
		BackendWithPortName("web-service", "http").
		Build()

	if path.Path != "/web" {
		t.Errorf("expected path /web, got %s", path.Path)
	}
	if *path.PathType != networkingv1.PathTypeExact {
		t.Errorf("expected PathTypeExact, got %v", *path.PathType)
	}
	if path.Backend.Service.Name != "web-service" {
		t.Errorf("expected service name web-service, got %s", path.Backend.Service.Name)
	}
	if path.Backend.Service.Port.Name != "http" {
		t.Errorf("expected port name http, got %s", path.Backend.Service.Port.Name)
	}
}

func TestIngressPathImplementationSpecific(t *testing.T) {
	path := IngressPath("/custom").
		PathTypeImplementationSpecific().
		Backend("custom-service", 9000).
		Build()

	if *path.PathType != networkingv1.PathTypeImplementationSpecific {
		t.Errorf("expected PathTypeImplementationSpecific, got %v", *path.PathType)
	}
}

func TestIngressRule(t *testing.T) {
	rule := IngressRule("example.com").
		Paths(
			IngressPath("/api").Backend("api-service", 8080),
			IngressPath("/web").Backend("web-service", 80),
		).
		Build()

	if rule.Host != "example.com" {
		t.Errorf("expected host example.com, got %s", rule.Host)
	}
	if len(rule.HTTP.Paths) != 2 {
		t.Errorf("expected 2 paths, got %d", len(rule.HTTP.Paths))
	}
	if rule.HTTP.Paths[0].Path != "/api" {
		t.Errorf("expected first path /api, got %s", rule.HTTP.Paths[0].Path)
	}
	if rule.HTTP.Paths[1].Path != "/web" {
		t.Errorf("expected second path /web, got %s", rule.HTTP.Paths[1].Path)
	}
}

func TestIngressBuilder(t *testing.T) {
	obj, err := Ingress("my-ingress").
		Namespace("default").
		Label("app", "myapp").
		IngressClassName("nginx").
		TLS([]string{"example.com"}, "tls-secret").
		Rules(
			IngressRule("example.com").
				Paths(
					IngressPath("/api").
						PathTypePrefix().
						Backend("api-service", 8080),
					IngressPath("/web").
						PathTypePrefix().
						Backend("web-service", 80),
				),
		).
		Build()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ingress, ok := obj.(*networkingv1.Ingress)
	if !ok {
		t.Fatalf("expected *networkingv1.Ingress, got %T", obj)
	}

	// Check metadata
	if ingress.Name != "my-ingress" {
		t.Errorf("expected name my-ingress, got %s", ingress.Name)
	}
	if ingress.Namespace != "default" {
		t.Errorf("expected namespace default, got %s", ingress.Namespace)
	}
	if ingress.Labels["app"] != "myapp" {
		t.Errorf("expected label app=myapp, got %s", ingress.Labels["app"])
	}

	// Check TypeMeta
	if ingress.Kind != "Ingress" {
		t.Errorf("expected Kind Ingress, got %s", ingress.Kind)
	}
	if ingress.APIVersion != "networking.k8s.io/v1" {
		t.Errorf("expected APIVersion networking.k8s.io/v1, got %s", ingress.APIVersion)
	}

	// Check spec
	if *ingress.Spec.IngressClassName != "nginx" {
		t.Errorf("expected ingress class nginx, got %s", *ingress.Spec.IngressClassName)
	}

	// Check TLS
	if len(ingress.Spec.TLS) != 1 {
		t.Fatalf("expected 1 TLS config, got %d", len(ingress.Spec.TLS))
	}
	if ingress.Spec.TLS[0].SecretName != "tls-secret" {
		t.Errorf("expected secret name tls-secret, got %s", ingress.Spec.TLS[0].SecretName)
	}
	if len(ingress.Spec.TLS[0].Hosts) != 1 || ingress.Spec.TLS[0].Hosts[0] != "example.com" {
		t.Errorf("expected TLS host example.com, got %v", ingress.Spec.TLS[0].Hosts)
	}

	// Check rules
	if len(ingress.Spec.Rules) != 1 {
		t.Fatalf("expected 1 rule, got %d", len(ingress.Spec.Rules))
	}
	if ingress.Spec.Rules[0].Host != "example.com" {
		t.Errorf("expected rule host example.com, got %s", ingress.Spec.Rules[0].Host)
	}
	if len(ingress.Spec.Rules[0].HTTP.Paths) != 2 {
		t.Errorf("expected 2 paths in rule, got %d", len(ingress.Spec.Rules[0].HTTP.Paths))
	}
}

func TestIngressDefaultBackend(t *testing.T) {
	obj, err := Ingress("default-backend-ingress").
		DefaultBackend("default-service", 80).
		Build()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ingress := obj.(*networkingv1.Ingress)

	if ingress.Spec.DefaultBackend == nil {
		t.Fatal("expected default backend to be set")
	}
	if ingress.Spec.DefaultBackend.Service.Name != "default-service" {
		t.Errorf("expected default backend service default-service, got %s", ingress.Spec.DefaultBackend.Service.Name)
	}
	if ingress.Spec.DefaultBackend.Service.Port.Number != 80 {
		t.Errorf("expected default backend port 80, got %d", ingress.Spec.DefaultBackend.Service.Port.Number)
	}
}

func TestIngressDefaultBackendWithPortName(t *testing.T) {
	obj, err := Ingress("default-backend-port-name-ingress").
		DefaultBackendWithPortName("default-service", "http").
		Build()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ingress := obj.(*networkingv1.Ingress)

	if ingress.Spec.DefaultBackend == nil {
		t.Fatal("expected default backend to be set")
	}
	if ingress.Spec.DefaultBackend.Service.Name != "default-service" {
		t.Errorf("expected default backend service default-service, got %s", ingress.Spec.DefaultBackend.Service.Name)
	}
	if ingress.Spec.DefaultBackend.Service.Port.Name != "http" {
		t.Errorf("expected default backend port name http, got %s", ingress.Spec.DefaultBackend.Service.Port.Name)
	}
}

func TestIngressMultipleTLS(t *testing.T) {
	obj, err := Ingress("multi-tls-ingress").
		TLS([]string{"example.com"}, "secret1").
		TLS([]string{"api.example.com", "admin.example.com"}, "secret2").
		Build()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ingress := obj.(*networkingv1.Ingress)

	if len(ingress.Spec.TLS) != 2 {
		t.Fatalf("expected 2 TLS configs, got %d", len(ingress.Spec.TLS))
	}
	if ingress.Spec.TLS[0].SecretName != "secret1" {
		t.Errorf("expected first TLS secret secret1, got %s", ingress.Spec.TLS[0].SecretName)
	}
	if ingress.Spec.TLS[1].SecretName != "secret2" {
		t.Errorf("expected second TLS secret secret2, got %s", ingress.Spec.TLS[1].SecretName)
	}
	if len(ingress.Spec.TLS[1].Hosts) != 2 {
		t.Errorf("expected 2 hosts in second TLS, got %d", len(ingress.Spec.TLS[1].Hosts))
	}
}

func TestIngressMultipleRules(t *testing.T) {
	obj, err := Ingress("multi-rule-ingress").
		Rules(
			IngressRule("example.com").
				Paths(IngressPath("/").Backend("main-service", 80)),
			IngressRule("api.example.com").
				Paths(IngressPath("/v1").Backend("api-v1", 8080)),
		).
		Build()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ingress := obj.(*networkingv1.Ingress)

	if len(ingress.Spec.Rules) != 2 {
		t.Fatalf("expected 2 rules, got %d", len(ingress.Spec.Rules))
	}
	if ingress.Spec.Rules[0].Host != "example.com" {
		t.Errorf("expected first rule host example.com, got %s", ingress.Spec.Rules[0].Host)
	}
	if ingress.Spec.Rules[1].Host != "api.example.com" {
		t.Errorf("expected second rule host api.example.com, got %s", ingress.Spec.Rules[1].Host)
	}
}
