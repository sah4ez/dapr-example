package balance

import (
	"context"

	"github.com/sah4ez/dapr-example/interfaces/types"
)

type Service struct {
	user iUser
}

func New(user iUser) *Service {
	return &Service{
		user: user,
	}
}

type iUser interface {
	GetNameByID(ctx context.Context, id types.ID) (user types.User, err error)
}
