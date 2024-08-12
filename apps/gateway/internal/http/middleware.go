package http

import (
	"context"
	"errors"
	"github.com/alekseytsvetkov/microservices/apps/gateway/internal/config"
	"github.com/alekseytsvetkov/microservices/libs/jwt"
	"github.com/alekseytsvetkov/microservices/libs/respond"
	userpb "github.com/alekseytsvetkov/microservices/proto/user"
	"net/http"
	"strings"
)

const ID_KEY = "id"

type Middleware struct {
	cfg        *config.Config
	userClient userpb.UserServiceClient
}

func NewMiddleware(cfg *config.Config, userClient userpb.UserServiceClient) *Middleware {
	return &Middleware{
		cfg:        cfg,
		userClient: userClient,
	}
}

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			next.ServeHTTP(w, r)
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			respond.Error(w, http.StatusUnauthorized, ErrInvalidFormat)
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		claims, err := jwt.Validate(token, m.cfg.JWT.Secret)
		if err != nil {
			if errors.Is(err, jwt.ErrInvalidToken) {
				respond.Error(w, http.StatusUnauthorized, ErrInvalidToken)
				return
			}

			respond.Error(w, http.StatusForbidden, ErrAccessDenied)
			return
		}

		ctx := context.WithValue(r.Context(), ID_KEY, claims.ID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) GetUserIDFromCtx(ctx context.Context) string {
	userID, ok := ctx.Value(ID_KEY).(string)
	if !ok {
		return ""
	}

	return userID
}
