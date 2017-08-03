package main

import (
	"bytes"
	"image"
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

func drawImage(i image.Image) {
	size := i.Bounds().Size()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			pixelColor := i.At(x, y)
			termbox.SetCell(x, y, ' ', termbox.ColorDefault, ColorToAttribute(pixelColor))
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

func loop(g *gif.GIF) {
	events := make(chan termbox.Event)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()

	currentImageIndex := 0

	for {
		select {
		case e := <-events:
			if e.Type == termbox.EventKey && (e.Key == termbox.KeyEsc || e.Key == termbox.KeyCtrlC) {
				return
			}
		default:
			drawImage(g.Image[currentImageIndex])
			time.Sleep(10 * time.Duration(g.Delay[currentImageIndex]) * time.Millisecond)

			currentImageIndex++
			if currentImageIndex >= len(g.Image) {
				currentImageIndex = 0
			}
		}
	}
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
