package directives

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/alekseytsvetkov/microservices/apps/gateway/internal/http"
)

type Directive struct {
	middleware *http.Middleware
}

func NewDirective(middleware *http.Middleware) *Directive {
	return &Directive{
		middleware: middleware,
	}
}

func (c *Directive) IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	userID := c.middleware.GetUserIDFromCtx(ctx)
	if userID == "" {
		return nil, http.ErrUnauthorized
	}

	return next(ctx)
}
