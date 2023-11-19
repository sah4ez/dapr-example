package permissions

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/sah4ez/dapr-example/interfaces"
	"github.com/sah4ez/dapr-example/interfaces/types"
)

var _ interfaces.User = (*UserAuth)(nil)

type UserAuth struct {
	next interfaces.User
}

func (p *UserAuth) GetNameByID(ctx context.Context, id types.ID) (user types.User, err error) {
	log.Ctx(ctx).Info().Msg("check access")
	return p.next.GetNameByID(ctx, id)
}

func NewUser(svc interfaces.User) *UserAuth {
	return &UserAuth{
		next: svc,
	}
}
