package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arg := os.Args
	if len(arg) == 1 {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			z01.PrintRune('A')
		}
		data2 := string(data)
		printStr(data2)
		os.Exit(0)
	} else {
		for i := 1; i < len(arg); i++ {
			file, err := os.Open(arg[i])
			if err != nil {
				printStr("ERROR: open ")
				printStr(arg[i])
				printStr(": no such file or directory")
				z01.PrintRune('\n')
				os.Exit(1)
			} else {
				data, err := ioutil.ReadAll(file)
				if err != nil {
					z01.PrintRune('A')
				}
				printStr(string(data))
			}
		}
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
