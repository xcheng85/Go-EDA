# Applicaton architecture

1. classic monolith
2. modular monolith
3. microservices


code organization: streaming architecture

ports, adapters on hexagonal architecture
accepts interface, return structs 



two concpets: 

DIP: interface sits with consumer

Interface as contract: interface sits in the central location 


## grpc
one grpc server (tcp port) serves multiple grpc services: players, etc. 
no namespace conflicts. such as: package playerspb;


## module interaction

Query 1: all the products sponsered by player

Answer: cross-module interaction

INF --> Players.GetSponserProducts
INF --> Racquets.GetRacquetByPlayerId
INF <-- Racquets.GetRacquetByPlayerId
INF --> Bag.GetBagsByPlayerId
INF <-- Bag.GetBagsByPlayerId
...
INF <-- Players.GetSponserProducts

player module 

raquet module

bag module

shoes module


