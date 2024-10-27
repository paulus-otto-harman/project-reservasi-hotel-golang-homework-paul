package view

import (
	"context"
)

type Screen interface {
	Render(ctx context.Context) int
}

func Render(screen Screen, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if screen.Render(ctx) == 0 {
				return
			}
		}
	}
}
