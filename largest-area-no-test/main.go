package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	// 读文件内容
	inputFile := flag.String("i", "inputs.json", "Specify input file.")
	flag.Parse()
	inputContent, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Println("input file error.")
		return
	}

	// 解出结构体
	shapes := []struct {
		Shape    string             `json:"shape"`
		Geometry map[string]float64 `json:"geometry"`
	}{}
	if err := json.Unmarshal(inputContent, &shapes); err != nil {
		fmt.Println("Unmarshal error: ", err)
		return
	}

	// 计算最大面积
	maxArea := float64(0)
	for _, s := range shapes {
		var area float64
		switch s.Shape {
		case "rect":
			area = s.Geometry["width"] * s.Geometry["height"]
		case "circle":
			area = 3.14 * s.Geometry["radius"] * s.Geometry["radius"]
		}

		if area > maxArea {
			maxArea = area
		}
	}

	fmt.Println("Shape with max area: ", maxArea)
}
