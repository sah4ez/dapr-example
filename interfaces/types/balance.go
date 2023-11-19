package types

import (
	"github.com/shopspring/decimal"
)

type Balance struct {
	ID ID `json:"id"`

	Amount   decimal.Decimal `json:"amount"`
	UserName string          `json:"userName"`
	UserID   ID              `json:"userID"`
}
