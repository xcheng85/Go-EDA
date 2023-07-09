package main

import (
	"context"
	"github.com/xcheng85/Go-EDA/internal/monolith"
	"github.com/xcheng85/Go-EDA/internal/worker"
)

// application in the hexongal arch
type app struct {
	// modules just like all the application controllers
	modules []monolith.Module
	// management multiple goroutines
	workerSyncer worker.WorkerSyncer
}

// like the builder of ioc
func (a *app) startupModules() error {
	return nil
}

// worker for running gRPC
func (a *app) runGRPC(ctx context.Context) error {
	return nil
}
