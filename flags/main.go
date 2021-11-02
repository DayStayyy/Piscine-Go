package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 || arguments[1] == "-h" || arguments[1] == "--help" {
		printHelp()
	} else if len(arguments) >= 1 && arguments[1] != "-h" && arguments[1] != "--help" {
		arguments = delChar(arguments, 0)
		arg0 := arguments[0]
		if arguments[0] == "--order" || arguments[0] == "-o" {
			order(arguments[1])
		} else if arg0[0:3] == "-i=" {
			if len(arguments) == 2 {
				str := insert(arguments[1], arg0[3:])
				fmt.Print(str)
				z01.PrintRune('\n')
			} else if len(arguments) > 2 {
				str := insert(arguments[2], arg0[3:])
				order(str)
			}
		} else if arg0[0:9] == "--insert=" {
			if len(arguments) == 2 {
				str := insert(arguments[1], arg0[9:])
				fmt.Print(str)
				z01.PrintRune('\n')
			} else if len(arguments) > 2 {
				str := insert(arguments[2], arg0[9:])
				order(str)
			}
		} else {
			table := []rune(arguments[0])
			for i := 0; i < len(table); i++ {
				z01.PrintRune(rune(table[i]))
			}
			z01.PrintRune('\n')
		}
	}

}

func insert(s string, add string) string {
	finalS := s + add
	return finalS
}

func order(s string) {
	table := []rune(s)
	var intTable []int

	for i := 0; i < len(table); i++ {
		intTable = append(intTable, int(table[i]))
	}
	intTable = SortIntegerTable(intTable)
	for j := 0; j < len(intTable); j++ {
		z01.PrintRune(rune(intTable[j]))
	}
	z01.PrintRune('\n')
}

func SortIntegerTable(table []int) []int {
	for i := 0; i < len(table); i++ {
		for j := 0; j < i+1; j++ {
			if table[i] < table[j] {
				temp := table[i]
				table[i] = table[j]
				table[j] = temp
			}
		}
	}
	return table
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
}

func delChar(s []string, index int) []string {
	var result []string
	for i := 0; i < len(s); i++ {
		if i > index {
			result = append(result, s[i])
		}
	}
	return result
}
