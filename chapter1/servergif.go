package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"log"
	"net/http"
	"sync"

)

var mu sync.Mutex
var count int


type IpType int

const (
	IPTYPE_NONE IpType = iota
	IPTYPE_V4
	IPTYPE_V6
)


var palette = []color.Color{color.White, color.RGBA{0x0, 0xff, 0x0, 1}}
const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func lissajous(out io.Writer) {
	const(
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount:0}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+ int(x*size+0.5), size+int(y*size+0.5), blackIndex)
			
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, 30)
		anim.Image = append(anim.Image, img)
		
	}
	gif.EncodeAll(out, &anim)


}
