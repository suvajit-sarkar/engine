package utilities

import (
	"strconv"
	"strings"
)

//IntegerToSting returns string converted from int
func IntegerToSting(x int) string {
	return strings.Join([]string{"", strconv.Itoa(x)}, "")
}

//Find if a slice has the value
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// MaxFloat return max of two floats
func MaxFloat(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

// MaxInt return max of two floats
func MaxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}
