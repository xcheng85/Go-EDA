# Go-EDA
Golang with Event-driven arch


## EDA
Benifits:
1. Resiliency: component offline
2. Agility: development standpoint
3. UX updates
4. Analytics and auditing

Challenges:
1. Eventual consistency: distributed nature
2. Dual writes: lose of events
3. Distributed and async workflow: saga Choreography and Orchestration
4. Debuggability




1. async
2. distributed
3. loosely coupled

three patterns:
1. event notification
2. event-carried state transfer (Rest)
3. event-sourcing, event store


message queue: no retention
stream: retention and replay events. 
event store: append-only 

producers
consumers: groups, 


## Arch

1. Hexagonal architecture

inputs: API and event subs
outputs: event pub, db

ports: interface
adapters: implementation
dependency inversion principle


think of asp.net core:

layers:

domain layer: domain model, services
appliation layer: services and interfaces(port) (this layer depends)


driver (input to applicatio layer): webui, api, event consumer
driven (applicaiton layer trigger): db, logging, event producer''



DDD:
1. model
2. bounded context: core and supporting

Domain-centric architecture


## CQRS
1. separation in the application layer
2. sepration in the db layer. (different db type)
3. seperation in the microservices

application in event sourcing
write: append-only log. redis stream
read: multiple read models

Task-based UI: command pattern, dispatch events


## Event Storming

### Big Picture

all the domain events of the application
source of the events
correct chronological order of events

people using the application: admin, customer,bot, ...

external system: services


### Design Level
single bounded context
no domain expert
goal: design and process building

figure out what is the core and support context. 

Tools: Gherkin and Godog tools


### Architecture Decision
Markdown 
in git


### screaming architecture
for code organization
folder name: bounded context,
docs
docker
cmd
internal

## Useful tools
https://github.com/grpc-ecosystem/grpc-gateway

2. grpc
https://github.com/bufbuild/buf
https://github.com/bufbuild/buf-examples


## Developer replay
for each api service
1. Generates OpenAPI definitions for Protobuf services.
2. 


### 1. Player bounded context
```shell
go mod init github.com/xcheng85/Go-EDA

mkdir -p players

# buf generate protobuf
# buf remote plugin:
# 1. openapiv2

```