package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jesperfj/intense-go/pkg/iplookup"
	"log"
	"net"
	"net/http"
	"os"
)

type LocationInfo struct {
	Host     string
	IP       string
	Provider string
}

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
		addrs, err := net.LookupHost(c.Request.Host)
		loc := &iplookup.IPAPIResponse{}
		if err != nil {
			fmt.Println("Error looking up host")
		} else {
			if len(addrs) > 0 {
				loc = iplookup.Lookup(addrs[0])
			}
		}
		c.HTML(http.StatusOK, "index.tmpl.html", loc)
	})

	router.Run(":" + port)
}
