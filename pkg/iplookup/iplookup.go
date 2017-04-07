package iplookup

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type WhoisResponse1 struct {
	Nets Nets1 `json:"nets"`
}

type WhoisResponse2 struct {
	Nets Nets2 `json:"nets"`
}

type Nets1 struct {
	Net Net `json:"net"`
}

type Nets2 struct {
	Net []Net `json:"net"`
}

type Net struct {
	OrgRef OrgRef `json:"orgRef"`
}

type OrgRef struct {
	Name string `json:"@name"`
}

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

func Lookup2() *IPAPIResponse {
	res, err := http.Get("http://ip-api.com/json")
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

func Lookup(addr string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://whois.arin.net/rest/nets;q="+addr+"?showDetails=true&showARIN=false&showNonArinTopLevelNet=false&ext=netref2", nil)
	req.Header.Set("Accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	r1 := new(WhoisResponse1)
	err = json.Unmarshal(body, &r1)
	if r1.Nets.Net.OrgRef.Name != "" {
		return r1.Nets.Net.OrgRef.Name
	} else {
		r2 := new(WhoisResponse2)
		err = json.Unmarshal(body, &r2)
		return r2.Nets.Net[0].OrgRef.Name
	}
}
