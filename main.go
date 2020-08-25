package main

import (
	"bufio"
	"fmt"
	"github.com/arcgemu/celscript/celscript"
	"os"
)

func main()  {
	lic := make(chan string)
	lexer := celscript.NewLexer(lic)

	go lexer.Run()

	for {
		r := bufio.Reader(os.Stdin)
		fmt.print(">")
		t, err := r.ReadString('\n')

		if err != nil {
			panic(err)
		}
		lexer.In <- t
	}
}
