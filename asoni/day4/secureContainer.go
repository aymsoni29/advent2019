package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Optimization
func getDigits(x int) [6]int {
	var digits [6]int
	for i := 0; i < 6; i++ {
		digits[5-i] = x % 10
		x = x / 10
	}
	return digits
}

func verifyDigits(x int) bool {
	for i := 0; i < 5; i++ {
		x = x / 10
	}
	if x != 0 && x/10 == 0 {
		return true
	}
	return false
}

func adjacentDigits(digits [6]int) bool {
	prev := -1
	for _, val := range digits {
		if prev != -1 && val == prev {
			return true
		}
		prev = val
	}
	return false
}

func verifyNonDecreasing(digits [6]int) bool {
	prev := -1
	for _, val := range digits {
		if prev != -1 && prev > val {
			return false
		}
		prev = val
	}
	return true
}

func calculateAnswer(min int, max int) (count int) {
	count = 0
	for i := min; i <= max; i++ {
		if verifyDigits(i) {
			digits := getDigits(i)
			if adjacentDigits(digits) && verifyNonDecreasing(digits) {
				count++
			}
		}
	}
	return
}

func matchingDigitFilter(digits [6]int) bool {
	var digitCount map[int]int
	digitCount = make(map[int]int)
	for _, val := range digits {
		digitCount[val]++
	}

	for _, v := range digitCount {
		if v == 2 {
			return true
		}
	}
	return false
}

func calculateAnswerWithFilter(min int, max int) (count int) {
	count = 0
	for i := min; i <= max; i++ {
		if verifyDigits(i) {
			digits := getDigits(i)
			if adjacentDigits(digits) && verifyNonDecreasing(digits) && matchingDigitFilter(digits) {
				count++
			}
		}
	}
	return
}
func main() {
	fmt.Println("Day 4: Secure Container")

	// input := "278384-824795"
	input := "172851-675869"
	inputRange := strings.Split(input, "-")
	min, _ := strconv.Atoi(inputRange[0])
	max, _ := strconv.Atoi(inputRange[1])

	fmt.Println("Part 1", calculateAnswer(min, max))
	fmt.Println("Part 2", calculateAnswerWithFilter(min, max))

}

// func verifyNonDecreasing(x int) bool {
// 	prev := -1
// 	for i := 0; i < 6; i++ {
// 		digit := x % 10
// 		if prev != -1 && digit > prev {
// 			return false
// 		}
// 		prev = digit
// 		x = x / 10
// 	}
// 	return true
// }

// func adjacentDigits(x int) bool {
// 	prev := -1

// 	for i := 0; i < 6; i++ {
// 		digit := x % 10
// 		if prev == digit {
// 			return true
// 		}
// 		prev = digit
// 		x = x / 10
// 	}
// 	return false
// }
