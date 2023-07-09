package main

import (
	"fmt"
	"os"
	"github.com/xcheng85/Go-EDA/internal/worker"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	myapp := app{}
	myapp.workerSyncer = worker.NewSyncer()

	return nil
}
