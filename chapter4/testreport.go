package main

import(
	"time"
	"encoding/json"
	"net/http"
	"net/url"
	"text/template"
	"strings"
	"fmt"
	"os"
	"log"
)

const IssueURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number  int
	HTMLURL string `json:"html_url"`
	Title   string
	State   string
	User    *User
	CreatedAt time.Time `json:"created_at"`
	Body string // in Markdown format
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}

const temp1 = `{{.TotalCount}} Issues:
{{range .Items}} ---------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt| daysAgo}} days
{{end}}
`
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func SearchIssues(term []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(term, " "))
	resp, err := http.Get(IssueURL+ "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.StatusCode)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil 
}

var report = template.Must(template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(temp1))

func main() {

	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Printf("%d issue:\n", result.TotalCount)
	//for _, item := range result.Items {
	//	fmt.Printf("#%-5d %9.9s %55s\n",
	//	item.Number, item.User.Login, item.Title)
	//}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
