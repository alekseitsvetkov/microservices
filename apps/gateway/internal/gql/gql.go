package gql

import (
	"example.com/microservices/apps/gateway/internal/gql/graph"
	"example.com/microservices/apps/gateway/internal/gql/resolvers"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func New(resolver *resolvers.Resolver) *handler.Server {
	cfg := graph.Config{
		Resolvers: resolver,
	}

	schema := graph.NewExecutableSchema(cfg)

	s := handler.New(schema)

	s.AddTransport(transport.Options{})
	s.AddTransport(transport.GET{})
	s.AddTransport(transport.POST{})

	s.Use(extension.Introspection{})

	return s
}
