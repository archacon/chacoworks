package service

import (
	"context"
)

// BugsService describes the service.
type BugsService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Create(ctx context.Context, bug string) error
}
