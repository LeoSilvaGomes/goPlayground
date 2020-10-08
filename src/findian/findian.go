package findian

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Execute() {
	var str string
	fmt.Printf("Write a string that ends with n, begins with i and contains a: ")

	x := bufio.NewScanner(os.Stdin)
	if x.Scan() {
		str = x.Text()
	}

	str = strings.ToLower(str)

	if strings.HasPrefix(str, "i") && strings.Contains(str, "a") && strings.HasSuffix(str, "n") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}

}
