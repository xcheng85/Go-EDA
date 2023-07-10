package monolith

import (
	"context"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

// define the interface of di
// Module requires Monolith interface

// chi.Mux is the implementation of chi.Router interface
type Monolith interface {
	Mux() *chi.Mux
	RPC() *grpc.Server
}

type Module interface {
	Startup(context.Context, Monolith) error
}
