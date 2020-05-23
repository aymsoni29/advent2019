package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)
const (
	inputValue = 5
)

func readCsvFile(filePath string) []string {
	// Load a csv file.
	f, _ := os.Open(filePath)
	defer f.Close()
	// Create a new reader.
	r := csv.NewReader(f)
	var data []string
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		data = record
	}
	fmt.Println(data)
	return data
}

func parseData(stringData []string) []int {
	var intData []int
	for _, val := range stringData {
		x, _ := strconv.Atoi(val)
		intData = append(intData, x)
	}
	return intData
}

func runOpCode(ind int, i *[]int) int {

	// TIL That is how you reference pointers - doing intData[index] will result in an error (cannot index variable of type *[]int)
	// For more info, see - https://flaviocopes.com/golang-does-not-support-indexing/

	opCode := (*i)[ind]%100

	switch (*i)[ind] {

	// Opcode list
	case 1:
		(*i)[(*i)[ind+3]] = (*i)[(*i)[ind+1]] + (*i)[(*i)[ind+2]]
		return 4

	case 2:
		(*i)[(*i)[ind+3]] = (*i)[(*i)[ind+1]] * (*i)[(*i)[ind+2]]
		return 4

	case 3:
		(*i)[(*i)[ind+1]] = inputValue
		return 2
	
	case 4:
		fmt.Println("Output:", (*i)[(*i)[ind+1]])
		return 2

	case 99:
		return 4

	default:
		fmt.Println("Encountered unknown opcode")
		return 4
	}
}

func runIntcode(intData []int) []int {
	var step int = 4
	for i := 0; i < len(intData); i += step {
		// Halt Program
		if intData[i] == 99 {
			return intData
		}
		step = runOpCode(i, &intData)
	}
	return intData
}

func part2(intData []int) (int, int) {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {

			newIntData := make([]int, len(intData))
			copy(newIntData, intData)

			newIntData[1] = i
			newIntData[2] = j

			runIntcode(newIntData)
			if newIntData[0] == 19690720 {
				return i, j
			}
		}
	}
	return 0, 0
}

func main() {
	stringData := readCsvFile("day-2-input.csv")
	intData := parseData(stringData)
	fmt.Println(intData[99])
	// intData = runIntcode(intData)
	// fmt.Println("Day 2 Part 1:", intData[0])
	x, y := part2(intData)
	fmt.Println("Day 2 Part 2:", ((100 * x) + y))
}
