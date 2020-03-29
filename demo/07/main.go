package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"time": time.Now().UTC(),
		})
	})

	fmt.Println("Hello World!")

	r.Run(":8080")
}
