package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wires [][]string
	for scanner.Scan() {
		wire := strings.Split(scanner.Text(), ",")
		wires = append(wires, wire)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wires
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

// Find the movements of each wires and map them to co-ordinates
func trackWire(wire []string) [][]int {

	var track [][]int
	var position []int
	position = append(position, 0, 0)
	newPosition := append(position[:0:0], position...)
	track = append(track, newPosition)

	for _, val := range wire {

		direction := val[0]
		magnitude, err := strconv.Atoi(stripchars(val, "RULD"))

		if err != nil {
			fmt.Println("Error while parsing ", err)
		}

		switch direction {
		case 'R':
			position[0] += magnitude
		case 'L':
			position[0] -= magnitude
		case 'U':
			position[1] += magnitude
		case 'D':
			position[1] -= magnitude
		}

		newPosition := append(position[:0:0], position...)
		track = append(track, newPosition)
	}

	return track
}

// Find minimum between two integers
func findMin(a int, b int) (c int) {
	c = b
	if a < c {
		c = a
	}
	return
}

// Find maximum between two integers
func findMax(a int, b int) (c int) {
	c = b
	if a > c {
		c = a
	}
	return
}

// O(n^2) search between each lines and finding intersections
func revertNumber(x int) (int) {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	wires := parseInput("crossedWires_input.txt")
	track1 := trackWire(wires[0])
	track2 := trackWire(wires[1])
	fmt.Println(findNearestIntersection(track1, track2))

}
