package memo

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/mitchellh/go-homedir"
	"github.com/russross/blackfriday"
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
	if err != nil{
		return err
	}
	for _, f := range files {
		fmt.Println(f)
	}
	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Publish All title
	home, _ := homedir.Expand("~/")
	fmt.Fprintln(w, "<h1> メモ一覧 </h1>")
	fmt.Fprintln(w, "<hr>")
	fmt.Fprintln(w, "<ul>")
	files, err := list()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, f := range files {
		fp, err := os.Open(home + "/memo/" + f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		scanner := bufio.NewScanner(fp)
		scanner.Scan()
		title := scanner.Text()[7:]
		fmt.Fprintln(w, "<li>")
		fmt.Fprintln(w, `<a href="./list/`+f+`">`+title+`</a>`)
		fmt.Fprintln(w, "</li>")
		fp.Close()
	}
	fmt.Fprintln(w, "</ul>")

}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	// Publish each web-page
	filename := r.URL.Path[6:] // URI is (http://localhost:XXXX/list/<filename>)
	home, _ := homedir.Expand("~/")
	md, err := ioutil.ReadFile(home + "/memo/" + filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	output := blackfriday.MarkdownCommon([]byte(md))
	fmt.Fprintln(w, string(output))
	fmt.Fprintln(w, "<hr>")
	fmt.Fprintln(w, `<a href="../../">戻る</a>`)

}

// Serve : service of html
func Serve() {
	fmt.Println("See http://localhost:9595/")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/list/", detailsHandler)
	http.ListenAndServe(":9595", nil)
}

// Help : show help
func Help() {
	fmt.Println("usage: go-memo [sub-command]")
	fmt.Println("sub-commands:")
	fmt.Println("  serve     You can see notes at http://localhost:9595/")
	fmt.Println("  list      You can see title of all notes")

}
