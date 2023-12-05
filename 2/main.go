package main

import (
	"fmt"
	"regexp"
	"os"
	"bufio"
	"strings"
	"strconv"
)

var colorCounts = map[string]int{
	"red":12,
	"blue":14,
	"green":13,
}

func main() {
	f, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println("error opening file")
	}

	s := bufio.NewScanner(f) 

	//Part1(s)
	Part2(s)
}

func Part2(s *bufio.Scanner) {
	sum := 0 
	for s.Scan() {
		blocks := SplitLine(s.Text(), ",;:")
		blockMap := make(map[string]int)
		for i := 1; i < len(blocks); i++ {
			re := regexp.MustCompile("[0-9]+")
			count := re.FindString(blocks[i])

			color := ""
			for colors, _ := range colorCounts {
				if strings.Contains(blocks[i], colors) {
					color = colors
				}
			}
			
			currentCount := blockMap[color]
			countInt, _ := strconv.Atoi(count)

			if countInt > currentCount {
				blockMap[color] = countInt
			}
		}

		power := 1
		for _, v := range blockMap {
			power *= v
		}
		sum += power
	}

	fmt.Println(sum)
}

func Part1(s *bufio.Scanner) {
	sum := 0
	for s.Scan() {
		re := regexp.MustCompile("[0-9]+")
		gameId := re.FindString(s.Text())

		blocks := SplitLine(s.Text(), ",;:")
		
		if !Over(blocks) {
			gameIdInt, _ := strconv.Atoi(gameId)
			sum += gameIdInt
		}
	}

	fmt.Println(sum)
}

func SplitLine(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func Over(blocks []string) bool {
	for i := 1; i < len(blocks); i++ {
		re := regexp.MustCompile("[0-9]+")
		count := re.FindString(blocks[i])
		
		color := ""
		for colors, _ := range colorCounts {
			if strings.Contains(blocks[i], colors) {
				color = colors
			}
		}

		countInt, _ := strconv.Atoi(count)
		if countInt > colorCounts[color] {
			return true
		}
	}

	return false
}