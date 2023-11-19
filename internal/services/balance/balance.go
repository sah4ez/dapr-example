package balance

import (
	"context"

	"github.com/sah4ez/dapr-example/interfaces/types"
	"github.com/sah4ez/dapr-example/pkg/errors"
	"github.com/shopspring/decimal"
)

func (svc *Service) GetBalance(ctx context.Context, id types.ID) (balance types.Balance, err error) {

	userID := 123
	user, err := svc.user.GetNameByID(ctx, userID)
	if err != nil {
		err = errors.BalanceUserErr.SetCause("userID: %d", userID)
		return
	}

	balance = types.Balance{
		Amount:   decimal.New(123, 0),
		UserName: user.Name,
		UserID:   userID,
	}

	return
}
