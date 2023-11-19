package user

import (
	"context"

	"github.com/sah4ez/dapr-example/interfaces/types"
	"github.com/sah4ez/dapr-example/internal/models"
)

func (svc *Service) GetNameByID(ctx context.Context, id types.ID) (user types.User, err error) {
	mUser, err := svc.store.GetUser(ctx, id)
	if err != nil {
		return
	}

	user = userFromModel(mUser)
	return
}

func userFromModel(user models.User) types.User {
	return types.User{
		ID:   user.ID,
		Name: user.Name,
	}
}
