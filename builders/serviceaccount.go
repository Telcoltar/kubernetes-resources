package builders

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type ServiceAccountBuilder struct {
	*objectMetaBuilder[ServiceAccountBuilder]
	imagePullSecrets []string
}

func ServiceAccount(name string) *ServiceAccountBuilder {
	sab := &ServiceAccountBuilder{}
	sab.objectMetaBuilder = ObjectMeta(sab).name(name)
	return sab
}

func (sab *ServiceAccountBuilder) ImagePullSecrets(secretNames ...string) *ServiceAccountBuilder {
	sab.imagePullSecrets = append(sab.imagePullSecrets, secretNames...)
	return sab
}

func (sab *ServiceAccountBuilder) Build() (runtime.Object, error) {
	pullSecrets := make([]corev1.LocalObjectReference, len(sab.imagePullSecrets))
	for i, secretName := range sab.imagePullSecrets {
		pullSecrets[i] = corev1.LocalObjectReference{Name: secretName}
	}
	return &corev1.ServiceAccount{
		TypeMeta:         serviceAccountType,
		ObjectMeta:       sab.objectMetaBuilder.Build(),
		ImagePullSecrets: pullSecrets,
	}, nil
}
