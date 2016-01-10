package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/unixpickle/gocube"
	"github.com/unixpickle/rubiksimg"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: scrambleimg <output.png> <scramble>")
		os.Exit(1)
	}

	cube := gocube.SolvedCubieCube()
	moves, err := gocube.ParseMoves(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid scramble.")
		os.Exit(1)
	}
	for _, move := range moves {
		cube.Move(move)
	}

	stickers := cube.StickerCube()
	image := rubiksimg.GenerateImage(1000, stickers)

	output, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to open output file:", err)
		os.Exit(1)
	}
	defer output.Close()

	if err := png.Encode(output, image); err != nil {
		fmt.Fprint(os.Stderr, "Failed to encode image:", err)
		os.Exit(1)
	}
}
