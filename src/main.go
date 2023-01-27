package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()
	router.LoadHTMLGlob("tmpl/*")
	
	router.Static("/tmpl", "./tmpl/")

	router.GET("when/:year", func(c *gin.Context) {
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
			fmt.Println(param, "не является целым числом.")
		}
	} else {
		fmt.Println(param, "отрицательное число.")
	}
	return r
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
