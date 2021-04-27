package main
import "time"
type Track struct {
	Title string
	Artist string
	Alblum string
	Year string
	Year string
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Deliab", "From", 2012, length("3m38s")},
	{"Go", "dd", "dd", 453, length("ddd")},
}


