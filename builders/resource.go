package builders

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

type ResourceBuilder struct {
	resources corev1.ResourceRequirements
}

func Resources() *ResourceBuilder {
	return &ResourceBuilder{
		resources: corev1.ResourceRequirements{
			Limits:   corev1.ResourceList{},
			Requests: corev1.ResourceList{},
		},
	}
}

func (b *ResourceBuilder) CPURequest(request string) *ResourceBuilder {
	b.resources.Requests[corev1.ResourceCPU] = resource.MustParse(request)
	return b
}

func (b *ResourceBuilder) CR(request string) *ResourceBuilder {
	return b.CPURequest(request)
}

func (b *ResourceBuilder) CPULimit(limit string) *ResourceBuilder {
	b.resources.Limits[corev1.ResourceCPU] = resource.MustParse(limit)
	return b
}

func (b *ResourceBuilder) CL(limit string) *ResourceBuilder {
	return b.CPULimit(limit)
}

func (b *ResourceBuilder) MemoryRequest(request string) *ResourceBuilder {
	b.resources.Requests[corev1.ResourceMemory] = resource.MustParse(request)
	return b
}

func (b *ResourceBuilder) MR(request string) *ResourceBuilder {
	return b.MemoryRequest(request)
}

func (b *ResourceBuilder) MemoryLimit(limit string) *ResourceBuilder {
	b.resources.Limits[corev1.ResourceMemory] = resource.MustParse(limit)
	return b
}

func (b *ResourceBuilder) ML(limit string) *ResourceBuilder {
	return b.MemoryLimit(limit)
}

func (b *ResourceBuilder) Build() corev1.ResourceRequirements {
	return b.resources
}
