# GRPC-Gateway
https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/
https://github.com/grpc-ecosystem/grpc-gateway/tree/main


## Installation
create tools.go (https://github.com/grpc-ecosystem/grpc-gateway/tree/main)
for running go mod tidy

make install-tools

## Highlights
1. provide APIs in both gRPC and HTTP/JSON at the same time?
2. generates a reverse-proxy server which translates a RESTful HTTP API into gRPC. This server is generated according to the google.api.http annotations in your service definitions.

## PLayers

### create proto


### buf.yaml


### buf.gen.yaml
local plugin

By default, a protoc-gen-<name> program is expected to be on your PATH so that it can be discovered and executed by buf. 

## gRPC external config
https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/grpc_api_configuration/#using-an-external-configuration-file

