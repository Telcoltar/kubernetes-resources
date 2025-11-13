package builders

import (
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type RoleBindingBuilder struct {
	*objectMetaBuilder[RoleBindingBuilder]
	roleRef  rbacv1.RoleRef
	subjects []rbacv1.Subject
}

func RoleBinding(name string) *RoleBindingBuilder {
	rbb := &RoleBindingBuilder{}
	rbb.objectMetaBuilder = ObjectMeta(rbb).name(name)
	return rbb
}

func (rbb *RoleBindingBuilder) RoleRef(apiGroup, kind, name string) *RoleBindingBuilder {
	rbb.roleRef = rbacv1.RoleRef{APIGroup: apiGroup, Kind: kind, Name: name}
	return rbb
}

func (rbb *RoleBindingBuilder) Subject(kind, name string) *RoleBindingBuilder {
	rbb.subjects = append(rbb.subjects, rbacv1.Subject{Kind: kind, Name: name})
	return rbb
}

func (rbb *RoleBindingBuilder) Build() (runtime.Object, error) {
	return &rbacv1.RoleBinding{
		TypeMeta:   roleBindingType,
		ObjectMeta: rbb.objectMetaBuilder.Build(),
		RoleRef:    rbb.roleRef,
		Subjects:   rbb.subjects,
	}, nil
}
