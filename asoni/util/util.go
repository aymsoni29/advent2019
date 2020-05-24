package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ParseInputFile - Each new line has the input parameter x
func ParseInputFile(filename string) []int {
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
