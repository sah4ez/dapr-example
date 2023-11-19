package postgres

import (
	"context"

	"github.com/sah4ez/dapr-example/internal/models"
)

type Mock struct{}

func (Mock) GetUser(ctx context.Context, id int) (user models.User, err error) {
	return (models.User{}).Mock(), nil
}
