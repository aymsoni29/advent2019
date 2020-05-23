package day1

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInputFile(t *testing.T) {
	expected := []int{1, 2, 3}
	actual := parseInputFile("day-1-input.txt")

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Output not as expected")
	}
}

/*
	Running tool: /usr/local/bin/go test -benchmem -run=^$ github.com/advent2019/asoni/day1 -bench ^(BenchmarkParseInputFile)$

	goos: darwin
	goarch: amd64
	pkg: github.com/advent2019/asoni/day1
	BenchmarkParseInputFile-8   	   38934	     30440 ns/op	    4272 B/op	       7 allocs/op
	PASS
	ok  	github.com/advent2019/asoni/day1	11.675s
*/
func BenchmarkParseInputFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseInputFile("testFile.txt")
	}
}

func TestCalculateFuel(t *testing.T) {
	testCases := []int{-1, 0, 1, 3, 6}
	expectedResults := []int{-2, -2, -2, -1, 0}

	for i, val := range testCases {
		assert.Equal(t, expectedResults[i], calculateFuel(val), "Expected output different from actual output")
	}
}

/*
	goos: darwin
	goarch: amd64
	pkg: github.com/advent2019/asoni/day1
	BenchmarkCalculateFuel-8   	1000000000	         0.570 ns/op	       0 B/op	       0 allocs/op
	PASS
	ok  	github.com/advent2019/asoni/day1	3.411s
*/
func BenchmarkCalculateFuel(b *testing.B) {
	testCase := 6
	for i := 0; i < b.N; i++ {
		calculateFuel(testCase)
	}
}
