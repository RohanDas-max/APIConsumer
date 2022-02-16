package controller

import (
	"context"
	"time"

	"github.com/rohandas-max/ghCrwaler/pkg/handler"
)

func Controller(ctx context.Context, username string, t time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(t):
		if err := handler.Handler(ctx, username); err != nil {
			return err
		}
	}
	return nil
}
