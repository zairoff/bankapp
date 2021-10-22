package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{ID: 1, Amount: 99},
		{ID: 2, Amount: 19},
		{ID: 8, Amount: 99},
		{ID: 4, Amount: 39},
		{ID: 5, Amount: 19},
	}

	max := Max(payments)

	fmt.Println(max)

	//Output: {1 99}

}
