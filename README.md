# My collection of Dagger modules

## Fly.io Module

This module enables you to do the following:

- deploy an application to Fly.io

### Usage

To deploy an application to Fly.io, run the following command:

```bash
dagger call deploy --src [source-folder] --token=env:FLYIO_TOKEN
```

Note that the token can come from an environment variable or a file as well. Make sure you set that before running the deploy command.

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
