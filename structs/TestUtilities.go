package structs

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/stretchr/testify/assert"
)

// TestCase represents a test case
type TestCase struct {
	Details string // what is the testcase testing? 
	Input   string
	Result1 string
	Result2 string
	WantErr bool // do we want an error to return?
}

// Test tests a test case
func (tc TestCase) Test(f func(*bufio.Scanner) (string, string, error), a *assert.Assertions) {
	scanner := bufio.NewScanner(strings.NewReader(tc.Input))
	answer1, answer2, err := f(scanner)
	details := tc.Details
	if tc.Details == "" {
		details = fmt.Sprintf("input: %v, ans1: %v, ans2: %v", tc.Input, tc.Result1, tc.Result2)
	}
	if tc.WantErr {
		a.Error(err, details)
		return
	}
	a.NoError(err)
	if tc.Result1 != "" {
		a.Equal(tc.Result1, answer1, details)
	}
	if tc.Result2 != "" {
		a.Equal(tc.Result2, answer2, details)
	}
}