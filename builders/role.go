package builders

import (
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type RoleBuilder struct {
	*objectMetaBuilder[RoleBuilder]
	rules []rbacv1.PolicyRule
}

func Role(name string) *RoleBuilder {
	rb := &RoleBuilder{}
	rb.objectMetaBuilder = ObjectMeta(rb).name(name)
	return rb
}

func (rb *RoleBuilder) Rules(rules ...*RoleRuleBuilder) *RoleBuilder {
	rb.rules = append(rb.rules, BuildAll(rules...)...)
	return rb
}

func (rb *RoleBuilder) Build() (runtime.Object, error) {
	return &rbacv1.Role{
		TypeMeta:   roleType,
		ObjectMeta: rb.objectMetaBuilder.Build(),
		Rules:      rb.rules,
	}, nil
}

type RoleRuleBuilder struct {
	rule rbacv1.PolicyRule
}

func RoleRule() *RoleRuleBuilder {
	return &RoleRuleBuilder{
		rule: rbacv1.PolicyRule{},
	}
}

func (rbb *RoleRuleBuilder) Verbs(verbs ...string) *RoleRuleBuilder {
	rbb.rule.Verbs = append(rbb.rule.Verbs, verbs...)
	return rbb
}

func (rbb *RoleRuleBuilder) APIGroups(groups ...string) *RoleRuleBuilder {
	rbb.rule.APIGroups = append(rbb.rule.APIGroups, groups...)
	return rbb
}

func (rbb *RoleRuleBuilder) Resources(resources ...string) *RoleRuleBuilder {
	rbb.rule.Resources = append(rbb.rule.Resources, resources...)
	return rbb
}

func (rbb *RoleRuleBuilder) Build() rbacv1.PolicyRule {
	return rbb.rule
}
