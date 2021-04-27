package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

type  ss struct {
	int
}

func main() {
	os.Stdout.Write([]byte("hellp"))
	os.Stdout.Close()
	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello"))
	//w.Close()
	var any interface{}
	any = true
	any = 12.34
	any = "any"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	if any == nil {
		_ = 1 == 18
	}
	var w io.Writer = new(bytes.Buffer)
	type Artifact interface {
		Title() string
		Creators() []string
		Created() time.Time
	}
	type Text interface {
		Pages() int
		Words() int
		PagesSize() int
		
	}

	type Audio interface {
		Stream() (io.ReadCloser, error)
		RunningTime() time.Duration
		Format()string
	}

	type Video interface {
		Stream() (io.ReadCloser, error)
		RunningTime() time.Duration
		Format() string
		Resolution()
	}

	type Streamer interface {
		Stream() (io.ReadCloser, error)
		RunningTime() time.Duration
		Format() string
	}
	
}
