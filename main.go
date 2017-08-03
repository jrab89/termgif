package main

import (
	"bytes"
	"image/color"
	"image/gif"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

// the same thing as https://processing.org/reference/map_.html
func lerp(value, istart, istop, ostart, ostop float32) float32 {
	return ostart + (ostop-ostart)*((value-istart)/(istop-istart))
}

// ColorToAttribute converts a color.Color to the closest equivalent termbox.Attribute.
// termbox.Output256 should set as Termbox's OutputMode.
func ColorToAttribute(c color.Color) termbox.Attribute {
	r16, g16, b16, _ := c.RGBA()

	r := r16 >> 8
	g := g16 >> 8
	b := b16 >> 8

	rLerped := uint8(lerp(float32(r), float32(0), float32(255), float32(0), float32(5)))
	gLerped := uint8(lerp(float32(g), float32(0), float32(255), float32(0), float32(5)))
	bLerped := uint8(lerp(float32(b), float32(0), float32(255), float32(0), float32(5)))

	// this color format is explained here: https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
	return termbox.Attribute(16 + 36*rLerped + 6*gLerped + bLerped + 1)
}

func draw(g *gif.GIF) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for y := 0; y < g.Config.Height; y++ {
		for x := 0; x < g.Config.Width; x++ {
			gifPixelColor := g.Image[0].At(x, y)
			termbox.SetCell(x, y, ' ', termbox.ColorDefault, ColorToAttribute(gifPixelColor))
		}
	}

	termbox.Flush()
}

func openGif(path string) (*gif.GIF, error) {
	gifBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	gifReader := bytes.NewReader(gifBytes)

	decodedGif, err := gif.DecodeAll(gifReader)
	if err != nil {
		return nil, err
	}

	return decodedGif, nil
}

func main() {
	gifPath := os.Args[1]
	inputGif, err := openGif(gifPath)
	if err != nil {
		log.Fatal(err)
	}

	err = termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)
	loop(inputGif)
}

func loop(g *gif.GIF) {
	events := make(chan termbox.Event)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()

	for {
		select {
		case e := <-events:
			if e.Type == termbox.EventKey && (e.Key == termbox.KeyEsc || e.Key == termbox.KeyCtrlC) {
				return
			}
		default:
			draw(g)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
