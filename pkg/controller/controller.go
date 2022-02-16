package controller

import (
	"context"

	"github.com/rohandas-max/ghCrwaler/pkg/handler"
)

func Controller(ctx context.Context, username string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if err := handler.Handler(ctx, username); err != nil {
			return err
		}
	}

	return nil
}
