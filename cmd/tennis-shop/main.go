package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/xcheng85/Go-EDA/internal/config"
	"github.com/xcheng85/Go-EDA/internal/monolith"
	"github.com/xcheng85/Go-EDA/internal/swagger"
	"github.com/xcheng85/Go-EDA/internal/worker"
	"github.com/xcheng85/Go-EDA/players"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	// set up driven adapters
	// create deps
	config, err := config.NewAppConfig()
	if err != nil {
		return err
	}
	rpc := createGRPCServer()
	mux := createMux()
	workerSyncer := worker.NewSyncer()
	modules := []monolith.Module{
		&players.Module{},
	}
	// setup application
	// build the app with deps
	myapp := app{
		config:       config,
		modules:      modules,
		mux:          mux,
		rpc:          rpc,
		workerSyncer: workerSyncer,
	}
	// set up Driver adapters
	// bind rest and grpc routes
	if err = myapp.startupModules(); err != nil {
		return err
	}
	// Mount swagger
	mux.Mount("/", http.FileServer(http.FS(web.WebUI)))

	fmt.Println("started tennis-shop application")
	defer fmt.Println("stopped tennis-shop application")

	// blocking main thread
	myapp.workerSyncer.Add(
		myapp.runGRPC,
		myapp.runRest,
	)
	return myapp.workerSyncer.Sync()
}

// the following like ioc singletion
func createGRPCServer() *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)
	return server
}

func createMux() *chi.Mux {
	return chi.NewMux()
}
