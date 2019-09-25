package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

const AsciiSize = 128
const aCode = 97
const zCode = 122
const ACode = 65
const ZCode = 90
const CaseOffset = 32
const SpaceCode = 32
const CommaCode = 44

func main() {
	files := os.Args[1:]
	modePtr := flag.Bool("s", false, "short")
	flag.Parse()
	shortMode := *modePtr

	if shortMode {
		files = files[1:]
	}

	if len(files) == 0 {
		fmt.Println("Ошибка! Необходимо указать файлы, подлежащие обработке")
		os.Exit(1)
	}

	for fi, f := range files {
		sCount := make([]int, AsciiSize)
		sP := make([]float64, AsciiSize)
		sE := make([]float64, AsciiSize)
		var fileLength int
		var entropy float64
		buf, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		inputString := string(buf)

		for _, r := range inputString {
			if r < AsciiSize {
				switch {
				case r >= aCode && r <= zCode:
					sCount[r-CaseOffset]++
				case r >= ACode && r <= ZCode:
					sCount[r]++
				case r == SpaceCode:
					sCount[SpaceCode]++
				default:
					sCount[CommaCode]++
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
				sE[i] = - sP[i] * math.Log2(sP[i])
				entropy += sE[i]
			}
		}

		fmt.Printf("Файл: %s, Суммарная энтропия: %.4f бит\n", f, entropy)
		for i, v := range sP {
			if v > 0.0 {
				fmt.Printf("\"%s\": H=%.4f P=%.4f\n", string(rune(i)), sE[i], v)
			}
		}

		if fi+1 != len(files) {
			fmt.Println()
		}
	}
}
