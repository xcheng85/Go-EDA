# players module

## clean architecture

entities: domain

repositories: sql or no sql, CRUD, ORM, layer above db

services: business logic related to player

handlers (controllers): grpc and rest

DTO: request params, qs, body, and response, layer between domain and handler

## internal grpc module (controller)
grpc: like controller in asp.net core / node.js

Two deps:
1. service: business logic
2. actual grpc server, which is passed in from cmd app


```shell
go get -u "google.golang.org/grpc"
```



https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/creating_main.go/

Attach the Player service to the server passed in


## domain (entity)


## service (business logic)


## swagger 
1. embed (go standard lib)

index.html and api.swagger.json
url: "api.swagger.json",

The FS type enables embedding a tree of files, such as a directory of static web server content

2. register 


## Debug dependency graph
```shell
godepgraph ./cmd/tennis-shop/ | dot -Tpng -o godepgraph.png
```