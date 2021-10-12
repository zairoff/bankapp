package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

// func ExampleWithdraw_positive() {
// 	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
// 	fmt.Println(result.Balance)
// 	// Output: 1000000
// }

// func ExampleWithdraw_noMoney() {
// 	Withdraw(types.Card{Balance: 20_000_00, Active: true}, 30_000_00)

// 	// Output: Операция невозвможна
// }

// func ExampleWithdraw_inactive() {
// 	Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)

// 	// Output: Операция невозвможна
// }

// func ExampleWithdraw_limit() {
// 	Withdraw(types.Card{Balance: 20_000_00, Active: false}, 20_000_00)

// 	// Output: Операция невозвможна
// }

func ExampleDeposit() {
	var card = types.Card{Balance: 2000000, Active: true}
	Deposit(&card, 1000000)
	fmt.Println(card.Balance)

	card.Active = false
	Deposit(&card, 1000000)
	fmt.Println(card.Balance)

	card.Active = true
	Deposit(&card, 10000000)
	fmt.Println(card.Balance)

	//Output: 3000000
	//3000000
	//3000000

}
