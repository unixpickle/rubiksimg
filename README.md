# rubiksimg

This is a Go package for rendering a Rubik's cube scramble as an image. It is useful for programs which automatically generate scrambles, such as [dailycube](https://github.com/unixpickle/dailycube).

# Usage

After a simple `go get`, you can render a scramble like this:

```go
import (
	"github.com/unixpickle/gocube"
	"github.com/unixpickle/rubiksimg"
)

...

const ImageSize = 500

...

func main() {
    cube := gocube.SolvedCubieCube()
    // ... apply moves to the cube here ...
	  stickers := cube.StickerCube()
    image := rubiksimg.GenerateImage(ImageSize, stickers)
    // ... output the image to a file here ...
}
```

# `scrambleimg`

This comes with a command-line tool called `scrambleimg`. It generates a PNG file using a specified scramble string. It is useful for quickly telling what a scramble should look like.

# Code generated

A lot of this code (i.e. [background.go](background.go) and [stickers.go](stickers.go)) was generated programmatically based on an SVG file. The code which generated the code can be found in [generation/](generation).
