packages
--------

https://github.com/nsf/termbox-go and https://golang.org/pkg/image/gif do all the heavy lifting

why go?
-------
- batteries included standard library
- performance, type system, concurrency model (and feels sort of like Python)
- excellent tooling and editor integrations
- quick and easy to build and deploy self contained binaries
- Docker, k8s, Hashicorp

termbox-go
----------
* library for creating cross-platform text-based interfaces, like `ncurses`
  - http://ctop.sh
* small API, similar to working with HTML5 canvas
  - termbox-go gives you a grid of cells, one for every character your terminal can display
  - you control individual cells with the [SetCell function](http://godoc.org/github.com/nsf/termbox-go#SetCell)
  - [minimal example](https://github.com/nsf/termbox-go/blob/master/_demos/random_output.go)
  - provides a nice API for dealing with keyboard input and terminal colors, no need to worry about ANSI escape codes!

gif package
-----------
* [DecodeAll](https://golang.org/pkg/image/gif/#DecodeAll)
* [GIF type](https://golang.org/pkg/image/gif/#GIF)

```go
type GIF struct {
        Image     []*image.Paletted // The successive images.
        Delay     []int             // The successive delay times, one per frame, in 100ths of a second.
        LoopCount int               // The loop count.
        // Disposal is the successive disposal methods, one per frame. For
        // backwards compatibility, a nil Disposal is valid to pass to EncodeAll,
        // and implies that each frame's disposal method is 0 (no disposal
        // specified).
        Disposal []byte
        // Config is the global color table (palette), width and height. A nil or
        // empty-color.Palette Config.ColorModel means that each frame has its own
        // color table and there is no global color table. Each frame's bounds must
        // be within the rectangle defined by the two points (0, 0) and
        // (Config.Width, Config.Height).
        //
        // For backwards compatibility, a zero-valued Config is valid to pass to
        // EncodeAll, and implies that the overall GIF's width and height equals
        // the first frame's bounds' Rectangle.Max point.
        Config image.Config
        // BackgroundIndex is the background index in the global color table, for
        // use with the DisposalBackground disposal method.
        BackgroundIndex byte
}
```
