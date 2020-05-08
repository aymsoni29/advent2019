package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Each new line has the input parameter x
func parseInputFile(filename string) []int {
	var data []int
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error in parsing")
		}
		data = append(data, val)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return data
}

// Consume input data and calculate fuel
func calculateFuel(mass int) int {
	return (mass / 3) - 2
}

func day1Part1(massData []int) (fuel int) {
	fuel = 0
	for _, val := range massData {
		fuel += calculateFuel(val)
	}
	return fuel
}

func day1Part2(massData []int) (fuel int) {
	fuel = 0
	for _, val := range massData {
		x := calculateFuel(val)
		for x >= 0 {
			fuel += x
			x = calculateFuel(x)
		}
	}
	return fuel
}

func main() {

	data := parseInputFile("day-1-input.txt")
	// totalFuel := day1Part1(data)
	totalFuel := day1Part2(data)
	fmt.Println("Total Fuel =", totalFuel)
}
