# kpt-remove-resource

> a [kpt](https://googlecontainertools.github.io/kpt/) function for removing a
Kubernetes resource

## CLI usage

To use this function via the kpt CLI run the following:

```bash
kpt fn run . --image dustinspecker/kpt-remove-resource:lastest -- kind=KIND name=NAME namespace=NAMESPACE
```

A kubernetes resource matching the `kind`, `name`, and `namespace` will be removed.

## declarative usage

Create a YAML file with the following content:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  annotations:
    config.kubernetes.io/function: |
      container:
        image: dustinspecker/kpt-remove-resource:latest
data:
  kind: Service
  name: ingress-nginx-controller
  namespace: ingress-nginx
```

and run:

```bash
kpt fn run .
```

A kubernetes resource matching the `kind`, `name`, and `namespace` will be removed.
