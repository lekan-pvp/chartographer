package imageservice

import (
	"encoding/base64"
	"golang.org/x/image/bmp"
	"golang.org/x/image/draw"
	"image"
	"internshipApplicationTemplate/internal/config"
	"log"
	"os"
	"strconv"
	"strings"
)

func SaveFragment(id string, x string, y string, width string, height string, data []byte) error {
	var readerDst *os.File
	var err error
	for _, idFile := range r {
		if id == idFile {
			fileName := id + ".bmp"
			filePath := config.Cfg.FilePath + fileName
			readerDst, err = os.Open(filePath)
			if err != nil {
				log.Println("file not found")
				return err
			}
		}
	}

	dst, err := bmp.Decode(readerDst)
	if err != nil {
		log.Println("dest file decode error")
		return err
	}

	xFragment, err := strconv.Atoi(x)
	if err != nil {
		return err
	}

	yFragment, err := strconv.Atoi(y)
	if err != nil {
		return err
	}

	widthFragment, err := strconv.Atoi(width)
	if err != nil {
		return err
	}

	heightFragment, err := strconv.Atoi(height)
	if err != nil {
		return err
	}

	readerSrc := base64.NewDecoder(base64.StdEncoding, strings.NewReader(string(data)))
	src, err := bmp.Decode(readerSrc)
	if err != nil {
		log.Println("sourse file decode error")
		return err
	}

	sp2tl := image.Point{X: xFragment, Y: yFragment}
	sp2bd := image.Point{X: xFragment + widthFragment, Y: yFragment + heightFragment}
	r2 := image.Rectangle{Min: sp2tl, Max: sp2bd}
	r1 := image.Rectangle{Min: image.Point{0, 0}, Max: r2.Max}

	rgba := image.NewRGBA(r1)

	draw.Draw(rgba, dst.Bounds(), dst, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r2, src, sp2tl, draw.Src)

	return nil
}
