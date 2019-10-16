package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const AsciiSize = 128

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

		freq, _ := getFreqAndLength(inputString)
		var freqString string

		for i,v := range freq {
			freqString += fmt.Sprintf("%d", v)
			if i < len(freq)-1 {
				freqString += "\n"
			}
		}

		err = ioutil.WriteFile("freq.txt", []byte(freqString), 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lBorder, rBorder := encode(inputString)

		err = ioutil.WriteFile("lBorder.txt", []byte(lBorder), 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = ioutil.WriteFile("rBorder.txt", []byte(rBorder), 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}
}
