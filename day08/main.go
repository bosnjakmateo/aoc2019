package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileString("day08/day08.txt", "-1")
	value := values[0]
	var image []string

	for i := 0; i < len(values[0]); i += 150 {
		batch := value[i:utils.Min(i+150, len(values[0]))]
		image = append(image, batch)
	}

	fmt.Printf("Solution part one: %d\n", VerifyImage(image))
	fmt.Println("Solution part two:")
	DecodeImage(image)
}
