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

var daemonSetType = metav1.TypeMeta{
	APIVersion: "apps/v1",
	Kind:       "DaemonSet",
}

var statefulSetType = metav1.TypeMeta{
	APIVersion: "apps/v1",
	Kind:       "StatefulSet",
}

var persistentVolumeClaimType = metav1.TypeMeta{
	APIVersion: "v1",
	Kind:       "PersistentVolumeClaim",
}

var roleType = metav1.TypeMeta{
	APIVersion: "rbac/v1",
	Kind:       "Role",
}

var roleBindingType = metav1.TypeMeta{
	APIVersion: "rbac/v1",
	Kind:       "RoleBinding",
}

var serviceAccountType = metav1.TypeMeta{
	APIVersion: "v1",
	Kind:       "ServiceAccount",
}

var routeType = metav1.TypeMeta{
	APIVersion: "route.openshift.io/v1",
	Kind:       "Route",
}

var ingressType = metav1.TypeMeta{
	Kind:       "Ingress",
	APIVersion: "networking.k8s.io/v1",
}
