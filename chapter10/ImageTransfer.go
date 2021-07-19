package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" //register PNG decoder
	"io"
	"os"
	_   "io/ioutil"
	_	"bufio"
	_   "bytes"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg:%v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality:95})
}

//cat 1.png      | ./ImageTransfer >2.jpg
