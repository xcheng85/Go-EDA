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
2. bounded context


Domain-centric architecture


## CQRS
1. separation in the application layer
2. sepration in the db layer. (different db type)
3. seperation in the microservices

application in event sourcing
write: append-only log. redis stream
read: multiple read models

Task-based UI: command pattern, dispatch events