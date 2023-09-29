package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CalcOptions struct {
	operator   string
	num1, num2 int
	isRoman    bool
}

func main() {
	fmt.Print("Enter math statement: ")
	input := getInput()
	options := parseInput(input)
	result := calc(options)
	outResult(result, options.isRoman)
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		gotError(errors.New("input error"))
	}
	str = strings.TrimSpace(str)

	return str
}

func parseInput(str string) CalcOptions {
	sep, err := getSep(str)
	if err != nil {
		gotError(err)
	}

	substr := strings.Split(str, sep)

	if len(substr) != 2 {
		gotError(fmt.Errorf("wrong math statement"))
	}

	left := strings.TrimSpace(substr[0])
	right := strings.TrimSpace(substr[1])
	num1, num2, isRoman := getNums(left, right)

	return CalcOptions{sep, num1, num2, isRoman}
}

func getNums(left, right string) (int, int, bool) {
	romanDigits := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	num1, isNum1Roman := romanDigits[left]
	num2, isNum2Roman := romanDigits[right]

	if isNum1Roman && isNum2Roman {
		return num1, num2, true
	}

	num1, errNum1 := strconv.Atoi(left)
	num2, errNum2 := strconv.Atoi(right)
	isNotValidNum := errNum1 != nil || errNum2 != nil || num1 == 0 || num2 == 0

	if isNotValidNum {
		gotError(fmt.Errorf("all numbers must be roman or arabic int between 1 and 10"))
	}

	return num1, num2, false
}

func getSep(str string) (string, error) {
	var sep string
	switch {
	case strings.Contains(str, "+"):
		sep = "+"
	case strings.Contains(str, "-"):
		sep = "-"
	case strings.Contains(str, "*"):
		sep = "*"
	case strings.Contains(str, "/"):
		sep = "/"
	default:
		return sep, fmt.Errorf("no valid operator")
	}

	return sep, nil
}

func calc(co CalcOptions) int {
	var result int
	switch co.operator {
	case "+":
		result = co.num1 + co.num2
	case "-":
		result = co.num1 - co.num2
	case "*":
		result = co.num1 * co.num2
	case "/":
		result = co.num1 / co.num2
	}

	return result
}

func arabicToRoman(num int) string {
	if num < 1 {
		gotError(fmt.Errorf("result in roman can't be less than I"))
	}

	arabicNums := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNums := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""

	for i, currArabic := range arabicNums {
		for num >= currArabic {
			result += romanNums[i]
			num -= currArabic
		}
	}

	return result
}

func outResult(res int, isRoman bool) {
	if isRoman {
		fmt.Println(arabicToRoman(res))
	} else {
		fmt.Println(res)
	}
}

func gotError(err error) {
	fmt.Println(err)
	os.Exit(0)
}
