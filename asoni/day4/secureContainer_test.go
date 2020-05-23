package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_verifyDigits(t *testing.T) {
	testCases := []struct {
		num            int
		expectedAnswer bool
	}{
		{
			num:            123456,
			expectedAnswer: true,
		},
		{
			num:            12345,
			expectedAnswer: false,
		},
		{
			num:            1234,
			expectedAnswer: false,
		},
		{
			num:            123,
			expectedAnswer: false,
		},
		{
			num:            12,
			expectedAnswer: false,
		},
		{
			num:            1,
			expectedAnswer: false,
		},
		{
			num:            1234567,
			expectedAnswer: false,
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			ans := verifyDigits(tc.num)
			assert.Equal(t, tc.expectedAnswer, ans)
		})
	}
}

func Benchmark_verifyDigits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		verifyDigits(278384)
	}
}

func Test_getDigits(t *testing.T) {
	testCases := []struct {
		num            int
		expectedAnswer [6]int
	}{
		{
			num:            123456,
			expectedAnswer: [6]int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			ans := getDigits(tc.num)
			assert.Equal(t, tc.expectedAnswer, ans)
		})
	}
}

func Benchmark_getDigits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getDigits(278384)
	}
}
func Test_adjacentDigits(t *testing.T) {
	testCases := []struct {
		digits         [6]int
		expectedAnswer bool
	}{
		{
			digits:         [6]int{1, 2, 3, 4, 5, 5},
			expectedAnswer: true,
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			ans := adjacentDigits(tc.digits)
			assert.Equal(t, tc.expectedAnswer, ans)
		})
	}

}
