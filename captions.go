package rubiksimg

import (
	"image"
	"image/color"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

var defaultFontData = draw2d.FontData{
	Name:   "Roboto",
	Family: draw2d.FontFamilySans,
	Style:  draw2d.FontStyleNormal,
}

var fontDirectory = "fonts"

const lineHeightRatio = 1.4
const captionSideMargin = 0.02

func generateCaptionImage(width int, fontSize float64, caption string) image.Image {
	lines := wrapText(fontSize, caption, float64(width)*(1-captionSideMargin*2))
	height := wrappedTextHeight(fontSize, lines)
	dest := image.NewRGBA(image.Rect(0, 0, width, int(height)))
	ctx := draw2dimg.NewGraphicContext(dest)

	draw2d.SetFontFolder(fontFolder())
	ctx.SetFontData(defaultFontData)
	ctx.SetFontSize(fontSize)
	ctx.SetFillColor(color.RGBA{0, 0, 0, 255})

	var y float64
	for _, line := range lines {
		lineWidth, lineHeight := measureText(fontSize, line)
		lineMargin := lineHeight * (lineHeightRatio - 1) / 2
		y += lineHeight + lineMargin
		ctx.FillStringAt(line, (float64(width)-lineWidth)/2, y)
		y += lineMargin
	}

	return dest
}

func wrappedTextHeight(fontSize float64, lines []string) float64 {
	var height float64
	for _, line := range lines {
		_, h := measureText(fontSize, line)
		height += h * lineHeightRatio
	}
	return height
}

func wrapText(fontSize float64, text string, maxWidth float64) []string {
	fields := strings.Fields(text)
	lines := make([]string, 0, len(fields))
	currentLine := []string{}
	for _, field := range fields {
		currentLine = append(currentLine, field)
		newText := strings.Join(currentLine, " ")
		if width, _ := measureText(fontSize, newText); width > maxWidth {
			oldLine := currentLine
			currentLine = []string{}
			if len(oldLine) > 1 {
				oldLine = oldLine[:len(oldLine)-1]
				currentLine = []string{field}
			}
			lines = append(lines, strings.Join(oldLine, " "))
		}
	}
	if len(currentLine) > 0 {
		lines = append(lines, strings.Join(currentLine, " "))
	}
	return lines
}

func measureText(fontSize float64, text string) (width, height float64) {
	dest := image.NewRGBA(image.Rect(0, 0, 1, 1))
	ctx := draw2dimg.NewGraphicContext(dest)

	fontData := draw2d.FontData{
		Name:   "Roboto",
		Family: draw2d.FontFamilySans,
		Style:  draw2d.FontStyleNormal,
	}
	draw2d.SetFontFolder(fontFolder())
	ctx.SetFontData(fontData)
	ctx.SetFontSize(fontSize)

	left, top, right, bottom := ctx.GetStringBounds(text)
	width = right - left
	height = bottom - top
	return
}

func fontFolder() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), fontDirectory)
}
