package main

import (
	"github.com/xcheng85/Go-EDA/internal/monolith"
)

// application in the hexongal arch
type app struct {
	// modules just like all the application controllers
	modules []monolith.Module
}

// like the builder of ioc
func (a *app) startupModules() error {
	return nil
}
