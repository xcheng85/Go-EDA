# Event Sourcing
Pattern

append-only stream

just like redis log file: to reconstruct state by read and exec command sequentially.

event: domain-event

history

in signle bounded context,

players, history events


## Implementations

event stores: 
1. consistency
2. concurrency control


### db selection

### serialize/deserialize

### event

1. event name
2. playload: anything use interface{}
3. event time

in go1.18 any alias for the interface{}

options ...EventOption: 
variadic parameters,
type is function type
fp example
variadic must be the last parameter


### aggregate

append event, hold all the events 


go 1.18 generics [T IEvent]
a generic satisfy an interface

save aggregate to db
load aggregate from db

versioning


## event store

event table in postgresSQL

compound key for concurrency write

short-lived aggregate
long-lived aggregate 


snapshot will improve performance when loading the aggregate events. leetcode snapshot binary search



appenevent --> insert into a table in postresql