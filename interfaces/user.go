package interfaces

import (
	"context"

	"github.com/sah4ez/dapr-example/interfaces/types"
)

// @tg http-prefix=v1
// @tg jsonRPC-server log metrics
type User interface {
	GetNameByID(ctx context.Context, id types.ID) (user types.User, err error)
}
