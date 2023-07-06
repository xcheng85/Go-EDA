package monolith

import (
	"context"
)

type Monolith interface {
}

type Module interface {
	Startup(context.Context, Monolith) error
}
