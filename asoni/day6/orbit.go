package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "https://github.com/sean-hart/advent2019/asoni/util"
)

// parseInputFile - Each new line has the input parameter x
func parseInputFile(filename string) []string {
	var data []string
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Failed opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
		if err != nil {
			fmt.Println("Error in parsing")
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return data
}

// We collect each node and all it's connections
func populateGraph(data []string) (map[string][]string, map[string]bool) {

	graph := make(map[string][]string)
	nodes := make(map[string]bool)

	for _, val := range data {

		pair := strings.Split(val, ")")

		// Add a connection to that node
		if val, ok := graph[pair[0]]; ok {
			val = append(val, pair[1])
			graph[pair[0]] = val
		} else {
			graph[pair[0]] = []string{pair[1]}
		}

		// Collect all the nodes while traversing
		nodes[pair[0]] = true
		nodes[pair[1]] = true
	}

	return graph, nodes
}

// For each node find it's direct+indirect connections
func calculateNodeChecksum(graph map[string][]string, nodeChecksum map[string]int, node string) int {

	// Base case when the node is a leaf node
	if _, ok := graph[node]; !ok {
		nodeChecksum[node] = 0
		return 0
	}

	// Another base case when we already have the checksum of the no
	// This won't work if we have cycles in the graph
	if _, ok := nodeChecksum[node]; ok {
		return nodeChecksum[node]
	}

	// Iterate through all it's edges
	for _, n := range graph[node] {

		// For each edge increment the checksum of node by 1 as it is a direct orbit
		if _, ok := nodeChecksum[node]; ok {
			nodeChecksum[node]++
		} else {
			nodeChecksum[node] = 1
		}

		// Check if we already calculated the checksum of the edge node
		if val, ok := nodeChecksum[n]; ok {
			nodeChecksum[node] += val
		} else {
			nodeChecksum[node] += calculateNodeChecksum(graph, nodeChecksum, n)
		}
	}

	return nodeChecksum[node]
}

// For all nodes, add their checksums
func calculateTotalChecksum(graph map[string][]string, nodes map[string]bool) int {

	totalChecksum := 0
	nodeChecksum := make(map[string]int)
	for node := range nodes {
		totalChecksum += calculateNodeChecksum(graph, nodeChecksum, node)
	}
	return totalChecksum
}

// Find all ancestors of YOU and SAN
func getAncestors(graph map[string][]string, node string, ancestors []string) []string {
	parent := ""
	for key, val := range graph {
		for _, x := range val {
			if x == node {
				ancestors = append(ancestors, key)
				parent = key
				break
			}
		}
		if parent != "" {
			break
		}
	}
	if parent != "" {
		ancestors = getAncestors(graph, parent, ancestors)
	}
	return ancestors
}

// Find first common ancestor
func minimumOrbitChange(you []string, san []string) int {
	count1 := 0
	count2 := 0
	for _, node := range you {
		count2 = 0
		for _, node2 := range san {
			if node == node2 {
				return count1 + count2
			}
			count2++
		}
		count1++
	}
	return -1
}

func main() {

	fmt.Println("Day 6")
	data := parseInputFile("input.txt")
	graph, nodes := populateGraph(data)
	fmt.Println(calculateTotalChecksum(graph, nodes))
	you := getAncestors(graph, "YOU", []string{})
	san := getAncestors(graph, "SAN", []string{})
	fmt.Println(minimumOrbitChange(you, san))
}
