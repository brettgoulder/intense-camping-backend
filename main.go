package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

type IPAPIResponse struct {
	As          string  `json:"as"`          // "as": "AS16509 Amazon.com, Inc.",
	City        string  `json:"city"`        // "city: "Boardman",
	Country     string  `json:"country"`     // "country": "United States",
	CountryCode string  `json:"countryCode"` // "countryCode": "US",
	ISP         string  `json:"isp"`         // "isp: "Amazon Technologies",
	Lat         float32 `json:"lat"`         // "lat": 45.8696,
	Long        float32 `json:"lon"`         // "lon": -119.688,
	Org         string  `json:"org"`         // "org": "Amazon.com",
	Query       string  `json:"query"`       // "query": "52.32.217.169",
	Region      string  `json:"region"`      // "region": "OR",
	RegionName  string  `json:"regionName"`  // "regionName": "Oregon",
	Status      string  `json:"status"`      // "status": "success",
	Timezone    string  `json:"timezone"`    // "timezone": "America/Los_Angeles",
	Zip         string  `json:"zip"`         // "zip": "97818"
}

func lookup(addr string) *IPAPIResponse {
	res, err := http.Get("http://ip-api.com/json/" + addr)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	data := &IPAPIResponse{}
	err = json.Unmarshal(body, data)
	return data
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
		loc := &IPAPIResponse{}
		if err != nil {
			fmt.Println("Error looking up host")
		} else {
			if len(addrs) > 0 {
				loc = lookup(addrs[0])
			}
		}
		c.HTML(http.StatusOK, "index.tmpl.html", loc)
	})

	router.Run(":" + port)
}
