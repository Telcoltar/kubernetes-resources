package builders

import (
	"encoding/json"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

// PodSecurityContextBuilder provides fluent setters for corev1.PodSecurityContext.
type PodSecurityContextBuilder struct {
	psc corev1.PodSecurityContext
}

// PodSecurityContext creates a new standalone PodSecurityContextBuilder.
func PodSecurityContext() *PodSecurityContextBuilder {
	return &PodSecurityContextBuilder{}
}

// PodSecurityContextFromYAML creates a builder prefilled from a YAML string.
// Returns an error if the YAML cannot be parsed into a PodSecurityContext.
func PodSecurityContextFromYAML(yml string) (*PodSecurityContextBuilder, error) {
	b := &PodSecurityContextBuilder{}
	if strings.TrimSpace(yml) == "" {
		return b, nil
	}
	jsonBytes, err := yaml.ToJSON([]byte(yml))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonBytes, &b.psc); err != nil {
		return nil, err
	}
	return b, nil
}

// RunAsUser sets the user ID to run the entrypoint of the container process.
func (b *PodSecurityContextBuilder) RunAsUser(uid int64) *PodSecurityContextBuilder {
	b.psc.RunAsUser = &uid
	return b
}

// RunAsGroup sets the primary group ID for all processes within any container of the pod.
func (b *PodSecurityContextBuilder) RunAsGroup(gid int64) *PodSecurityContextBuilder {
	b.psc.RunAsGroup = &gid
	return b
}

// RunAsNonRoot indicates that processes in the pod must run as a non-root user.
func (b *PodSecurityContextBuilder) RunAsNonRoot(v bool) *PodSecurityContextBuilder {
	b.psc.RunAsNonRoot = &v
	return b
}

// FSGroup specifies the group ID for volumes mounted by the pod.
func (b *PodSecurityContextBuilder) FSGroup(gid int64) *PodSecurityContextBuilder {
	b.psc.FSGroup = &gid
	return b
}

// FSGroupChangePolicy defines behavior for changing ownership and permissions of volumes.
func (b *PodSecurityContextBuilder) FSGroupChangePolicy(policy corev1.PodFSGroupChangePolicy) *PodSecurityContextBuilder {
	b.psc.FSGroupChangePolicy = &policy
	return b
}

// SupplementalGroups adds additional group IDs for processes in the pod.
func (b *PodSecurityContextBuilder) SupplementalGroups(groups ...int64) *PodSecurityContextBuilder {
	b.psc.SupplementalGroups = append(b.psc.SupplementalGroups, groups...)
	return b
}

// SeccompProfileType sets the seccomp profile type for the pod.
func (b *PodSecurityContextBuilder) SeccompProfileType(t corev1.SeccompProfileType) *PodSecurityContextBuilder {
	if b.psc.SeccompProfile == nil {
		b.psc.SeccompProfile = &corev1.SeccompProfile{}
	}
	b.psc.SeccompProfile.Type = t
	return b
}

// SeccompProfileLocalhost sets the localhost seccomp profile path for the pod.
func (b *PodSecurityContextBuilder) SeccompProfileLocalhost(path string) *PodSecurityContextBuilder {
	if b.psc.SeccompProfile == nil {
		b.psc.SeccompProfile = &corev1.SeccompProfile{}
	}
	b.psc.SeccompProfile.LocalhostProfile = &path
	return b
}

// SELinuxOptions sets SELinux options for all containers in the pod.
func (b *PodSecurityContextBuilder) SELinuxOptions(opts corev1.SELinuxOptions) *PodSecurityContextBuilder {
	b.psc.SELinuxOptions = &opts
	return b
}

// WindowsOptions sets Windows-specific security options for the pod.
func (b *PodSecurityContextBuilder) WindowsOptions(opts corev1.WindowsSecurityContextOptions) *PodSecurityContextBuilder {
	b.psc.WindowsOptions = &opts
	return b
}

// Sysctls appends kernel parameters to set for the pod.
func (b *PodSecurityContextBuilder) Sysctls(sysctls ...corev1.Sysctl) *PodSecurityContextBuilder {
	b.psc.Sysctls = append(b.psc.Sysctls, sysctls...)
	return b
}

// Build returns the constructed PodSecurityContext value.
func (b *PodSecurityContextBuilder) Build() *corev1.PodSecurityContext {
	return &b.psc
}
