package main

import (
	"math/big"
)

func getTempBigFloat() *big.Float {
	temp, _, _ := big.ParseFloat("0.0", 10, 250, big.ToZero)
	return temp
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

		//fmt.Println(r, lBorders[r].Text('f',30), rBorders[r].Text('f',30))

	}

	return
}
