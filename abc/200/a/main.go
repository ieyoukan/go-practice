package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	//x/100 切上げ的な？
	isamari := 0
	if x%100 != 0 {
		isamari = 1
	}
	fmt.Println(x/100 + isamari)
}
