package lab1

import (
	"flag"
	"fmt"
	"os"
)

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
		return
	}

	fmt.Println(files)
	fmt.Println(len(files))
	fmt.Println(*modePtr)
}
