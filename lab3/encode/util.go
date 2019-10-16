package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

func getTempBigFloat() *big.Float {
	temp, _, _ := big.ParseFloat("0.0", 10, 250, big.ToZero)
	return temp
}

func getFreqAndLength(s string) (freq []int, length int) {

	freq = make([]int, AsciiSize)
	length = 0

	for _, r := range s {
		if r < AsciiSize {
			freq[r]++
			length++
		} else {
			fmt.Println("Ошибка! Строка содержит недопустимые символы")
			os.Exit(1)
		}
	}

	return
}

func getBorders(freq []int, length int) (lBorders []*big.Float, rBorders []*big.Float) {

	ptr, _, _ := big.ParseFloat("0.0", 10, 250, big.ToZero)
	temp, _, _ := big.ParseFloat("0.0", 10, 250, big.ToZero)
	var fLength float64 = float64(length)
	lBorders, rBorders = make([]*big.Float, AsciiSize), make([]*big.Float, AsciiSize)

	for r, f := range freq {

		if f <= 0 {
			lBorders[r], _, _ = big.ParseFloat("-1.0", 10, 250, big.ToZero)
			rBorders[r], _, _ = big.ParseFloat("-1.0", 10, 250, big.ToZero)
			continue
		}

		rangeSize := temp.Quo(big.NewFloat(float64(f)), big.NewFloat(fLength))
		//fmt.Println(rangeSize.Text('f', 250))
		lBorders[r], _, _ = big.ParseFloat("0.0", 10, 250, big.ToZero)
		rBorders[r], _, _ = big.ParseFloat("0.0", 10, 250, big.ToZero)
		lBorders[r].Copy(ptr)
		rBorders[r].Copy(ptr.Add(ptr, rangeSize))

		//fmt.Println(r, rBorders[r].Text('f',30))

	}

	return
}

func encode(s string) (lBorderString string, rBorderString string) {

	freq, fileLength := getFreqAndLength(s)
	lBorders, rBorders := getBorders(freq, fileLength)
	prevLBorder, _, _ := big.ParseFloat("0.0", 10, 250, big.ToZero)
	prevRBorder, _, _ := big.ParseFloat("1.0", 10, 250, big.ToZero)
	lBorder, _, _ := big.ParseFloat("0.0", 10, 250, big.ToZero)
	rBorder, _, _ := big.ParseFloat("1.0", 10, 250, big.ToZero)

	lBorderString, rBorderString = "0.", "0."

	for _, r := range s {
		if lBorders[r].Cmp(big.NewFloat(0.0)) < 0 {
			continue
		}
		//fmt.Println(prevLBorder.Text('f',15)[:17], prevRBorder.Text('f',15)[:17])

		lBorder = getTempBigFloat().Add(prevLBorder, getTempBigFloat().Mul(getTempBigFloat().Sub(prevRBorder, prevLBorder), lBorders[r]))
		rBorder = getTempBigFloat().Add(prevLBorder, getTempBigFloat().Mul(getTempBigFloat().Sub(prevRBorder, prevLBorder), rBorders[r]))

		//fmt.Println(lBorder.Text('f',30), "\n", rBorder.Text('f',30))
		var lPart, rPart string = "", ""
		offset := 1

		//fmt.Println(lBorder.Text('f',15)[:17], rBorder.Text('f',15)[:17])
		//fmt.Println(lBorder.Text('f',250)[2:3], rBorder.Text('f',250)[2:3])

		for {
			if lBorder.Text('f', 250)[2:offset+2] == rBorder.Text('f', 250)[2:offset+2] {
				offset++
			} else {
				break
			}
		}

		if offset > 1 {
			lPart = lBorder.Text('f', 250)[2 : offset+1]
			rPart = rBorder.Text('f', 250)[2 : offset+1]
			prevLBorder, _, _ = big.ParseFloat("0." + lBorder.Text('f', 250)[offset+1:], 10, 250, big.ToZero)
			prevRBorder, _, _ = big.ParseFloat("0." + rBorder.Text('f', 250)[offset+1:], 10, 250, big.ToZero)
		} else {
			prevLBorder.Copy(lBorder)
			prevRBorder.Copy(rBorder)
		}

		if lPart != "" {
			lBorderString += lPart
		}

		if rPart != "" {
			rBorderString += rPart
		}

		//fmt.Println(lBorderString, rBorderString)
		//if i>10 {break}
	}

	prevLBorderString := strings.Split(prevLBorder.Text('f',15),".")[0]
	prevRBorderString := strings.Split(prevRBorder.Text('f',15),".")[0]
	lBorderString += prevLBorderString
	rBorderString += prevRBorderString

	return
}
