package permissions

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/sah4ez/dapr-example/interfaces"
	"github.com/sah4ez/dapr-example/interfaces/types"
)

var _ interfaces.Balance = (*BalanceAuth)(nil)

type BalanceAuth struct {
	next interfaces.Balance
}

func (p *BalanceAuth) GetBalance(ctx context.Context, id types.ID) (user types.Balance, err error) {
	log.Ctx(ctx).Info().Msg("check access")
	return p.next.GetBalance(ctx, id)
}

func NewBalance(svc interfaces.Balance) *BalanceAuth {
	return &BalanceAuth{
		next: svc,
	}
}
