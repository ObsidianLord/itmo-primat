package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const AsciiSize = 128

var items = make(map[rune]int)
var fileLength = 0
var rating = make([]rune, len(items))

func main() {

	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("Ошибка! Необходимо указать файлы, подлежащие обработке")
		os.Exit(1)
	}

	for _, f := range files {

		buf, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		inputString := string(buf)

		for _, r := range inputString {
			if r < AsciiSize {
				if r >= 32 {
					items[r]++
					fileLength++
				}
			} else {
				fmt.Printf("Ошибка! Файл %s содержит недопустимые символы\n", f)
				os.Exit(1)
			}
		}

		rating = make([]rune, len(items))
		fmt.Println(f + "\n")
		fmt.Println("Символ\tВероятность\tКодовое слово\tДлина кодового слова")
		huffmanTable := huffman(items)
		fmt.Println()
		fmt.Println("Символ\tВероятность\tКодовое слово\tДлина кодового слова")

		shannonFanoTable := shannonFano(items)
		var huffmanText, shannonFanoText string

		for _, v := range inputString {
			huffmanText += huffmanTable[v]
			shannonFanoText += shannonFanoTable[v]
		}

		fmt.Println("\nЭнтропия\nИсходный файл\tкод Хаффмана\tкод Шеннона-Фано")
		fmt.Printf("%.4f\t%.4f\t%.4f\n", entropy(inputString), entropy(huffmanText), entropy(shannonFanoText))

		items = make(map[rune]int)
		fileLength = 0
		rating = make([]rune, len(items))
	}
}
