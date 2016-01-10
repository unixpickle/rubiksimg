package rubiksimg

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/unixpickle/gocube"
)

func GenerateImage(size int, cube gocube.StickerCube) image.Image {
	dest := image.NewRGBA(image.Rect(0, 0, size, size))
	ctx := draw2dimg.NewGraphicContext(dest)

	scale := float64(size) / 500
	drawBackground(ctx, scale)

	drawTopFace(ctx, scale, colorsForStickers(cube[:9]))
	drawFrontFace(ctx, scale, colorsForStickers(cube[9*2:9*3]))
	drawRightFace(ctx, scale, colorsForStickers(cube[9*4:9*5]))

	return dest
}

func colorsForStickers(nums []int) []color.Color {
	res := make([]color.Color, len(nums))
	for i, num := range nums {
		res[i] = colorForSticker(num)
	}
	return res
}

func colorForSticker(num int) color.Color {
	return []color.RGBA{
		{0xff, 0xff, 0xff, 0xff},
		{0xff, 0xff, 0x00, 0xff},
		{0x38, 0xd8, 0x38, 0xff},
		{0x3b, 0x6d, 0xe0, 0xff},
		{0xe9, 0x36, 0x4b, 0xff},
		{0xff, 0xa5, 0x00, 0xff},
	}[num-1]
}
