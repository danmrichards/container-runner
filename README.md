# Container Runner
A quick and dirty demonstration on the way in which workloads can be run inside
a container.

This repo consists of two applications, the server and the runner. The server
is just a dummy application that acts as a workload to put inside the container. 
The runner is the tool that can place the workload in a container using one of
it's supported runtimes.

## Requirements
* Go 1.13+
* Docker

## Building From Source

Clone this repo and build the binaries:

```bash
$ make build
```

Ensure you also build the rootfs image:

```bash
$ make rootfs.tar
```

## Dependencies
### Go
Update the Go dependencies like so:

```bash
$ make deps
```

## Usage

```bash
Usage of ./bin/server-linux-amd64:
  -host string
    	host for the HTTP clients (default "0.0.0.0")
  -port string
    	port for the HTTP clients (default "8080")
```
