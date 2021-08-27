#5G-C Operator

## Quickstart

* Install binary folllowing up-to-date instructions showed [here](https://sdk.operatorframework.io/docs/installation/). Note we are following the lin in [here](https://marketplace.redhat.com/en-us/blog/building-an-operator)

* The operator-sdk asking for go1.13 using skip-go-version-check as a quick workaround:

```console
operator-sdk init --domain=5gc.dev --owner "Jose Nuñez" --repo=github.com/jnunyez/5gc-k8s-operator --skip-go-version-check
```

* Create an API and controller logic skeleton:

```console
operator-sdk create api --version v1alpha1 --kind Core5g --group open5gcore --resource=true --controller=true
```