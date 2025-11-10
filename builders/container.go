package builders

import corev1 "k8s.io/api/core/v1"

type ContainerBuilder struct {
	container       corev1.Container
	image           *ImageBuilder
	ports           []ContainerPortBuilderI
	envs            []EnvBuilderI
	mounts          []*VolumeMountBuilder
	securityContext *ContainerSecurityContextBuilder
	readinessProbe  ProbeBuilderI
	livenessProbe   ProbeBuilderI
}

func Container(name string) *ContainerBuilder {
	return &ContainerBuilder{
		container: corev1.Container{
			Name: name,
		},
	}
}

func (b *ContainerBuilder) Ports(ports ...ContainerPortBuilderI) *ContainerBuilder {
	b.ports = append(b.ports, ports...)
	return b
}

func (b *ContainerBuilder) PullPolicy(policy corev1.PullPolicy) *ContainerBuilder {
	b.container.ImagePullPolicy = policy
	return b
}

func (b *ContainerBuilder) Envs(envs ...EnvBuilderI) *ContainerBuilder {
	b.envs = append(b.envs, envs...)
	return b
}

func (b *ContainerBuilder) Image(image *ImageBuilder) *ContainerBuilder {
	b.image = image
	return b
}

func (b *ContainerBuilder) LivenessProbe(probe ProbeBuilderI) *ContainerBuilder {
	b.livenessProbe = probe
	return b
}

func (b *ContainerBuilder) ReadinessProbe(probe ProbeBuilderI) *ContainerBuilder {
	b.readinessProbe = probe
	return b
}

func (b *ContainerBuilder) SecurityContext(cscb *ContainerSecurityContextBuilder) *ContainerBuilder {
	b.securityContext = cscb
	return b
}

func (b *ContainerBuilder) Resources(rb *ResourceBuilder) *ContainerBuilder {
	b.container.Resources = rb.Build()
	return b
}

func (b *ContainerBuilder) VolumeMounts(mounts ...*VolumeMountBuilder) *ContainerBuilder {
	b.mounts = append(b.mounts, mounts...)
	return b
}

func (b *ContainerBuilder) Build() corev1.Container {
	b.container.Ports = BuildAll(b.ports...)
	if b.image != nil {
		b.container.Image = b.image.Build()
	}
	b.container.Env = BuildAll(b.envs...)
	if b.livenessProbe != nil {
		b.container.LivenessProbe = b.livenessProbe.Build()
	}
	if b.readinessProbe != nil {
		b.container.ReadinessProbe = b.readinessProbe.Build()
	}
	if b.securityContext != nil {
		b.container.SecurityContext = b.securityContext.Build()
	}
	b.container.VolumeMounts = append(b.container.VolumeMounts, BuildAll(b.mounts...)...)
	return b.container
}
