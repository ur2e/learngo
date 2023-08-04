package main

import (
	"fmt"
)

func main() {
	var input int
	sum := 0
	fmt.Scanln(&input)

	for i := 1; ; i++ {
		sum += i

		if sum > input {
			fmt.Println(i - 1)
			break
		}
	}
}
