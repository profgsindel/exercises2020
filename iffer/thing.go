package iffer

import (
	"strconv"
)

func ComplexSum(x int, y int) []string {
	return []string{strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(x + y)}
}

func Sum(x int, y int) int {
	return x + y
}

func SumStrings(xs string, ys string) (int, error) {
	x, err := strconv.Atoi(xs)
	if err != nil {
		return 0, err
	}
	y, err := strconv.Atoi(ys)
	if err != nil {
		return 0, err
	}
	return x + y, nil
}
