package resolvers

import (
	"github.com/alekseytsvetkov/microservices/apps/gateway/internal/config"
	"github.com/alekseytsvetkov/microservices/apps/gateway/internal/http"
	authpb "github.com/alekseytsvetkov/microservices/proto/auth"
	productpb "github.com/alekseytsvetkov/microservices/proto/product"
	userpb "github.com/alekseytsvetkov/microservices/proto/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	cfg           *config.Config
	middleware    *http.Middleware
	userClient    userpb.UserServiceClient
	productClient productpb.ProductServiceClient
	authClient    authpb.AuthServiceClient
}

func NewResolver(
	cfg *config.Config,
	middleware *http.Middleware,
	userClient userpb.UserServiceClient,
	productClient productpb.ProductServiceClient,
	authClient authpb.AuthServiceClient,
) *Resolver {
	return &Resolver{
		cfg:           cfg,
		middleware:    middleware,
		userClient:    userClient,
		productClient: productClient,
		authClient:    authClient,
	}
}
