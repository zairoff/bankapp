package transfer

import (
	"fmt"
)

func ExampleTotal() {

	res := Total(0)
	fmt.Println(res)

	res = Total(500000)
	fmt.Println(res)

	res = Total(1000000)
	fmt.Println(res)

	//Output:
	//0
	//502500
	//1005000

}
