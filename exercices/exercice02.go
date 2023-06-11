package exercices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MultiplicationOfTenTable() string {
	scanner := bufio.NewScanner(os.Stdin)

	var number int
	var err error
	var text string

	fmt.Println("Insert a number to multiply by Ten: ")

	for {
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		text += fmt.Sprintf("%d x %d = %d \n", number, i, number*i)
	}

	return text
}
