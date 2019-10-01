package main

import (
	"fmt"
	"math"
	"os"
)

func entropy(s string) (result float64) {

	sCount := make([]int, AsciiSize)
	sP := make([]float64, AsciiSize)
	sH := make([]float64, AsciiSize)
	var fileLength int

	for _, r := range s {
		if r < AsciiSize {
			if r >= 32 {
				sCount[r]++
				fileLength++
			}
		} else {
			fmt.Println("Ошибка! Строка содержит недопустимые символы")
			os.Exit(1)
		}
	}

	for i, _ := range sP {
		if sCount[i] > 0 {
			sP[i] = float64(sCount[i]) / float64(fileLength)
			sH[i] = - sP[i] * math.Log2(sP[i])
			result += sH[i]
		}
	}

	return result
}
