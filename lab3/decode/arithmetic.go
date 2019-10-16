package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/big"

	"os"
	"strconv"
)

const AsciiSize = 128

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Ошибка! Необходимо указать файлы, подлежащие обработке")
		os.Exit(1)
	}

	var length int

	buf, err := ioutil.ReadFile(files[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lBorder := string(buf)

	buf, err = ioutil.ReadFile(files[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rBorder := string(buf)

	var freq []int = make([]int, AsciiSize)

	file, err := os.Open(files[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctr := 0
	reader := bufio.NewReader(file)
	var count []byte
	for {
		count, _, err = reader.ReadLine()

		if err != nil {
			break
		}

		freq[ctr], err = strconv.Atoi(string(count))
		length += freq[ctr]
		ctr++
	}
	_ = file.Close()

	lBorders, rBorders := getBorders(freq, length)

	cutLBorder, _, _ := big.ParseFloat(lBorder, 10, 250, big.ToZero)
	cutRBorder, _, _ := big.ParseFloat(rBorder, 10, 250, big.ToZero)

	var decodedString string

	for i, v := range lBorders {
		if cutLBorder.Cmp(v) >= 0 && cutRBorder.Cmp(rBorders[i]) <= 0 {
			decodedString = string(rune(i))
		}
	}

	lFloatBorder, _, _ := big.ParseFloat(lBorder, 10, uint(len(lBorder)-2), big.ToZero)
	rFloatBorder, _, _ := big.ParseFloat(rBorder, 10, uint(len(rBorder)-2), big.ToZero)
	temp, _, _ := big.ParseFloat(lBorder, 10, uint(len(lBorder)-2), big.ToZero)
	temp.Add(lFloatBorder, rFloatBorder)
	temp.Quo(temp, big.NewFloat(2))
	code, _, _ := big.ParseFloat(rBorder, 10, uint(len(rBorder)-2), big.ToZero)
	code.Copy(temp)

	for j := 1; j < length; j++ {
		temp1, _, _ := big.ParseFloat(lBorder, 10, uint(len(lBorder)-2), big.ToZero)
		temp2, _, _ := big.ParseFloat(lBorder, 10, uint(len(lBorder)-2), big.ToZero)
		temp3, _, _ := big.ParseFloat(lBorder, 10, uint(len(lBorder)-2), big.ToZero)

		lSymBorder, rSymBorder := lBorders[decodedString[j-1]], rBorders[decodedString[j-1]]

		//fmt.Println(code.Text('f', 30), lSymBorder.Text('f', 30), rSymBorder.Text('f', 30))
		upper := temp1.Sub(code, lSymBorder)
		lower := temp2.Sub(rSymBorder, lSymBorder)
		code.Copy(temp3.Quo(upper, lower))

		//fmt.Println(code.Text('f', 30))

		for i, v := range lBorders {
			if v.Cmp(getTempBigFloat()) >= 0 {
				if v.Cmp(code) <= 0 && rBorders[i].Cmp(code) >= 0 {
					decodedString += string(rune(i))
					break
				}
			}
		}

		//if j>3 {break}
	}

	err = ioutil.WriteFile("result.txt", []byte(decodedString), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
