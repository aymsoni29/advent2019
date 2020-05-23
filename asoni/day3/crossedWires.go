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
func revertNumber(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findNearestIntersection(track1 [][]int, track2 [][]int) (int, int) {

	nearestIntersection := math.MaxInt64
	timingIntersection := math.MaxInt64

	sum1 := 0
	sum2 := 0

	for i := 0; i < len(track1)-1; i++ {
		point1 := track1[i]
		point2 := track1[i+1]

		// Check which index is common
		commonIndex1 := 0
		if point1[1] == point2[1] {
			commonIndex1 = 1
		}

		// Find Min and max of both co-ordinates
		min1 := findMin(point1[1-commonIndex1], point2[1-commonIndex1])
		max1 := findMax(point1[1-commonIndex1], point2[1-commonIndex1])
		commonVal1 := point1[commonIndex1]
		sum1 += max1 - min1
		sum2 = 0
		// For each line of wire 1, check intersection with each line of wire 2
		for j := 0; j < len(track2)-1; j++ {

			point3 := track2[j]
			point4 := track2[j+1]

			commonIndex2 := 0
			if point3[1] == point4[1] {
				commonIndex2 = 1
			}

			min2 := findMin(point3[1-commonIndex2], point4[1-commonIndex2])
			max2 := findMax(point3[1-commonIndex2], point4[1-commonIndex2])
			commonVal2 := point3[commonIndex2]
			sum2 += max2 - min2

			var intersect [2]int

			// These are parallel lines which intersect only when they overlap
			if commonIndex1 == commonIndex2 && point1[commonIndex1] == point3[commonIndex2] {
				// Handle overlapping lines
				globalMin := findMin(min1, min2)
				globalMax := findMax(max1, max2)
				for k := globalMin; k <= globalMax; k++ {
					if min1 <= k && k <= max1 && min2 <= k && k <= max2 {
						intersect[commonIndex2] = point3[commonIndex2]
						intersect[1-commonIndex2] = k
						sum1 = revertNumber(sum1)
						sum2 = revertNumber(sum2)
						fmt.Println("Sum1 ->", sum1, "Sum2 ->", sum2)
						if sum1+sum2 <= timingIntersection {
							timingIntersection = sum1 + sum2
						}
						break
					}
				}
			} else if commonIndex1 != commonIndex2 {
				var probableIntersection [2]int
				probableIntersection[commonIndex1] = commonVal1
				probableIntersection[commonIndex2] = commonVal2
				// Verify if the point lies on both the bounded lines
				if (min1 <= probableIntersection[commonIndex2] && probableIntersection[commonIndex2] <= max1) && (min2 <= probableIntersection[commonIndex1] && probableIntersection[commonIndex1] <= max2) {
					intersect[0] = probableIntersection[0]
					intersect[1] = probableIntersection[1]
					sum1 = revertNumber(sum1)
					sum2 = revertNumber(sum2)
					fmt.Println("Sum1 ->", sum1, "Sum2 ->", sum2)
					if sum1+sum2 <= timingIntersection {
						timingIntersection = sum1 + sum2
					}
				} else {
					continue
				}
			}

			intersect[0] = revertNumber(intersect[0])
			intersect[1] = revertNumber(intersect[1])
			sum1 = revertNumber(sum1)
			sum2 = revertNumber(sum2)

			if intersect[0]+intersect[1] != 0 {
				nearestIntersection = findMin(intersect[0]+intersect[1], nearestIntersection)
			}
			// fmt.Println()

		}
	}

	return nearestIntersection, timingIntersection
}

func main() {
	wires := parseInput("crossedWires_input.txt")
	track1 := trackWire(wires[0])
	track2 := trackWire(wires[1])
	fmt.Println(findNearestIntersection(track1, track2))

}
