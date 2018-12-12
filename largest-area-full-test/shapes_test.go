package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectGoodArea(t *testing.T) {
	r := &Rect{
		Width:  1,
		Height: 2,
	}
	area, err := r.Area()
	assert.NoError(t, err)
	assert.Equal(t, float64(2), area)
}

func TestRectError(t *testing.T) {
	rs := map[string]Rect{
		"width <= 0": {
			Width:  0,
			Height: 2,
		},
		"width is NaN": {
			Width:  math.NaN(),
			Height: 2,
		},
		"width is +Inf": {
			Width:  math.Inf(1),
			Height: 2,
		},
		"height <= 0": {
			Width:  1,
			Height: -1,
		},
		"height is NaN": {
			Width:  1,
			Height: math.NaN(),
		},
		"height is +Inf": {
			Width:  1,
			Height: math.Inf(1),
		},
	}
	for name, r := range rs {
		_, err := r.Area()
		assert.Error(t, err, "error case: %s", name)
	}
}

func TestCircleGoodArea(t *testing.T) {
	r := &Circle{
		Radius: 1,
	}
	area, err := r.Area()
	assert.NoError(t, err)
	assert.Equal(t, float64(3.14), area)
}

func TestCircleError(t *testing.T) {
	rs := map[string]Circle{
		"radius <= 0": {
			Radius: 0,
		},
		"radius is NaN": {
			Radius: math.NaN(),
		},
		"radius is +Inf": {
			Radius: math.Inf(1),
		},
	}
	for name, r := range rs {
		_, err := r.Area()
		assert.Error(t, err, "error case: %s", name)
	}
}
