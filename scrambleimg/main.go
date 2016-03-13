package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/unixpickle/gocube"
	"github.com/unixpickle/rubiksimg"
)

func main() {
	if len(os.Args) != 3 && len(os.Args) != 4 {
		dieUsage()
	}

	args := os.Args[1:]
	var caption bool
	if len(args) == 3 {
		if args[0] == "--caption" {
			caption = true
		} else {
			dieUsage()
		}
		args = args[1:]
	}

	cube := gocube.SolvedCubieCube()
	moves, err := gocube.ParseMoves(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid scramble.")
		os.Exit(1)
	}
	for _, move := range moves {
		cube.Move(move)
	}

	stickers := cube.StickerCube()
	var img image.Image
	if caption {
		img = rubiksimg.GenerateCaptionedImage(1000, 50, args[1], stickers)
	} else {
		img = rubiksimg.GenerateImage(1000, stickers)
	}

	output, err := os.Create(args[0])
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to open output file:", err)
		os.Exit(1)
	}
	defer output.Close()

	if err := png.Encode(output, img); err != nil {
		fmt.Fprint(os.Stderr, "Failed to encode image:", err)
		os.Exit(1)
	}
}

func dieUsage() {
	fmt.Fprintln(os.Stderr, "Usage: scrambleimg [--caption] <output.png> <scramble>")
	os.Exit(1)
}
