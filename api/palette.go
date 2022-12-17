package api

import (
	"errors"
	"image"
	"image/png"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mccutchen/palettor"
	"github.com/nfnt/resize"
	"github.com/RA-Balaji/color-palette-maker/model"
)

const (
	numOfColors = 6
	maxIterations = 100
)

func GetPaletteFromImage(c *gin.Context) {

	var data model.UrlReq
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("url"))
	}

	image, err := loadImageFromURL(data.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	log.Println("Image bounds:", image.Bounds())
	imgResized := resize.Thumbnail(200, 200, image, resize.Lanczos3)
    palette, err := palettor.Extract(numOfColors, maxIterations, imgResized)
	if err != nil {
        log.Fatalf("image too small")
    }
    for _, color := range palette.Colors() {
        log.Printf("color: %v; weight: %v", color, palette.Weight(color))
    }
}

func loadImageFromURL(url string) (image.Image, error) {
    //Get the response bytes from the url
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    if response.StatusCode != 200 {
        return nil, errors.New("received non 200 response code")
    }

	log.Println("response body=", response.Body)
    img, err := png.Decode(response.Body)
    if err != nil {
        return nil, err
    }

    return img, nil
} 