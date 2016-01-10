# rubiksimg

This is a Go package for generating images given a Rubik's cube scramble. It is useful for programs which automatically generate scrambles.

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

# The scrambleimg command

This comes with a command called `scrambleimg` which takes a scramble string and outputs it to a png file. It is useful for quickly telling what a cube looks like.

# This code was generated

A lot of this code (i.e. [background.go](background.go) and [stickers.go](stickers.go)) was generated programmatically based on an SVG file. The code which generated the code can be found in [generation/](generation).
