# Golang Sandbox

This repository is aimed to help me learn Golang in a Docker container

## Prerequisites

Infrastructure wise you'll need:
* Docker environment reachable by the docker client, I recommend installing [Docker Toolbox](https://www.docker.com/toolbox)
* Being in a Unix-like system which has `make` available

## Start development

The make command will start the development environment and run the `make dev` and `make logs` targets

```bash
make
```

## Other commands

### Accessing the container

```bash
make enter

  Write the name of the service that you want to access, possible choices are:

  golang

  Service name: redis
root@b9616088331e:/#
```

### Viewing logs

```bash
make logs
```

### Stopping it

```bash
make kill
```

### Restarting it

```bash
make restart
```

### Removing everything

It will also stop the environment if its running

```bash
make nodev
```

## Folders

### ops/

This folder will enclose all of the System Operations information, and code necessary to get the application up and running.
