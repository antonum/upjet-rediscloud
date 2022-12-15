# Provider RedisCloud

`provider-rediscloud` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
RedisCloud API.

## QuickStart (user mode)

Install [UP CLI](https://docs.upbound.io/cli/) 
```
brew install upbound/tap/up
```
or see link above for instructions on other platforms.

Install latest Crossplane to your current kubernetes cluster:
```
up uxp install
```
That would create `upbound-system` namespace and install core crossplane components.

Install Rediscloud provider:
```
kubectl apply -n upbound-system -f examples/install-open.yaml 
```
Note: in pre-released stage `install-open.yaml` uses public, but potentially outdataed version of provider image. To access the latest you need to use examples/install.yaml that requires credentials to the private repo. 

Create `secret.yaml` from `examples/providerconfig/secret.yaml.tmpl` using your Redis Cloud key/secret. Install  secret and provider config.
```
kubectl apply -f examples/providerconfig/secret.yaml                    
kubectl apply -n upbound-system -f examples/providerconfig/providerconfig.yaml 
```

Use `examples/database.yaml` as example to create the database. It must use pre-existing subscription id.

```
kubectl apply -f - << EOF  
apiVersion: database.rediscloud.upbound.io/v1alpha1
kind: Database
metadata:
  name: my-first-database
spec:
  forProvider:
    subscriptionId: 1922316
    name: "my-first-database"
    protocol: "redis"
    memoryLimitInGb: 1
    dataPersistence: "none"
    throughputMeasurementBy: "operations-per-second"
    throughputMeasurementValue: 1000
  providerConfigRef:
    name: default
EOF
```



## Getting Started (Development out-of-cluster mode)

Build the provider:
```
make generate
```

Apply CRDs
```
kubectl apply -f package/crds
```



Create secret.yaml from examples/providerconfig/secret.yaml.tmpl using your Redis Cloud key/secred and apply it, along with the provider config.
```
kubectl apply -f examples/providerconfig/secret.yaml
kubectl apply -f examples/providerconfig/providerconfig.yaml
```

Run provider:
```
make run
```

Use `examples/database.yaml` as example to create the database. It must use pre-existing subscription id.

```
kubectl apply -f - << EOF  
apiVersion: database.rediscloud.upbound.io/v1alpha1
kind: Database
metadata:
  name: my-first-database
spec:
  forProvider:
    subscriptionId: 1889645
    name: "my-first-database"
    protocol: "redis"
    memoryLimitInGb: 1
    dataPersistence: "none"
    throughputMeasurementBy: "operations-per-second"
    throughputMeasurementValue: 1000
  providerConfigRef:
    name: default
EOF
```

Verify that database successfully created in RedisCloud console and by running:
```
kubectl describe database my-first-database
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/antonum/upjet-rediscloud).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/antonum/upjet-rediscloud/issues).
