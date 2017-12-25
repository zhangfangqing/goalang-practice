package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"gopl.io/ch4/github"
)

var issueList = template.Must(template.New("issuelist").Parse(`
	<h1>{{.TotalCount}} issues</h1>
	<table>
	<tr style='text-align: left'>
		<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
	</tr>
	{{range .Items}}
	<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`))

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		makeHtml(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func makeHtml(out io.Writer) {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(out, result); err != nil {
		log.Fatal(err)
	}
}
