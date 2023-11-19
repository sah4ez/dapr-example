package interfaces

import (
	"context"

	"github.com/sah4ez/dapr-example/interfaces/types"
)

// @tg http-prefix=v1
// @tg jsonRPC-server log metrics
type Balance interface {
	GetBalance(ctx context.Context, id types.ID) (balance types.Balance, err error)
}
