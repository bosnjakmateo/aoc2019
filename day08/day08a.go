package main

import "strings"

func VerifyImage(image []string) int {
	layer := image[0]
	numberOfZeros := strings.Count(image[0], "0")

	for i := range image {
		if strings.Count(image[i], "0") < numberOfZeros {
			numberOfZeros = strings.Count(image[i], "0")
			layer = image[i]
		}
	}

	return strings.Count(layer, "1") * strings.Count(layer, "2")
}
