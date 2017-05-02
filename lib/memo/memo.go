package memo

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/russross/blackfriday"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func list() []string {
	path, _ := homedir.Expand("~/memo")
	files, _ := ioutil.ReadDir(path)
	res := make([]string, len(files))
	for i := 0; i < len(files); i++ {
		res[i] = files[i].Name()
	}
	return res
}

// ViewList : show list
func ViewList() {
	files := list()
	for _, f := range files {
		fmt.Println(f)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Publish All title
	home, _ := homedir.Expand("~/")
	fmt.Fprintln(w, "<h1> メモ一覧 </h1>")
	fmt.Fprintln(w, "<hr>")
	fmt.Fprintln(w, "<ul>")
	files := list()
	for _, f := range files {
		fp, err := os.Open(home + "/memo/" + f)
		check(err)
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
	check(err)
	output := blackfriday.MarkdownCommon([]byte(md))
	fmt.Fprintln(w, string(output))
	fmt.Fprintln(w, "<hr>")
	fmt.Fprintln(w, `<a href="../../">戻る</a>`)

}

// Serve : service of html
func Serve() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/list/", detailsHandler)
	http.ListenAndServe(":9595", nil)
	fmt.Println("See http://localhost:9595/")
}

// Help : show help
func Help() {
	fmt.Println("usage: go-memo [sub-command]")
	fmt.Println("sub-commands:")
	fmt.Println("  serve     You can see notes at http://localhost:9595/")
	fmt.Println("  list      You can see title of all notes")

}
