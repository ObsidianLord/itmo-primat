package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

const AsciiSize = 128
const ACode = 65
const ZCode = 90
const CaseOffset = 32
const SpaceCode = 32
const CommaCode = 44

func sliceSum(a []int) (sum int) {
	for _, v := range a {
		sum += v
	}
	return
}

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Ошибка! Необходимо указать файлы, подлежащие обработке")
		os.Exit(1)
	}

	for fi, f := range files {
		sCount := make([]int, AsciiSize)
		sP := make([]float64, AsciiSize)
		sH := make([]float64, AsciiSize)
		spCount := make([][]int, AsciiSize)
		spP := make([][]float64, AsciiSize)
		for i := 0; i < AsciiSize; i++ {
			spCount[i], spP[i] = make([]int, AsciiSize), make([]float64, AsciiSize)
		}
		var fileLength int
		var entropy, pairEntropy float64
		var prev rune
		buf, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		inputString := string(buf)

		for ri, r := range inputString {
			if r < AsciiSize {
				switch {
				case r >= aCode && r <= zCode:
					sCount[r-CaseOffset]++
					if ri > 0 {
						spCount[prev][r-CaseOffset]++
					} else {
						spCount[CommaCode][r-CaseOffset]++
					}
					prev = r - CaseOffset
				case r >= ACode && r <= ZCode:
					sCount[r]++
					if ri > 0 {
						spCount[prev][r]++
					} else {
						spCount[CommaCode][r]++
					}
					prev = r
				case r == SpaceCode:
					sCount[SpaceCode]++
					if ri > 0 {
						spCount[prev][SpaceCode]++
					} else {
						spCount[CommaCode][SpaceCode]++
					}
					prev = SpaceCode
				default:
					sCount[CommaCode]++
					if ri > 0 {
						spCount[prev][CommaCode]++
					} else {
						spCount[CommaCode][CommaCode]++
					}
					prev = CommaCode
				}
				fileLength++
			} else {
				fmt.Printf("Ошибка! Файл %s содержит недопустимые символы\n", f)
				os.Exit(1)
			}
		}

		for i, _ := range sP {
			if sCount[i] > 0 {
				sP[i] = float64(sCount[i]) / float64(fileLength)
				sH[i] = - sP[i] * math.Log2(sP[i])
				entropy += sH[i]
			}
		}

		if !shortMode {
			for j, _ := range spP {
				for i, _ := range spP[j] {
					if spCount[j][i] > 0 {
						spP[j][i] = float64(spCount[j][i]) / float64(sliceSum(spCount[j]))
						pairEntropy += - spP[j][i] * sP[j] * math.Log2(spP[j][i])
					}
				}
			}
		}

		fmt.Printf("Файл: %s, Суммарная энтропия: %.4f бит\n", f, entropy)
		if !shortMode {
			fmt.Printf("Суммарная энтропия с учетом частот вхождений пар символов:  %.4f бит\n", pairEntropy)
		}
		for i, v := range sP {
			if v > 0.0 {
				fmt.Printf("\"%s\": H=%.4f P=%.4f\n", string(rune(i)), sH[i], v)
			}
		}

		if fi+1 != len(files) {
			fmt.Println()
		}
	}
}
