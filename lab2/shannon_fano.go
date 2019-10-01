package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Code struct {
	Char rune
	Prob float64
	Bits uint32
	Size int
}

func shannonFano(items map[rune]int) map[rune]string {

	var divide func([]Code)
	codes := make([]Code, 0, len(items))

	for r, freq := range items {
		codes = append(codes, Code{
			Char: r,
			Prob: float64(freq) / float64(fileLength),
		})
	}

	sort.Slice(codes, func(a, b int) bool {
		return codes[a].Prob > codes[b].Prob
	})

	divide = func(codes []Code) {
		var p int

		if len(codes) < 2 {
			return
		}

		// sum the total probability for this slice
		freq := 0.0
		for _, code := range codes {
			freq += code.Prob
		}

		// probability of the left half
		left := codes[0].Prob
		best := 1.0

		// find the optimal pivot
		for p = 1; p < len(codes)-1; p++ {
			if diff := math.Abs((freq - left) - left); diff < best {
				best = diff
			} else {
				break
			}

			left += codes[p].Prob
		}

		for i := 0; i < len(codes); i++ {
			codes[i].Bits <<= 1
			codes[i].Size++

			if i >= p {
				codes[i].Bits |= 1
			}
		}

		divide(codes[:p])
		divide(codes[p:])
	}

	divide(codes)

	table := make(map[rune]string)
	for _, code := range codes {
		if len(strconv.FormatInt(int64(code.Bits), 2)) < code.Size {
			sizeDiff := code.Size - len(strconv.FormatInt(int64(code.Bits), 2))
			table[code.Char] = strings.Repeat("0", sizeDiff) + strconv.FormatInt(int64(code.Bits), 2)
		} else {
			table[code.Char] = strconv.FormatInt(int64(code.Bits), 2)
		}

	}

	fmt.Println("Код Шеннона-Фано:")
	for _, v := range rating {
		fmt.Printf("%s\t%.4f\t%s\t%d\n", string(v), float64(items[v])/float64(fileLength), table[v], len(table[v]))
	}

	return table
}
