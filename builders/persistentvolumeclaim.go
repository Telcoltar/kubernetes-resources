package builders

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type PVCBuilder = PersistentVolumeClaimBuilder

// Base builder type with generic parent
type pvcBuilder[T any] struct {
	spec   corev1.PersistentVolumeClaimSpec
	parent T
}

func (b *pvcBuilder[T]) AccessModes(modes ...corev1.PersistentVolumeAccessMode) T {
	b.spec.AccessModes = append(b.spec.AccessModes, modes...)
	return b.parent
}

func (b *pvcBuilder[T]) SetStorageClass(name string) {
	b.spec.StorageClassName = &name
}

func (b *pvcBuilder[T]) StorageClassP(name *string) T {
	b.spec.StorageClassName = name
	return b.parent
}

func (b *pvcBuilder[T]) StorageClass(name string) T {
	b.SetStorageClass(name)
	return b.parent
}

func (b *pvcBuilder[T]) Size(size string) T {
	b.spec.Resources = corev1.VolumeResourceRequirements{
		Requests: corev1.ResourceList{
			corev1.ResourceStorage: resource.MustParse(size),
		},
	}
	return b.parent
}

// PersistentVolumeClaimBuilder for standalone PVCs
type PersistentVolumeClaimBuilder struct {
	*objectMetaBuilder[PersistentVolumeClaimBuilder]
	*pvcBuilder[*PersistentVolumeClaimBuilder]
}

func PVC(name string) *PersistentVolumeClaimBuilder {
	return PersistentVolumeClaim(name)
}

func PersistentVolumeClaim(name string) *PersistentVolumeClaimBuilder {
	b := &PersistentVolumeClaimBuilder{}
	b.pvcBuilder = &pvcBuilder[*PersistentVolumeClaimBuilder]{
		spec:   corev1.PersistentVolumeClaimSpec{},
		parent: b,
	}
	b.objectMetaBuilder = ObjectMeta(b).name(name)
	return b
}

func (b *PersistentVolumeClaimBuilder) Build() (runtime.Object, error) {
	return &corev1.PersistentVolumeClaim{
		TypeMeta:   persistentVolumeClaimType,
		ObjectMeta: b.objectMetaBuilder.Build(),
		Spec:       b.spec,
	}, nil
}

func (b *PersistentVolumeClaimBuilder) Volume() *VolumePVCBuilder {
	return Volume(b.GetName()).PVC(b.GetName())
}

func (b *PersistentVolumeClaimBuilder) Mount(path string) *VolumeMountBuilder {
	return VolumeMount(b.GetName()).Path(path)
}

// PVCTemplateBuilder for StatefulSet volume claim templates
type PVCTemplateBuilder struct {
	*objectMetaBuilder[PVCTemplateBuilder]
	*pvcBuilder[*PVCTemplateBuilder]
}

func PVCTemplate(name string) *PVCTemplateBuilder {
	b := &PVCTemplateBuilder{}
	b.pvcBuilder = &pvcBuilder[*PVCTemplateBuilder]{
		spec:   corev1.PersistentVolumeClaimSpec{},
		parent: b,
	}
	b.objectMetaBuilder = ObjectMeta(b).name(name)
	return b
}

func (b *PVCTemplateBuilder) Build() corev1.PersistentVolumeClaim {
	return corev1.PersistentVolumeClaim{
		ObjectMeta: b.objectMetaBuilder.Build(),
		Spec:       b.spec,
	}
}

func (b *PVCTemplateBuilder) Volume() *VolumePVCBuilder {
	return Volume(b.GetName()).PVC(b.GetName())
}

func (b *PVCTemplateBuilder) Mount(path string) *VolumeMountBuilder {
	return VolumeMount(b.GetName()).Path(path)
}
