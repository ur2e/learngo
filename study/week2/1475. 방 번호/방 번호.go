package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var s [10]int
	fmt.Scanf("%s", &input)

	for i := 0; i < len(input); i++ {

		// input[i] = []byte
		num, _ := strconv.Atoi(string(input[i]))
		if num == 6 || num == 9 {
			if s[6] <= s[9] {
				s[6]++
			} else {
				s[9]++
			}
		} else {
			s[num]++
		}
	}

	max := 0
	for _, cnt := range s {
		if cnt > max {
			max = cnt
		}
	}

	fmt.Println(max)
}
