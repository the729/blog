package main

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockShape struct {
	AreaVal float64
	Err     error
}

func (s *mockShape) Area() (float64, error) {
	return s.AreaVal, s.Err
}

func TestLargest(t *testing.T) {
	shapes := []Shape{
		&mockShape{
			AreaVal: 1.0,
		},
		&mockShape{
			AreaVal: 2.0,
		},
		&mockShape{
			AreaVal: 1.5,
		},
	}
	maxArea, err := largest(shapes)
	assert.NoError(t, err)
	assert.InEpsilon(t, 2.0, maxArea, 1e-9)
}

func TestLargestError(t *testing.T) {
	shapes := []Shape{
		&mockShape{
			AreaVal: 1.0,
		},
		&mockShape{
			AreaVal: 2.0,
		},
		&mockShape{
			AreaVal: 1.5,
			Err:     errors.New("some error"),
		},
	}
	_, err := largest(shapes)
	assert.Error(t, err)
}

func TestReadShapes(t *testing.T) {
	inputStr := `
	[
		{
			"shape": "rect",
			"geometry": {
				"width": 1,
				"height": 2
			}
		},
		{
			"shape": "circle",
			"geometry": {
				"radius": 1
			}
		}
	]`

	shapes, err := readShapes(bytes.NewReader([]byte(inputStr)))
	assert.NoError(t, err)
	assert.ElementsMatch(t, []Shape{
		&Rect{
			Width:  1,
			Height: 2,
		},
		&Circle{
			Radius: 1,
		},
	}, shapes)
}

func TestReadShapes_UnsupportedShape(t *testing.T) {
	inputStr := `
	[
		{
			"shape": "unsupported",
			"geometry": {}
		}
	]`

	_, err := readShapes(bytes.NewReader([]byte(inputStr)))
	assert.Error(t, err)
}

func TestLargestAreaFromReader(t *testing.T) {
	inputStr := `
	[
		{
			"shape": "rect",
			"geometry": {
				"width": 1,
				"height": 2
			}
		},
		{
			"shape": "circle",
			"geometry": {
				"radius": 1
			}
		}
	]`

	maxArea, err := LargestAreaFromReader(bytes.NewReader([]byte(inputStr)))
	assert.NoError(t, err)
	assert.InEpsilon(t, 3.14, maxArea, 1e-9)
}
