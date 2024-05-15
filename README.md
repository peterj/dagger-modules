# My collection of Dagger modules



## Envoy Proxy Module

This module enables you to do the following:

- run an instance of Envoy proxy using the provided configuration
- validate the Envoy configuration file


### Usage

To run an instance of Envoy proxy and expose the ports on the host, run the following command:

```bash
dagger call envoy-proxy-service --config $PWD/examples/httpbin-sample.yaml --ports 10000 --ports 9901 up
```

To validate the Envoy configuration file, run the following command:

```bash
dagger call validate-config  --config $PWD/examples/httpbin-sample.yaml
```

```console
configuration '/etc/envoy/envoy.yaml' OK
```