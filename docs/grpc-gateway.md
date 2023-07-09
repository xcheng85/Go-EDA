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


## Architecture

Each domain is a module, and will be the adapters

port is in the shared domain. 

core domain,
supporting domain,
application domain

in the cmd, main app, 

app: just like webbuilder in asp.net core, 

application, driven, driver




## PLayers Module

### interface of the module (port)
port
multiple adapters: player is one of the adapters.

players/module.go


### player domain and bounded context


### player domain model






### create proto


### buf.yaml


### buf.gen.yaml
local plugin

By default, a protoc-gen-<name> program is expected to be on your PATH so that it can be discovered and executed by buf. 

## gRPC external config
https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/grpc_api_configuration/#using-an-external-configuration-file

api.annotations.yaml
selector: <pkg>.<servicename>.<rpcMethod>

using an OpenAPI configuration file
example: https://github.com/grpc-ecosystem/grpc-gateway/blob/main/examples/internal/proto/examplepb/unannotated_echo_service.swagger.yaml


this will generate swagger.json

## start with buf
WARN    Specified deps are not covered in your buf.lock, run "buf mod update":
        - buf.build/googleapis/googleapis
        - buf.build/grpc-ecosystem/grpc-gateway
buf mod update
buf build
echo $?

buf generate

it will create 