package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
    fmt.Println("Test output")

    r := gin.Default()
    r.Static("/public", "./public")
    r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "testbubble",
        })
	})
	r.Run()
}

