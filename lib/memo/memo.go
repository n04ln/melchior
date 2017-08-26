package memo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"log"

	"github.com/mitchellh/go-homedir"
)

func list() ([]string, error) {
	var res []string
	path, _ := homedir.Expand("~/memo")
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		isDotFile, err := regexp.MatchString(`^\..*$`, file.Name())
		if err != nil {
			return []string{}, nil
		}
		if !isDotFile {
			res = append(res, file.Name())
		}
	}
	return res, nil
}

// ViewList : show list
func ViewList() error {
	files, err := list()
	if err != nil {
		return err
	}
	for _, f := range files {
		fmt.Println(f)
	}
	return nil
}

// Serve : service of html
func Serve() {
	fmt.Println("See http://localhost:9595/")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/list/", detailsHandler)
	http.HandleFunc("/style.css", cssHandler)
	if err := http.ListenAndServe(":9595", nil); err != nil {
		log.Fatal(err)
	}
}

// Help : show help
func Help() {
	fmt.Println("usage: go-memo [sub-command]")
	fmt.Println("sub-commands:")
	fmt.Println("  serve     You can see notes at http://localhost:9595/")
	fmt.Println("  list      You can see title of all notes")

}

