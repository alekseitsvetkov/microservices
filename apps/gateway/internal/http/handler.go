package http

import (
	"example.com/microservices/apps/gateway/internal/grpc/clients"
)

type Handler struct {
	grpc *clients.Client
}

func NewHandler(grpc *clients.Client) *Handler {
	return &Handler{
		grpc: grpc,
	}
}
