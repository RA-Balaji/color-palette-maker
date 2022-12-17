package main

import (
	"github.com/gin-gonic/gin"
	"github.com/RA-Balaji/color-palette-maker/api"
)

func main() {

	router := gin.Default()	
	routerGrp := router.Group("/image")
	routerGrp.GET("/palette", api.GetPaletteFromImage)

	router.Run("localhost:8088")
}