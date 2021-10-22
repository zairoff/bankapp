package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}

	return card
}

func Withdraw(card types.Card, amount types.Money) types.Card {
	if !card.Active || card.Balance < amount || amount < 0 || amount > 2000000 {
		fmt.Println("Операция невозвможна")
		return card
	}
	card.Balance = card.Balance - amount
	return card
}

func Deposit(card *types.Card, amount types.Money) {
	if !card.Active || amount > 5000000 || amount < 0 {
		return
	}
	card.Balance += amount

}

func Total(cards []types.Card) types.Money {
	var temp types.Money
	for i := 0; i < len(cards); i++ {
		temp += cards[i].Balance
	}
	return temp
}
