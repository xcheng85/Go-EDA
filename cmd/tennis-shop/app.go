package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/xcheng85/Go-EDA/internal/config"
	"github.com/xcheng85/Go-EDA/internal/monolith"
	"github.com/xcheng85/Go-EDA/internal/worker"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"time"
	"net/http"
)
// composition root
// application in the hexongal arch
// app must implement monolith interface, which is required in each sub module
type app struct {
	config config.AppConfig
	// modules just like all the application controllers
	modules []monolith.Module
	mux     *chi.Mux
	rpc     *grpc.Server
	// management multiple goroutines
	workerSyncer worker.WorkerSyncer
}

func (a *app) Config() config.AppConfig {
	return a.config
}

func (a *app) Mux() *chi.Mux {
	return a.mux
}

func (a *app) RPC() *grpc.Server {
	return a.rpc
}

func (a *app) WorkerSyncer() worker.WorkerSyncer {
	return a.workerSyncer
}

// like the builder of ioc
func (a *app) startupModules() error {
	for _, module := range a.modules {
		if err := module.Startup(a.workerSyncer.Context(), a); err != nil {
			return err
		}
	}
	return nil
}

// worker for running Rest server for reverse proxy
func (a *app) runRest(ctx context.Context) error {
	// chi.Mux has the http.Handler embedded
	restServer := &http.Server{
		Addr:    a.config.Web.Address(),
		Handler: a.mux,
	}

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Println("web server started")
		defer fmt.Println("web server shutdown")
		if err := restServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	group.Go(func() error {
		// received cancel signal from the derived
		<-gCtx.Done()
		fmt.Println("web server to be shutdown")
		// gracefully shut down rest server
		ctx, cancel := context.WithTimeout(context.Background(), a.config.ShutdownTimeout)
		defer cancel()
		if err := restServer.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})
	// block here
	return group.Wait()
}

// worker for running gRPC
func (a *app) runGRPC(ctx context.Context) error {
	// 	lis, err := net.Listen("tcp", ":8080")
	// grpc server requires a tcp port
	listener, err := net.Listen("tcp", a.config.Rpc.Address())
	if err != nil {
		return err
	}

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Println("rpc server started")
		defer fmt.Println("rpc server shutdown")
		if err := a.RPC().Serve(listener); err != nil && err != grpc.ErrServerStopped {
			return err
		}
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		fmt.Println("rpc server to be shutdown")
		// create bi-rectional channel
		stopped := make(chan struct{})
		go func() {
			a.RPC().GracefulStop()
			close(stopped)
		}()
		// after timecout, it will send ts to the channel timout.C
		timeout := time.NewTimer(a.config.ShutdownTimeout)
		select {
		case <-timeout.C:
			// Force it to stop
			a.RPC().Stop()
			return fmt.Errorf("rpc server failed to stop gracefully")
		case <-stopped:
			return nil
		}
	})

	return group.Wait()
}
