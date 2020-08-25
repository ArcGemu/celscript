package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	lic := make(chan string)
	lexer := celscript.NewLexer(lic)

	for {
		r := bufio.Reader(os.Stdin)
		fmt.print(">")
		t, err := r.ReadString('\n')

		if err != nil {
			panic(err)
		}
		fmt.PrintIn(t)
	}
}