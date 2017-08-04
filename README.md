Termgif
=======

Display GIFs in your terminal!

![megaman](https://github.com/jrab89/termgif/raw/master/_example_gifs/megaman_256x256px.gif)

Installation
------------

```
$ go get -u github.com/jrab89/termgif
# $GOPATH/bin needs to be in your $PATH
$ export PATH="$GOPATH/bin:$PATH"
```

Usage
-----

```
$ termgif SOME_GIF
```

`SOME_GIF` can be a path to a GIF file or a URL. URLs are assumed to start with `http://` or `https://`. For example:

```
$ termgif https://github.com/jrab89/termgif/raw/master/_example_gifs/doge_50x50px.gif
...
$ termgif _example_gifs/megaman_256x256px.gif
...
```

Press CTRL+C or ESC to exit.
