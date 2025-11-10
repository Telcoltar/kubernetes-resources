package builders

import corev1 "k8s.io/api/core/v1"

type PodSpecBuilder[P any] struct {
	spec corev1.PodSpec

	containers      *ContainersBuilder
	initContianers  *ContainersBuilder
	securityContext *PodSecurityContextBuilder

	volumes []VolumeBuilderI

	pullSecrets []string

	parent *P
}

func PodSpec[P any](parent *P) *PodSpecBuilder[P] {
	b := &PodSpecBuilder[P]{}
	b.containers = Containers()
	b.initContianers = Containers()

	b.parent = parent
	return b
}

func (b *PodSpecBuilder[P]) TerminationGracePeriod(period int64) *P {
	b.spec.TerminationGracePeriodSeconds = &period
	return b.parent
}

func (b *PodSpecBuilder[P]) Containers(cb ...*ContainerBuilder) *P {
	b.containers.Containers(cb...)
	return b.parent
}

func (b *PodSpecBuilder[P]) InitContainers(cb ...*ContainerBuilder) *P {
	b.initContianers.Containers(cb...)
	return b.parent
}

func (b *PodSpecBuilder[P]) ServiceAccountName(name string) *P {
	b.spec.ServiceAccountName = name
	return b.parent
}

func (b *PodSpecBuilder[P]) SecurityContext(pscb *PodSecurityContextBuilder) *P {
	b.securityContext = pscb
	return b.parent
}

func (b *PodSpecBuilder[P]) PullSecrets(secrets ...string) *P {
	b.pullSecrets = append(b.pullSecrets, secrets...)
	return b.parent
}

func (b *PodSpecBuilder[P]) Volumes(volumes ...VolumeBuilderI) *P {
	b.volumes = append(b.volumes, volumes...)
	return b.parent
}

func (b *PodSpecBuilder[P]) Build() corev1.PodSpec {
	b.spec.Containers = b.containers.Build()
	b.spec.InitContainers = b.initContianers.Build()
	if b.securityContext != nil {
		b.spec.SecurityContext = b.securityContext.Build()
	}
	pullSecrets := make([]corev1.LocalObjectReference, len(b.pullSecrets))
	for i, secretName := range b.pullSecrets {
		pullSecrets[i] = corev1.LocalObjectReference{
			Name: secretName,
		}
	}
	b.spec.ImagePullSecrets = pullSecrets
	b.spec.Volumes = BuildAll(b.volumes...)
	return b.spec
}
