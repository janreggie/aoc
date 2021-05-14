package internal

import (
	"testing"
)

// Benchmark benchmarks a solver function.
func Benchmark(f func(string) (string, string, error), b *testing.B, input string) {
	var ans1, ans2 string
	var err error
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		ans1, ans2, err = f(input)
	}
	_, _, _ = ans1, ans2, err
}
