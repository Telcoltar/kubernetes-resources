package builders

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"
)

type VolumeBuilderI interface {
	Build() corev1.Volume
}

type volumeBaseBuilder struct {
	name          string
	actualBuilder VolumeBuilderI
}

func Volume(name string) *volumeBaseBuilder {
	return &volumeBaseBuilder{
		name: name,
	}
}

func (b *volumeBaseBuilder) Mount(path string) *VolumeMountBuilder {
	return VolumeMount(b.name).Path(path)
}

func (b *volumeBaseBuilder) Build() corev1.Volume {
	if b.actualBuilder != nil {
		return b.actualBuilder.Build()
	}
	return corev1.Volume{
		Name: b.name,
	}
}

type VolumeKeyPathBuilder struct {
	*volumeBaseBuilder
	volumeSource func(items []corev1.KeyToPath, defaultMode *int32, optional *bool) corev1.VolumeSource
	items        []corev1.KeyToPath
	defaultMode  *int32
	optional     *bool
}

func (b *volumeBaseBuilder) ConfigMap(name string) *VolumeKeyPathBuilder {
	vkpb := &VolumeKeyPathBuilder{
		volumeBaseBuilder: b,
		volumeSource: func(items []corev1.KeyToPath, defaultMode *int32, optional *bool) corev1.VolumeSource {
			return corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: name,
					},
					Items:       items,
					DefaultMode: defaultMode,
					Optional:    optional,
				},
			}
		},
	}
	b.actualBuilder = vkpb
	return vkpb
}

func (b *volumeBaseBuilder) Secret(name string) *VolumeKeyPathBuilder {
	vkpb := &VolumeKeyPathBuilder{
		volumeBaseBuilder: b,
		volumeSource: func(items []corev1.KeyToPath, defaultMode *int32, optional *bool) corev1.VolumeSource {
			return corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName:  name,
					Items:       items,
					DefaultMode: defaultMode,
					Optional:    optional,
				},
			}
		},
	}
	b.actualBuilder = vkpb
	return vkpb
}

func (b *VolumeKeyPathBuilder) Items(items []corev1.KeyToPath) *VolumeKeyPathBuilder {
	b.items = append(b.items, items...)
	return b
}

func (b *VolumeKeyPathBuilder) Item(key, path string) *VolumeKeyPathBuilder {
	b.items = append(b.items, corev1.KeyToPath{Key: key, Path: path})
	return b
}

func (b *VolumeKeyPathBuilder) DefaultMode(mode int32) *VolumeKeyPathBuilder {
	b.defaultMode = &mode
	return b
}

func (b *VolumeKeyPathBuilder) Optional() *VolumeKeyPathBuilder {
	b.optional = ptr.To(true)
	return b
}

func (b *VolumeKeyPathBuilder) Build() corev1.Volume {
	return corev1.Volume{
		Name:         b.name,
		VolumeSource: b.volumeSource(b.items, b.defaultMode, b.optional),
	}
}

type VolumeEmptyDirBuilder struct {
	*volumeBaseBuilder
	medium corev1.StorageMedium
}

func (b *volumeBaseBuilder) EmptyDir() *VolumeEmptyDirBuilder {
	vedb := &VolumeEmptyDirBuilder{
		volumeBaseBuilder: b,
	}
	b.actualBuilder = vedb
	return vedb
}

func (b *VolumeEmptyDirBuilder) Medium(medium corev1.StorageMedium) *VolumeEmptyDirBuilder {
	b.medium = medium
	return b
}

func (b *VolumeEmptyDirBuilder) Build() corev1.Volume {
	return corev1.Volume{
		Name: b.name,
		VolumeSource: corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{
				Medium: b.medium,
			},
		},
	}
}

type VolumePVCBuilder struct {
	*volumeBaseBuilder
	claim string
}

func (b *volumeBaseBuilder) PVC(claim string) *VolumePVCBuilder {
	vpb := &VolumePVCBuilder{
		b, claim,
	}
	b.actualBuilder = vpb
	return vpb
}

func (b *VolumePVCBuilder) Build() corev1.Volume {
	return corev1.Volume{
		Name: b.name,
		VolumeSource: corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: b.claim,
			},
		},
	}
}
