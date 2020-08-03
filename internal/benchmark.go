package internal

import (
	"bufio"
	"strings"
	"testing"
)

// Benchmark benchmarks a solver function.
func Benchmark(f func(*bufio.Scanner) (string, string, error), b *testing.B, input string) {
	var ans1, ans2 string
	var err error
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		input := bufio.NewScanner(strings.NewReader(input))
		b.StartTimer()
		ans1, ans2, err = f(input)
	}
	_, _, _ = ans1, ans2, err
}
