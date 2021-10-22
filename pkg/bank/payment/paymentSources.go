package payment

import (
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.Paymentsource {

	temp := []types.Paymentsource{}
	count := 0

	for i := 0; i < len(cards); i++ {
		if cards[i].Active && cards[i].Balance > 0 {
			temp[count].Number = cards[i].Number
			temp[count].Type = cards[i].Type
			temp[count].Balance = cards[i].Balance
			count++
		}
	}
	return temp
}
