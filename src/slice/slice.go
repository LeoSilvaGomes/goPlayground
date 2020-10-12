package slice

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Execute() {
	init := make([]int, 3)
	i := 0
	fmt.Printf("Write a integer to add a number in your sort slice, write 'X' to exit the program\n\n")

	for {
		var called string
		var num int

		fmt.Scan(&called)
		num, err := strconv.Atoi(called)

		if strings.ToLower(called) == "x" {
			break
		} else if err == nil {
			if i < 3 {
				switch {
				case init[0] == 0:
					init[0] = num
				case init[1] == 0:
					init[1] = num
				case init[2] == 0:
					init[2] = num
				}

				i++
			} else {
				init = append(init, num)
			}
			sort.Ints(init)

			fmt.Printf("Your sort slice: %v :P\n", init)
		} else {
			fmt.Printf("The slice accept just integer or 'X' to exit the program\n\n")
		}

	}
}
