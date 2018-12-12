package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() (float64, error)
}

type Circle struct {
	Radius float64 `json:"radius"`
}

type Rect struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func (s *Circle) Area() (float64, error) {
	if !(s.Radius > 0.0) || math.IsInf(s.Radius, 1) {
		return 0, fmt.Errorf("circle radius should be positive, got %f", s.Radius)
	}
	return 3.14 * s.Radius * s.Radius, nil
}

func (s *Rect) Area() (float64, error) {
	//	if s.Height <= 0.0 {
	if !(s.Height > 0.0) || math.IsInf(s.Height, 1) {
		return 0, fmt.Errorf("rect height should be positive, got %f", s.Height)
	}
	if !(s.Width > 0.0) || math.IsInf(s.Width, 1) {
		return 0, fmt.Errorf("rect width should be positive, got %f", s.Width)
	}
	return s.Height * s.Width, nil
}
