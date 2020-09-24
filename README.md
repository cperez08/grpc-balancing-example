# gRPC balancing example

This repo contains different examples about implementing golang gRPC server/client in Kubernetes and shows different approaches on how to achieve a proper load balancing.

### Demo 1

check [demo 1](./cmd/demo1) and [k8s demo1](./k8s/demo1) folders to see the a common default configuration.


### Demo 2

check [demo2](./cmd/demo2) and [k8s demo2](./k8s/demo2) folders to see the a tweaked configuration using the grpc DNS resolver and a headless servic.

### Demo 3

check [demo3](./cmd/demo3) and [k8s demo3](./k8s/demo3) folders to see the implementation of a custom gRPC resolver also the usage of a custom gRPC proxy along with the headless service.
