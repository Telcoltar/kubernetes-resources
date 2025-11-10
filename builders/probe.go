package builders

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type ProbeBuilderI interface {
	Build() *corev1.Probe
}

type commonProbeBuilder[P ProbeBuilderI] struct {
	probe  corev1.Probe
	parent P
}

func commonProbe[P ProbeBuilderI](parent P) *commonProbeBuilder[P] {
	return &commonProbeBuilder[P]{
		corev1.Probe{},
		parent,
	}
}

func changeParent[P ProbeBuilderI, Q ProbeBuilderI](b *commonProbeBuilder[P], parent Q) *commonProbeBuilder[Q] {
	return &commonProbeBuilder[Q]{
		probe:  b.probe,
		parent: parent,
	}
}

func (b *commonProbeBuilder[P]) Build() *corev1.Probe {
	return b.parent.Build()
}

func (b *commonProbeBuilder[P]) InitialDelay(delay int32) P {
	b.probe.InitialDelaySeconds = delay
	return b.parent
}

func (b *commonProbeBuilder[P]) Period(period int32) P {
	b.probe.PeriodSeconds = period
	return b.parent
}

func (b *commonProbeBuilder[P]) Timeout(timeout int32) P {
	b.probe.TimeoutSeconds = timeout
	return b.parent
}

func (b *commonProbeBuilder[P]) FailureThreshold(threshold int32) P {
	b.probe.FailureThreshold = threshold
	return b.parent
}

type ProbeBuilder struct {
	*commonProbeBuilder[*ProbeBuilder]
}

func (b *ProbeBuilder) Build() *corev1.Probe {
	return &b.probe
}

func Probe() *ProbeBuilder {
	b := &ProbeBuilder{}
	b.commonProbeBuilder = commonProbe(b)
	return b
}

type HTTPProbeBuilder struct {
	*commonProbeBuilder[*HTTPProbeBuilder]
	action corev1.HTTPGetAction
}

func (b *ProbeBuilder) HTTP() *HTTPProbeBuilder {
	httpPB := &HTTPProbeBuilder{
		action: corev1.HTTPGetAction{},
	}
	httpPB.commonProbeBuilder = changeParent(b.commonProbeBuilder, httpPB)
	return httpPB
}

func (b *HTTPProbeBuilder) Path(path string) *HTTPProbeBuilder {
	b.action.Path = path
	return b
}

func (b *HTTPProbeBuilder) Port(num int32) *HTTPProbeBuilder {
	b.action.Port = intstr.FromInt32(num)
	return b
}

func (b *HTTPProbeBuilder) PortName(name string) *HTTPProbeBuilder {
	b.action.Port = intstr.FromString(name)
	return b
}

func (b *HTTPProbeBuilder) Scheme(scheme corev1.URIScheme) *HTTPProbeBuilder {
	b.action.Scheme = scheme
	return b
}

func (b *HTTPProbeBuilder) Build() *corev1.Probe {
	b.probe.ProbeHandler = corev1.ProbeHandler{
		HTTPGet: &b.action,
	}
	return &b.probe
}
