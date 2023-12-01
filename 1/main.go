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
	re := regexp.MustCompile("[0-9]")
	
	sum := 0
	for s.Scan() {
		allDigits := re.FindAllString(s.Text(), -1)

		firstDigit := allDigits[0]
		lastDigit := allDigits[len(allDigits) - 1]

		total := firstDigit + lastDigit
		intTotal, _ := strconv.Atoi(total)
		sum += intTotal
	}

	fmt.Println(sum)
}