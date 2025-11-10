package builders

import (
	corev1 "k8s.io/api/core/v1"
)

// ContainerSecurityContextBuilder provides fluent setters for corev1.SecurityContext.
type ContainerSecurityContextBuilder struct {
	sc corev1.SecurityContext
}

// ContainerSecurityContext creates a new standalone ContainerSecurityContextBuilder.
func ContainerSecurityContext() *ContainerSecurityContextBuilder {
	return &ContainerSecurityContextBuilder{}
}

// RunAsUser sets the user ID to run the entrypoint of the container process.
func (b *ContainerSecurityContextBuilder) RunAsUser(uid int64) *ContainerSecurityContextBuilder {
	b.sc.RunAsUser = &uid
	return b
}

// RunAsGroup sets the primary group ID for all processes within the container.
func (b *ContainerSecurityContextBuilder) RunAsGroup(gid int64) *ContainerSecurityContextBuilder {
	b.sc.RunAsGroup = &gid
	return b
}

// RunAsNonRoot indicates that the container must run as a non-root user.
func (b *ContainerSecurityContextBuilder) RunAsNonRoot(v bool) *ContainerSecurityContextBuilder {
	b.sc.RunAsNonRoot = &v
	return b
}

// ReadOnlyRootFilesystem sets the container's root filesystem to read-only.
func (b *ContainerSecurityContextBuilder) ReadOnlyRootFilesystem(v bool) *ContainerSecurityContextBuilder {
	b.sc.ReadOnlyRootFilesystem = &v
	return b
}

// AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process.
func (b *ContainerSecurityContextBuilder) AllowPrivilegeEscalation(v bool) *ContainerSecurityContextBuilder {
	b.sc.AllowPrivilegeEscalation = &v
	return b
}

// Privileged runs the container in privileged mode.
func (b *ContainerSecurityContextBuilder) Privileged(v bool) *ContainerSecurityContextBuilder {
	b.sc.Privileged = &v
	return b
}

// ProcMount sets the proc filesystem mount type for the container.
func (b *ContainerSecurityContextBuilder) ProcMount(v corev1.ProcMountType) *ContainerSecurityContextBuilder {
	b.sc.ProcMount = &v
	return b
}

// CapabilitiesAdd adds Linux capabilities to the container.
func (b *ContainerSecurityContextBuilder) CapabilitiesAdd(caps ...corev1.Capability) *ContainerSecurityContextBuilder {
	if b.sc.Capabilities == nil {
		b.sc.Capabilities = &corev1.Capabilities{}
	}
	b.sc.Capabilities.Add = append(b.sc.Capabilities.Add, caps...)
	return b
}

// CapabilitiesDrop drops Linux capabilities from the container.
func (b *ContainerSecurityContextBuilder) CapabilitiesDrop(caps ...corev1.Capability) *ContainerSecurityContextBuilder {
	if b.sc.Capabilities == nil {
		b.sc.Capabilities = &corev1.Capabilities{}
	}
	b.sc.Capabilities.Drop = append(b.sc.Capabilities.Drop, caps...)
	return b
}

// SeccompProfileType sets the seccomp profile type for the container.
func (b *ContainerSecurityContextBuilder) SeccompProfileType(t corev1.SeccompProfileType) *ContainerSecurityContextBuilder {
	if b.sc.SeccompProfile == nil {
		b.sc.SeccompProfile = &corev1.SeccompProfile{}
	}
	b.sc.SeccompProfile.Type = t
	return b
}

// SeccompProfileLocalhost sets the localhost seccomp profile path for the container.
func (b *ContainerSecurityContextBuilder) SeccompProfileLocalhost(path string) *ContainerSecurityContextBuilder {
	if b.sc.SeccompProfile == nil {
		b.sc.SeccompProfile = &corev1.SeccompProfile{}
	}
	b.sc.SeccompProfile.LocalhostProfile = &path
	return b
}

// SELinuxOptions sets SELinux options for the container.
func (b *ContainerSecurityContextBuilder) SELinuxOptions(opts corev1.SELinuxOptions) *ContainerSecurityContextBuilder {
	b.sc.SELinuxOptions = &opts
	return b
}

// WindowsOptions sets Windows-specific security options for the container.
func (b *ContainerSecurityContextBuilder) WindowsOptions(opts corev1.WindowsSecurityContextOptions) *ContainerSecurityContextBuilder {
	b.sc.WindowsOptions = &opts
	return b
}

// Build returns the constructed SecurityContext value.
func (b *ContainerSecurityContextBuilder) Build() *corev1.SecurityContext {
	return &b.sc
}
