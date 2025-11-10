package builders

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type PVCBuilder = PersistentVolumeClaimBuilder

// Base builder type with generic parent
type pvcBaseBuilder[T any] struct {
	*objectMetaBuilder[T]
	spec   corev1.PersistentVolumeClaimSpec
	parent *T
}

func pvcBase[T any](name string, parent *T) *pvcBaseBuilder[T] {
	return &pvcBaseBuilder[T]{
		objectMetaBuilder: ObjectMeta(parent).name(name),
		spec:              corev1.PersistentVolumeClaimSpec{},
		parent:            parent,
	}
}

func (b *pvcBaseBuilder[T]) AccessModes(modes ...corev1.PersistentVolumeAccessMode) *T {
	b.spec.AccessModes = append(b.spec.AccessModes, modes...)
	return b.parent
}

func (b *pvcBaseBuilder[T]) SetStorageClass(name string) {
	b.spec.StorageClassName = &name
}

func (b *pvcBaseBuilder[T]) StorageClassP(name *string) *T {
	b.spec.StorageClassName = name
	return b.parent
}

func (b *pvcBaseBuilder[T]) StorageClass(name string) *T {
	b.SetStorageClass(name)
	return b.parent
}

func (b *pvcBaseBuilder[T]) Size(size string) *T {
	b.spec.Resources = corev1.VolumeResourceRequirements{
		Requests: corev1.ResourceList{
			corev1.ResourceStorage: resource.MustParse(size),
		},
	}
	return b.parent
}

// PersistentVolumeClaimBuilder for standalone PVCs
type PersistentVolumeClaimBuilder struct {
	*pvcBaseBuilder[PersistentVolumeClaimBuilder]
}

func PVC(name string) *PersistentVolumeClaimBuilder {
	return PersistentVolumeClaim(name)
}

func PersistentVolumeClaim(name string) *PersistentVolumeClaimBuilder {
	b := &PersistentVolumeClaimBuilder{}
	b.pvcBaseBuilder = pvcBase(name, b)
	return b
}

func (b *PersistentVolumeClaimBuilder) Build() (runtime.Object, error) {
	return &corev1.PersistentVolumeClaim{
		TypeMeta:   persistentVolumeClaimType,
		ObjectMeta: b.objectMetaBuilder.Build(),
		Spec:       b.spec,
	}, nil
}

func (b *pvcBaseBuilder[T]) Volume() *VolumePVCBuilder {
	return Volume(b.GetName()).PVC(b.GetName())
}

func (b *pvcBaseBuilder[T]) Mount(path string) *VolumeMountBuilder {
	return VolumeMount(b.GetName()).Path(path)
}

// PVCTemplateBuilder for StatefulSet volume claim templates
type PVCTemplateBuilder struct {
	*pvcBaseBuilder[PVCTemplateBuilder]
}

func PVCTemplate(name string) *PVCTemplateBuilder {
	b := &PVCTemplateBuilder{}
	b.pvcBaseBuilder = pvcBase(name, b)
	return b
}

func (b *PVCTemplateBuilder) Build() corev1.PersistentVolumeClaim {
	return corev1.PersistentVolumeClaim{
		ObjectMeta: b.objectMetaBuilder.Build(),
		Spec:       b.spec,
	}
}
