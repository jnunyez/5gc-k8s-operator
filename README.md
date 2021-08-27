# 5G-C Operator

Generator 5G-C configuration when kubernetes 5G-C resources change. The operator renders the configuration templates provided in the resource into a new configmap and mounts that to the 5G-C containers

## Quickstart

* Install binary folllowing up-to-date instructions showed [here](https://sdk.operatorframework.io/docs/installation/). Note we are following the lin in [here](https://marketplace.redhat.com/en-us/blog/building-an-operator)

* The operator-sdk asking for go1.13 using skip-go-version-check as a quick workaround:

```console
operator-sdk init --domain=5gc.dev --owner "Jose Nuñez" --repo=github.com/jnunyez/5gc-k8s-operator --skip-go-version-check
```

The resulting operator scaffolding should look like this:

```console
.
├── bin
├── config
│   ├── certmanager
│   ├── default
│   ├── manager
│   ├── prometheus
│   ├── rbac
│   ├── scorecard
│   │   ├── bases
│   │   └── patches
│   └── webhook
└── hack
```

* Create an API to extend the k8s API and create controller logic skeleton:

```console
operator-sdk create api --version v1alpha1 --kind Core5g --group open5gcore --resource=true --controller=true
```

This will scaffold the operator API at `api/v1alpha1/operator_types.go` and the controller at `controllers/operator_controller.go`:

```console
.
├── api
│   └── v1alpha1
├── bin
├── config
│   ├── certmanager
│   ├── crd
│   │   └── patches
│   ├── default
│   ├── manager
│   ├── prometheus
│   ├── rbac
│   ├── samples
│   ├── scorecard
│   │   ├── bases
│   │   └── patches
│   └── webhook
├── controllers
└── hack
```