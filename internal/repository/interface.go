package repository

import (
	"context"

	"github.com/sah4ez/dapr-example/internal/models"
)

type User interface {
	GetUser(ctx context.Context, id int) (user models.User, err error)
}
