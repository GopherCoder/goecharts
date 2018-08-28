package main

import (
	"net/http"

	"goecharts/infra/option"

	"github.com/gin-gonic/gin"
)

type Chart struct {
	Theme  string
	Width  string
	Height string
}

func main() {

	var chart = Chart{
		Theme:  "light",
		Width:  "1200",
		Height: "800",
	}
	router := gin.Default()
	router.Static("/js", "./static/js")
	router.LoadHTMLGlob("./static/html/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", chart)
	})

	testChart := option.ChartBase{
		YAxis: option.YAxis{
			Name: "销量",
		},
		Title: option.Title{
			Text: "非常牛逼的图",
			Show: true,
			TextStyle: option.TextStyle{
				Color: "#CC00FF",
			},
		},
	}

	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.html", testChart)
	})

	router.Run(":8088")
}
