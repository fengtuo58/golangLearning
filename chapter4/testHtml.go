package main

import(
	"time"
	"encoding/json"
	"net/http"
	"net/url"
	"html/template"
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

var issueList = template.Must(template.New("issuelist").Parse(` <h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}} <tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td> <td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td> <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr> {{end}} </table> `))

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
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
