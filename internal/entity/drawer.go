package entity

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Drawer interface {
	DrawFace(rectangle image.Rectangle, name string)
	drawLine(img draw.Image, x0, y0, x1, y1 int, color color.RGBA)
	drawName(txt string, x, y int, h float64)
	loadImage(path string) error
	loadFont(path string) error
	SaveImage(path string) error
}

type DrawerImpl struct {
	img   *image.Image
	dst   *image.RGBA
	font  *truetype.Font
	color color.RGBA
}

func (d *DrawerImpl) DrawFace(rectangle image.Rectangle, name string) {
	d.drawLine(d.dst, rectangle.Min.X, rectangle.Min.Y, rectangle.Max.X, rectangle.Min.Y, d.color)
	d.drawLine(d.dst, rectangle.Max.X, rectangle.Min.Y, rectangle.Max.X, rectangle.Max.Y, d.color)
	d.drawLine(d.dst, rectangle.Max.X, rectangle.Max.Y, rectangle.Min.X, rectangle.Max.Y, d.color)
	d.drawLine(d.dst, rectangle.Min.X, rectangle.Max.Y, rectangle.Min.X, rectangle.Min.Y, d.color)

	d.drawName(name, rectangle.Min.X, rectangle.Max.Y, float64(rectangle.Max.Y-rectangle.Min.Y))
}

func (d *DrawerImpl) drawLine(img draw.Image, x0, y0, x1, y1 int, color color.RGBA) {
	dx := math.Abs(float64(x1 - x0))
	dy := -math.Abs(float64(y1 - y0))
	sx := 1
	sy := 1
	if x0 > x1 {
		sx = -1
	}
	if y0 > y1 {
		sy = -1
	}
	err := dx + dy
	for {
		img.Set(x0, y0, color)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x0 += sx
		}
		if e2 <= dx {
			err += dx
			y0 += sy
		}
	}
}

func (d *DrawerImpl) drawName(txt string, x, y int, h float64) {

	// Define the font size as a fraction of the face height
	fontSize := h / 3

	dr := &font.Drawer{
		Dst:  d.dst,
		Src:  image.NewUniform(d.color),
		Face: truetype.NewFace(d.font, &truetype.Options{Size: fontSize, Hinting: font.HintingNone, DPI: 72.0}),
	}

	dr.Dot = fixed.Point26_6{
		X: fixed.I(x),
		Y: fixed.I(y + int(fontSize)),
	}

	dr.DrawString(txt)
}

func (d *DrawerImpl) loadFont(path string) error {
	fontBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	f, err := truetype.Parse(fontBytes)

	if err != nil {
		return err
	}

	d.font = f

	return nil
}

func (d *DrawerImpl) loadImage(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	img, err := jpeg.Decode(f)

	if err != nil {
		return err
	}

	d.img = &img

	return nil

}

func (d *DrawerImpl) SaveImage(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	return jpeg.Encode(f, d.dst, nil)
}

func NewDrawer(imagePath, fontPath string) Drawer {
	drawer := &DrawerImpl{}
	drawer.loadImage(imagePath)
	drawer.loadFont(fontPath)
	drawer.dst = image.NewRGBA((*drawer.img).Bounds())
	draw.Draw(drawer.dst, drawer.dst.Bounds(), (*drawer.img), (*drawer.img).Bounds().Min, draw.Src)

	drawer.color = color.RGBA{0, 255, 0, 255}

	return drawer
}
