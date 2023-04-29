package main

import "fmt"

type Graphic interface {
	Draw()
}

type Image struct {
	FileName string
}

func (im Image) Draw() {
	fmt.Println("draw " + im.FileName)
}

type ImageProxy struct {
	FileName string
	_image   *Image
}

func (ip ImageProxy) GetImage() *Image {
	if ip._image == nil {
		ip._image = &Image{ip.FileName}
	}
	return ip._image
}

func (ip ImageProxy) Draw() {
	ip.GetImage().Draw()
}

func main() {
	proxy := ImageProxy{FileName: "1.png"}

	fileName := proxy.FileName

	proxy.Draw()
	fmt.Println(fileName)
}
