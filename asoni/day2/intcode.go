package day2

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// TODO: Move to utils package
func readCsvFile(filePath string) []string {
	f, _ := os.Open(filePath)
	defer f.Close()
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
		data = record
	}
	return data
}

// TODO: Move to utils package
func parseData(stringData []string) []int {
	var intData []int
	for _, val := range stringData {
		x, _ := strconv.Atoi(val)
		intData = append(intData, x)
	}
	return intData
}

func runOpCode(index int, intData *[]int) {
	// TIL That is how you reference pointers - doing intData[index] will result in an error (cannot index variable of type *[]int)
	// For more info, see - https://flaviocopes.com/golang-does-not-support-indexing/
	switch (*intData)[index] {

	// Opcode list
	case 1:
		(*intData)[(*intData)[index+3]] = (*intData)[(*intData)[index+1]] + (*intData)[(*intData)[index+2]]
		return

	case 2:
		(*intData)[(*intData)[index+3]] = (*intData)[(*intData)[index+1]] * (*intData)[(*intData)[index+2]]
		return

	case 99:
		return

	default:
		fmt.Println("Encountered unknown opcode")
		return
	}
}

func runIntcode(intData []int) []int {
	for i := 0; i < len(intData); i += 4 {
		// Halt Program
		if intData[i] == 99 {
			return intData
		}
		runOpCode(i, &intData)
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
