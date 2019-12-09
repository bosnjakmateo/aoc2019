package main

import "strings"

func DecodeImage(image []string) {
	var decodedImage [150]string

	for i := 0; i < 150; i++ {
		for j := range image {
			layer := strings.Split(image[j], "")
			switch layer[i] {
			case "0":
				decodedImage[i] = "░"
			case "1":
				decodedImage[i] = "█"
			case "2":
				continue
			}
			break
		}
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 25; j++ {
			print(decodedImage[j+(i*25)])
		}
		println("")
	}
}
