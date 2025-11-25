# Kubernetes Resource Builder

Work in Progress golang lib to create Kubernetes Resources with a fluent Builder pattern.

## Available Builders

### Core Kubernetes Resources

- **Deployment** - Create Kubernetes Deployments
- **DaemonSet** - Create Kubernetes DaemonSets
- **StatefulSet** - Create Kubernetes StatefulSets
- **Service** - Create Kubernetes Services
- **ConfigMap** - Create ConfigMaps
- **Secret** - Create Secrets
- **PersistentVolumeClaim** - Create PVCs
- **ServiceAccount** - Create Service Accounts
- **Role** - Create RBAC Roles
- **RoleBinding** - Create RBAC Role Bindings

### OpenShift Resources

- **Route** - Create OpenShift Routes

## Usage Example

```go
package main

import (
    "fmt"
    "log"
    "github.com/Telcoltar/kubernetes-resources/builders"
)

func main() {
    // Create an OpenShift Route
    route := builders.Route("my-app-route").
        Namespace("default").
        Host("myapp.example.com").
        Path("/api").
        To("my-app-service").
        PortNamed("http").
        TLSEdge()

    // Build YAML
    yaml, err := builders.BuildYaml(route)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(string(yaml))
}
```

See the `examples/` directory for more usage examples.
