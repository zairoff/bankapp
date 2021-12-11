package main

import (
	"bank/pkg/bank/transfer"
	"fmt"
)

func ExampleTotal() {

	res := transfer.Total(0)
	fmt.Println(res)

	res = transfer.Total(500000)
	fmt.Println(res)

	res = transfer.Total(1000000)
	fmt.Println(res)
}
