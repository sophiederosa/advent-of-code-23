package main

import (
	"fmt"
	"bufio"
	"os" 
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}

	s := bufio.NewScanner(f)
	sum := 0
	for s.Scan() {
		re := regexp.MustCompile("[0-9]")
		allDigits := re.FindAllString(s.Text(), -1)

		firstDigit := allDigits[0]
		lastDigit := allDigits[len(allDigits) - 1]

		total := firstDigit + lastDigit
		intTotal, _ := strconv.Atoi(total)
		sum += intTotal
	}

	fmt.Println(sum)
}