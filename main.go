package main

import (
	"os"

	"github.com/NoahOrberg/melchior/lib/memo"
)

func main() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "list":
			memo.ViewList()
		case "serve":
			memo.Serve()
		case "help":
			memo.Help()
		default:
			memo.Help()
		}
	} else {
		memo.Help()
	}
}
