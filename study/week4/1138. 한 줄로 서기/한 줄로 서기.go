package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var left int
		fmt.Scan(&left)
		for j := 1; j <= n; j++ {
			if left == 0 && arr[j] == 0 {
				arr[j] = i
				break
			} else if arr[j] == 0 {
				left--
			}
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
