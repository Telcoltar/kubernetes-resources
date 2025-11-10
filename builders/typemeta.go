package builders

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

var serviceType = metav1.TypeMeta{
	APIVersion: "v1",
	Kind:       "Service",
}

var secretType = metav1.TypeMeta{
	APIVersion: "v1",
	Kind:       "Secret",
}

var configMapType = metav1.TypeMeta{
	APIVersion: "v1",
	Kind:       "ConfigMap",
}

var deploymentType = metav1.TypeMeta{
	APIVersion: "apps/v1",
	Kind:       "Deployment",
}

var statefulSetType = metav1.TypeMeta{
	APIVersion: "apps/v1",
	Kind:       "StatefulSet",
}

var persistentVolumeClaimType = metav1.TypeMeta{
	APIVersion: "v1",
	Kind:       "PersistentVolumeClaim",
}
