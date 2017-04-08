package iplookup

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

func Lookup(addr string) *IPAPIResponse {
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
