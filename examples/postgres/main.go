package main

import (
	"log"
	"os"

	postgres "github.com/Telcoltar/kubernetes-resources/examples/postgres/module"
)

func main() {
	postgres := postgres.Postgres().Version("18").Password("pass").Assemble()
	postgres.StorageClass("ceph").Namespace("default")
	yamls, err := postgres.AsYaml()
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("postgres.yaml", []byte(yamls), 0o600); err != nil {
		log.Fatal(err)
	}
}
