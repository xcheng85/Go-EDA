# internal module


## context 
see ch4 of book "cloud native go building reliable"

interface: context.Context

idiomatic for
1. deadlines
2. cancellation signals
3. requested-scoped values between processes (stored in the kv pairs )

context --> derived context: coordination throught done signal, sharing context
thread safe, goroutine

### create context
main function uses Background(), which should be derived
factory method
Context.WithCancel, WithTimeout


## Groups of Goroutines management
go get -u golang.org/x/sync/errgroup

https://pkg.go.dev/golang.org/x/sync/errgroup

Package errgroup provides synchronization, error propagation, and Context cancelation for groups of goroutines working on subtasks of a common task.

https://pkg.go.dev/golang.org/x/sync/errgroup

parallel/pipeline patterns for goroutine

1. Two ways to to ceate group

method1: new

method2: WithContext:
rules: canceled the first time a function passed to Go returns a non-nil error or the first time Wait returns, whichever occurs first.

2. Wait blocks


## gracefull shut down
signal.NotifyContext
go 1.16
https://millhouse.dev/posts/graceful-shutdowns-in-golang-with-signal-notify-context