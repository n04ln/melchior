package memo

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/russross/blackfriday"
	"log"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html>")
	fmt.Fprintln(w, "<head>")
	fmt.Fprintln(w, `<link rel="stylesheet" type="text/css" href="style.css">`)
	fmt.Fprintln(w, "</head>")
	fmt.Fprintln(w, "<body>")
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
	fmt.Fprint(w, "</body>")
	fmt.Fprint(w, "</html>")

}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(os.Getenv("GOPATH")+"/src/github.com/NoahOrberg/go-memo/css/style.css")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintf(w, string(data))
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html>")
	fmt.Fprintln(w, "<head>")
	fmt.Fprintln(w, `<link rel="stylesheet" type="text/css" href="../style.css">`)
	fmt.Fprintln(w, "</head>")
	fmt.Fprintln(w, "<body>")
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
	fmt.Fprint(w, "</body>")
	fmt.Fprint(w, "</html>")

}
