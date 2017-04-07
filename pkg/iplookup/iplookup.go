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
