package main

import (
	"github.com/NoahOrberg/go-memo/lib/memo"
	"os"
)


func main() {
	if len(os.Args) ==2 {
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
	}else {
		memo.Help()
	}
}
