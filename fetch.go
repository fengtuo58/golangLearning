package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if (err != nil ) {
			fmt.Fprintf(os.Stderr, "fetch err : %v \n", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : %v \n", err)
			os.Exit(1)
		}

		//		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		//fmt.Printf("%s", b)
	}
	
}
