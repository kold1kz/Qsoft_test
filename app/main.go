package main

import (
	"embed"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("tmpl/*.html")
	router.Use(Logger())

	router.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
		  if c.Request.Header.Get("X-PING") == "ping" {
			c.Writer.Header().Set("X-PONG", "pong")
		  }
		}
	  }()).GET("/when/:year", func(c *gin.Context) {
		param := c.Param("year")
		c.HTML(
			http.StatusOK,

			"index.html",

			gin.H{
				"param": daywhat(param),
			},
		)
	})


	router.Run()
}

func daywhat(param string) string {
	var r string = "error"
	date := time.Now()
	day := date.Day()
	month := int(date.Month())
	year := date.Year()

	if param[0] != '-' {
		if n, err := strconv.Atoi(param); err == nil {
			a := Date(year, month, day)
			b := Date(n, 1, 1)
			c := int(a.Sub(b).Hours() / 24)
			if n > year {
				r = "Days left: " + strconv.Itoa(c*-1)
			} else {
				r = "Days gone: " + strconv.Itoa(c)
			}
		} else {
			log.Println(param, "не является целым числом.")
		}
	} else {
		log.Println(param, "отрицательное число.")
	}
	return r
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
