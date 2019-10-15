package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

type Middleware func(UsersService) UsersService

type loggingMiddleware struct {
	logger log.Logger
	next   UsersService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UsersService) UsersService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, email string) (e0 error) {
	defer func() {
		l.logger.Log("method", "Create", "email", email, "e0", e0)
	}()
	return l.next.Create(ctx, email)
}
