package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	bridge := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&bridge[i])
	}

	var start, destination int
	fmt.Scan(&start, &destination)

	dq := make([]int, 0)
	dq = append(dq, start-1)
	check := make([]int, N)
	for i := 0; i < N; i++ {
		check[i] = -1
	}
	check[start-1] = 0

	var breakCheck bool

	for len(dq) > 0 {
		now := dq[0]
		dq = dq[1:]
		for n := 0; n < N; n++ {
			if (n-now)%bridge[now] == 0 && check[n] == -1 {
				dq = append(dq, n)
				check[n] = check[now] + 1
				if n == destination-1 {
					fmt.Println(check[n])
					breakCheck = true
					break
				}
			}
		}
		if breakCheck {
			break
		}
	}

	if check[destination-1] == -1 {
		fmt.Println(-1)
	}
}