package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	var cycles = 5

	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			cyclesStr := r.FormValue("cycles")
			if cyclesStr != "" {
				var err error
				cycles, err = strconv.Atoi(cyclesStr)
				if err != nil {
					fmt.Fprintf(w, "bad cycles param: %s", cyclesStr, err)
					return
				}
			}
			lissajous(cycles, w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(cycles, os.Stdout)
}

func lissajous(cycles int, out io.Writer) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	palette := make([]color.Color, 0, nframes)
	palette = append(palette, color.RGBA{0, 0, 0, 255})
	for i := 0; i < nframes; i++ {
		scale := float64(i) / float64(nframes)
		c := color.RGBA{uint8(55 + 200*scale), uint8(55 + 200*scale), uint8(55 + 200*scale), 255}
		palette = append(palette, c)
	}
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8((i%(len(palette)-1))+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
