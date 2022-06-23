package garray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduceInt(t *testing.T) {
	numbers := []int32{1, 2, 3, 4, 5}

	value := Reduce(numbers, func(value int32, element int32) int32 {
		return value + element
	}, 0)

	assert.Equal(t, 15, int(value))

	value = Reduce(numbers, func(value int32, element int32) int32 {
		return value - element
	}, 15)

	assert.Equal(t, 0, int(value))

	value = Reduce(numbers, func(value int32, element int32) int32 {
		return value + element*2
	}, 0)

	assert.Equal(t, 30, int(value))

	value = Reduce(numbers, func(value int32, element int32) int32 {
		return value + element/2
	}, 0)

	assert.Equal(t, 6, int(value))

}

func TestReduceFloat(t *testing.T) {
	numbers := []float32{1.5, 2.5, 3.5, 4.5, 5.5}

	value := Reduce(numbers, func(value float32, element float32) float32 {
		return value + element
	}, 0)

	assert.Equal(t, float32(17.5), value)

	value = Reduce(numbers, func(value float32, element float32) float32 {
		return value - element
	}, 17.5)

	assert.Equal(t, float32(0.0), value)

	value = Reduce(numbers, func(value float32, element float32) float32 {
		return value + element*2
	}, 0)

	assert.Equal(t, float32(35.0), value)

	value = Reduce(numbers, func(value float32, element float32) float32 {
		return value + element/2
	}, 0)

	assert.Equal(t, float32(8.75), value)
}
