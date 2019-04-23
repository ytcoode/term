package main

import (
	"fmt"

	"github.com/wangyuntao/term"
)

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Cleanup()

	for {
		e := term.PollEvent()

		switch v := e.(type) {
		case term.WinResize:
			row, col, err := term.WinSize()
			if err != nil {
				panic(err)
			}
			fmt.Println("winResize:", row, col)

		case term.Key:
			fmt.Println("Key:", v)

		case rune:
			fmt.Printf("Rune: %d\n", v)
		}
	}
}