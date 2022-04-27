package piscine

import (
	"ft"
	"os"
)

func printLine(line string) {
	for _,c := range line {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func recursivePrintArray(args []string) {
	args_len := 0
	for i, _ := range args {
		args_len = i
	}
	for ; args_len >= 0; args_len-- {
		printLine(args[args_len])
	}
}

func revParams() {
	recursivePrintArray(os.Args)
    //for i, v := range os.Args {
	//	if i >= 1 {
	//		printLine(v)
	//	}
    //}
}