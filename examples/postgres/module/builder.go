// Package postgres contains functions to build Postgres modul
package postgres

import (
	"fmt"

	"github.com/Telcoltar/kubernetes-resources/builders"
	"github.com/Telcoltar/kubernetes-resources/utils"

	"k8s.io/apimachinery/pkg/labels"
)

type PostgresBuilder struct {
	name     string
	version  string
	password string
}

func Postgres() *PostgresBuilder {
	return &PostgresBuilder{}
}

func (b *PostgresBuilder) Name(name string) *PostgresBuilder {
	b.name = name
	return b
}

func (b *PostgresBuilder) Version(version string) *PostgresBuilder {
	b.version = version
	return b
}

func (b *PostgresBuilder) Password(pass string) *PostgresBuilder {
	b.password = pass
	return b
}

func (b *PostgresBuilder) Assemble() builders.Builders {
	instance := "postgres"
	if b.name != "" {
		instance = fmt.Sprintf("postgres-%s", b.name)
	}
	labels := labels.Set{
		utils.LabelName:     "postgres",
		utils.LabelInstance: instance,
	}

	components := builders.Builders{}

	dataTemplate := builders.PVCTemplate("pgdata").Size("50Gi")

	port := builders.ContainerPort().
		Port(5432).Name("postgres")

	passwordSecret := builders.Secret(fmt.Sprintf("%s-user-pass", instance)).
		StringDatum("password", b.password)

	mainContainer := builders.Container("postgres").
		Envs(
			builders.Env("POSTGRES_PASSWORD").FromSecret(passwordSecret.GetName()).Key("password"),
		).
		VolumeMounts(dataTemplate.Mount("/var/lib/postgresql/data")).
		Resources(builders.Resources().ML("4Gi").CR("2").MR("4Gi")).
		Image(builders.Image("postgres").Tag(fmt.Sprintf("%s-alpine3.22", b.version)))

	service := builders.Service(instance).
		Ports(port.Service()).
		Selectors(labels)

	statefulSet := builders.StatefulSet(instance).Containers(mainContainer).
		Selector(labels).ServiceName(service.GetName()).
		ClaimTepmplates(dataTemplate).MinReady(50)

	components = append(components, passwordSecret, service, statefulSet)

	components.Labels(labels).Label(utils.LabelVersion, b.version)
	return components
}
