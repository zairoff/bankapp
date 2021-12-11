package payment

import (
	"bank/pkg/bank/types"
)

func Max(payments []types.Payment) types.Payment {
	temp := 0
	for i := 1; i < len(payments); i++ {
		if payments[i].Amount > payments[temp].Amount {
			temp = i
		}
	}
	return payments[temp]
}
