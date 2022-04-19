package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := getKangaroosParams()

	if err != nil {
		fmt.Print(err)
		return
	}

	var (
		x1 = (*data)[0]
		x2 = (*data)[2]
		v1 = (*data)[1]
		v2 = (*data)[3]
	)

	// Second kangaroo is too fast
	if x2 > x1 && v2 > v1 {
		fmt.Print("No\n")
		return
	}

	for x2 > x1 {
		x2 = x2 + v2
		x1 = x1 + v1

		if x1 == x2 {
			fmt.Print("Yes\n")
		}

		if x1 > x2 {
			fmt.Print("No\n")
		}
	}
}

func checkIfNumbers(num []string, input string) bool {
	return len(num) == len(strings.Fields(input))
}

func getKangaroosParams() (*[]int, error) {
	kangaroosParam := make([]int, 4)

	fmt.Print("Enter 4 digits: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("an error occurred while reading input. Please try again")
	}
	input = strings.TrimSuffix(input, "\n")

	numsCollection := regexp.MustCompile("[0-9]+").FindAllString(input, -1)

	if !checkIfNumbers(numsCollection, input) {
		return nil, errors.New("incorrect input")
	}

	if len(numsCollection) != 4 {
		return nil, errors.New("please enter 4 digits")
	}

	for i, num := range numsCollection {
		kangaroosParam[i], _ = strconv.Atoi(num)
	}

	return &kangaroosParam, nil
}
