package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCsvFile(t *testing.T) {
	output := []string{"1", "2", "3"}
	assert.Equal(t, output, readCsvFile("test.csv"))
}

func BenchmarkReadCsvFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readCsvFile("test.csv")
	}
}

func TestParseData(t *testing.T) {

}

func BenchmarkParseData(b *testing.B) {

}

func TestRunOpCode(t *testing.T) {

}

func BenchmarkRunOpCode(b *testing.B) {

}

func TestPart2(t *testing.T) {

}

func BenchmarkPart2(b *testing.B) {

}
