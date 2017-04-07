package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		fmt.Println("Host: " + c.Request.Host)
		addrs, err := net.LookupHost(c.Request.Host)
		if err {
			fmt.Println("Error looking up host")
		} else {
			if len(addrs) > 0 {
				fmt.Println("This app reached via IP address " + addrs[0])
			}
		}

		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)
}
