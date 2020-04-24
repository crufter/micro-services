# Logspammer Service

This is the Logspammer service

Generated with

```
micro new logspammer --namespace=go.micro --type=service
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.logspammer
- Type: service
- Alias: logspammer

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./logspammer-service
```

Build a docker image
```
make docker
```