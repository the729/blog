package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func largest(shapes []Shape) (float64, error) {
	maxArea := float64(0)
	for _, shape := range shapes {
		area, err := shape.Area()
		if err != nil {
			return 0, fmt.Errorf("get area of shape error: %s", err)
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea, nil
}

func readShapes(r io.Reader) ([]Shape, error) {
	inputs := []struct {
		Shape    string           `json:"shape"`
		Geometry *json.RawMessage `json:"geometry"`
	}{}
	if err := json.NewDecoder(r).Decode(&inputs); err != nil {
		return nil, fmt.Errorf("JSON decode error: %s", err)
	}

	shapes := []Shape{}
	for _, input := range inputs {
		var shape Shape
		switch input.Shape {
		case "rect":
			shape = &Rect{}
		case "circle":
			shape = &Circle{}
		}
		if err := json.Unmarshal(*input.Geometry, shape); err != nil {
			return nil, fmt.Errorf("JSON decode shape error: %s", err)
		}
		shapes = append(shapes, shape)
	}
	return shapes, nil
}

func LargestAreaFromReader(r io.Reader) (float64, error) {
	shapes, err := readShapes(r)
	if err != nil {
		return 0, fmt.Errorf("read shape error: %s", err)
	}

	maxArea, err := largest(shapes)
	if err != nil {
		return 0, fmt.Errorf("get largest shape error: %s", err)
	}

	return maxArea, nil
}

func main() {
	inputFile := flag.String("i", "inputs.json", "Specify input file.")
	flag.Parse()
	inputReader, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("input file error: ", err)
		return
	}
	defer inputReader.Close()

	maxArea, err := LargestAreaFromReader(inputReader)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Shape with max area: ", maxArea)
}
