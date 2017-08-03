package main

import (
	"image/color"
	"testing"

	"github.com/nsf/termbox-go"
)

func TestColorToAttribute(t *testing.T) {
	cases := []struct {
		input  color.Color
		output termbox.Attribute
	}{
		{color.RGBA{0, 0, 0, 0}, termbox.Attribute(17)},
		{color.RGBA{255, 255, 255, 0}, termbox.Attribute(232)},
		{color.RGBA{255, 0, 0, 0}, termbox.Attribute(197)},
		{color.RGBA{0, 255, 0, 0}, termbox.Attribute(47)},
		{color.RGBA{0, 0, 255, 0}, termbox.Attribute(22)},
		{color.RGBA{255, 175, 0, 0}, termbox.Attribute(215)},
	}

	for _, c := range cases {
		result := ColorToAttribute(c.input)
		if result != c.output {
			t.Errorf("For input %v got %v, expecting %v", c.input, result, c.output)
		}
	}
}
