package builders

import (
	corev1 "k8s.io/api/core/v1"
)

type VolumeMountBuilder struct {
	mount corev1.VolumeMount
}

func VolumeMount(name string) *VolumeMountBuilder {
	return &VolumeMountBuilder{
		mount: corev1.VolumeMount{
			Name: name,
		},
	}
}

func (b *VolumeMountBuilder) Path(path string) *VolumeMountBuilder {
	b.mount.MountPath = path
	return b
}

func (b *VolumeMountBuilder) SubPath(path string) *VolumeMountBuilder {
	b.mount.SubPath = path
	return b
}

func (b *VolumeMountBuilder) ReadOnly() *VolumeMountBuilder {
	b.mount.ReadOnly = true
	return b
}

func (b *VolumeMountBuilder) Build() corev1.VolumeMount {
	return b.mount
}
