package main

import (
	"fmt"
	"log"

	"github.com/Telcoltar/kubernetes-resources/builders"
)

func main() {
	// Example 1: Simple HTTP route
	httpRoute := builders.Route("my-app-route").
		Namespace("default").
		Host("myapp.example.com").
		Path("/api").
		To("my-app-service").
		PortName("http")

	// Example 2: HTTPS route with edge termination
	httpsRoute := builders.Route("secure-route").
		Namespace("default").
		Host("secure.example.com").
		To("secure-service").
		Port(8443).
		TLSEdge().
		InsecureEdgeTerminationPolicy("Redirect")

	// Example 3: Route with passthrough TLS
	passthroughRoute := builders.Route("passthrough-route").
		Namespace("default").
		Host("passthrough.example.com").
		To("tls-service").
		PortName("https").
		TLSPassthrough()

	// Build and print YAML
	routes := builders.Builders{httpRoute, httpsRoute, passthroughRoute}

	yaml, err := routes.AsYaml()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(yaml))
}
