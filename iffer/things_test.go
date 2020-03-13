package iffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComplexSum(t *testing.T) {
	total := ComplexSum(5, 5)
	assert.ElementsMatch(t, []string{"5", "5", "10"}, total)
}

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	assert.Equal(t, 10, total)
}

func TestSumStrings(t *testing.T) {
	sum, err := SumStrings("1", "2")
	assert.NoError(t, err)
	assert.Equal(t, 3, sum)

	sum, err = SumStrings("0", "3")
	assert.NoError(t, err)
	assert.Equal(t, 3, sum)

	sum, err = SumStrings("-1", "5")
	assert.NoError(t, err)
	assert.Equal(t, 4, sum)

	sum, err = SumStrings("O", "3") // This is an O not a 0
	assert.Error(t, err)

	sum, err = SumStrings("3", "O") // This is an O not a 0
	assert.Error(t, err)
}
