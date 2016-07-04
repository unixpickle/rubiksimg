package rubiksimg

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/unixpickle/gocube"
)

func GenerateCaptionedImage(size int, fontSize float64, caption string,
	cube gocube.StickerCube) image.Image {
	dest := image.NewRGBA(image.Rect(0, 0, size, size))
	ctx := draw2dimg.NewGraphicContext(dest)

	captionImage := generateCaptionImage(size, fontSize, caption)
	captionHeight := float64(captionImage.Bounds().Dy())
	remainingSize := float64(size) - captionHeight

	ctx.DrawImage(captionImage)

	cubeImage := GenerateImage(int(remainingSize), cube)
	ctx.Translate((float64(size)-remainingSize)/2, captionHeight)
	ctx.DrawImage(cubeImage)

	return dest
}

func GenerateImage(size int, cube gocube.StickerCube) image.Image {
	dest := image.NewRGBA(image.Rect(0, 0, size, size))
	ctx := draw2dimg.NewGraphicContext(dest)

	scale := float64(size) / 500
	ctx.Scale(-1, 1)
	ctx.Translate(-float64(size), 0)

	drawBackground(ctx, scale)
	drawTopFace(ctx, scale, mirrorDiagonally(colorsForStickers(cube[:9])))
	drawRightFace(ctx, scale, mirrorHorizontally(colorsForStickers(cube[9*4:9*5])))
	drawFrontFace(ctx, scale, mirrorHorizontally(colorsForStickers(cube[9*2:9*3])))

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

func mirrorHorizontally(colors []color.Color) []color.Color {
	colors[0], colors[2] = colors[2], colors[0]
	colors[3], colors[5] = colors[5], colors[3]
	colors[6], colors[8] = colors[8], colors[6]
	return colors
}

func mirrorDiagonally(colors []color.Color) []color.Color {
	colors[1], colors[3] = colors[3], colors[1]
	colors[2], colors[6] = colors[6], colors[2]
	colors[5], colors[7] = colors[7], colors[5]
	return colors
}
