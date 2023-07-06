package players

import (
	"context"
	"github.com/xcheng85/Go-EDA/internal/monolith"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	return nil
}
