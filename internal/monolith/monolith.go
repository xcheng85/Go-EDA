package monolith

import (
	"context"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"github.com/xcheng85/Go-EDA/internal/config"
	"github.com/xcheng85/Go-EDA/internal/worker"
)

// define the interface of di
// Module requires Monolith interface

// chi.Mux is the implementation of chi.Router interface
type Monolith interface {
	Config() config.AppConfig
	Mux() *chi.Mux
	RPC() *grpc.Server
	WorkerSyncer() worker.WorkerSyncer
}

type Module interface {
	Startup(context.Context, Monolith) error
}
