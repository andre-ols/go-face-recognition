package entity

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math"
	"os"

	"github.com/Kagami/go-face"
)

type Drawer interface {
	DrawFace(face face.Face)
	loadImage(path string) error
	SaveImage(path string) error
}

type DrawerImpl struct {
	img *image.Image
	dst *image.RGBA
}

func (d *DrawerImpl) DrawFace(face face.Face) {
	d.dst = image.NewRGBA((*d.img).Bounds())
	rectColor := color.RGBA{0, 255, 0, 255}
	draw.Draw(d.dst, d.dst.Bounds(), (*d.img), (*d.img).Bounds().Min, draw.Src)
	d.drawLine(d.dst, face.Rectangle.Min.X, face.Rectangle.Min.Y, face.Rectangle.Max.X, face.Rectangle.Min.Y, rectColor)
	d.drawLine(d.dst, face.Rectangle.Max.X, face.Rectangle.Min.Y, face.Rectangle.Max.X, face.Rectangle.Max.Y, rectColor)
	d.drawLine(d.dst, face.Rectangle.Max.X, face.Rectangle.Max.Y, face.Rectangle.Min.X, face.Rectangle.Max.Y, rectColor)
	d.drawLine(d.dst, face.Rectangle.Min.X, face.Rectangle.Max.Y, face.Rectangle.Min.X, face.Rectangle.Min.Y, rectColor)
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

func NewDrawer(imagePath string) Drawer {
	drawer := &DrawerImpl{}
	drawer.loadImage(imagePath)
	return drawer
}
