package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-c" && len(os.Args) < 5 {
		printFile(os.Args[3])
	} else if len(os.Args) > 1 && os.Args[1] == "-c" {
		fileInfo, err := os.Stat(os.Args[3])
		if err != nil {
			fmt.Printf("ERROR1")
		}
		fmt.Printf("=> " + fileInfo.Name() + " <==\n")
		printFile(os.Args[3])

		fmt.Printf("\n")
		fileInfo2, err2 := os.Stat(os.Args[4])
		if err2 != nil {
			fmt.Printf("ERROR2")
		}
		fmt.Printf("=> " + fileInfo2.Name() + " <==\n")
		printFile(os.Args[4])
	}
}

func printFile(arg string) {
	//If file exists
	fileExist, err := os.Stat(arg)
	if err != nil && fileExist == nil {
		if os.IsNotExist(err) {
			fmt.Printf("ERROR: open " + os.Args[3] + ": no such file or directory\n")
			os.Exit(1)
		}
	}

	//Open file
	file, err := os.Open(arg)
	if err != nil {
		fmt.Printf(string("ERROR OPENING FILE"))
		os.Exit(0)
	}

	//Read file
	var arr []byte
	data, err := file.Read(arr)
	data2 := string(data)
	if err != nil {
		fmt.Printf(string("ERROR READING FILE"))
		os.Exit(0)
	}
	for i := len(data2) - Atoi2(os.Args[2]); i < len(data2); i++ {
		fmt.Printf(string(string(data)[i]))
	}
	file.Close()
}

func Atoi2(s string) int {
	strTable := []rune(s)
	var table = make([]int, len(strTable))
	var res int
	var isNegative bool

	for i := 0; i < len(strTable); i++ {
		table[i] = int(strTable[i] - 48)
	}
	for j := 0; j < len(strTable); j++ {
		if (table[j] == 0) || (table[j] == 1) || (table[j] == 2) || (table[j] == 3) || (table[j] == 4) || (table[j] == 5) || (table[j] == 6) || (table[j] == 7) || (table[j] == 8) || (table[j] == 9) || (table[0] == 43-48) || (table[0] == 45-48) {
			if table[0] == 45-48 {
				isNegative = true
				table[0] = 0
			}
			if table[0] == 43-48 {
				table[0] = 0
			}
			if len(s) > 2 {
				if (table[1] == 43-48) || (table[1] == 45-48) {
					return 0
				}
			}
			res = res*10 + table[j]
		} else {
			return 0
		}
	}
	if isNegative == true {
		return res * -1
	} else {
		return res
	}
}
