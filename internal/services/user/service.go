package user

import (
	"context"
	"github.com/sah4ez/dapr-example/internal/models"
)

type Service struct {
	store iUserStorage
}

func New(store iUserStorage) *Service {
	return &Service{
		store: store,
	}
}

type iUserStorage interface {
	GetUser(ctx context.Context, id int) (user models.User, err error)
}
