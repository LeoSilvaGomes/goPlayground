package read

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	strLenght = 20
)

type Name struct {
	fname string
	lname string
}

func Execute() {
	var nameFile string

	fmt.Printf("Enter file name: ")
	fmt.Scan(&nameFile)

	f, _ := os.Open("assets/" + nameFile)

	scanner := bufio.NewScanner(f)
	names := make([]Name, 0)

	for scanner.Scan() {
		fullName := strings.Split(scanner.Text(), " ")

		if len(fullName[0]) > strLenght {
			fullName[0] = fullName[0][:strLenght]
		}

		if len(fullName[1]) > strLenght {
			fullName[1] = fullName[1][:strLenght]
		}

		names = append(names, Name{fullName[0], fullName[1]})
	}

	fmt.Printf("Slice complete: %s\n\n", names)

	fmt.Printf("Each struct:\n")
	for i := 0; i < len(names); i++ {
		fmt.Printf("Struct[%d]: %s %s\n", i, names[i].fname, names[i].lname)
	}

	f.Close()
}
