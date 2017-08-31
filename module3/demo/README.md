# Demo Srv

This is the Demo service with fqdn com.brianketelsen.srv.demo.  This version has been modified to use NATS for broker, registry, and transport.

## Getting Started

### Prerequisites

Start NATS:
```
docker run -d -p 4222:4222  nats
``` 

### Run Service

```
go build
```

```
$ ./demo --registry=nats --registry_address=127.0.0.1:4222 --broker=nats --broker_address=127.0.0.1:4222 --transport=nats --transport_address=127.0.0.1:4222
```

### Building a container

If you would like to build the docker container do the following
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o demo-srv 
docker build -t demo-srv .

```