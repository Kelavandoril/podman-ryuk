# Podman Ryuk

This project is a fork of moby-ryuk to add support for Podman clients without workarounds.

## Building

To build the binary only, run:

```shell
go build
```

To build the Linux podman container as the latest tag:

```shell
podman build -f linux/Containerfile -t kelavandoril/ryuk:latest-podman .
```

## Usage
TBD

## Ryuk configuration

In addition to the moby-ryuk environment variables, the following environment variables can be configured to change the behaviour:

| Environment Variable          | Default   | Description  |
| ----------------------------- | -------  | ------------ |
| `RYUK_CLIENT`     | `podman` | The container client to use for starting and tearing down containers |
