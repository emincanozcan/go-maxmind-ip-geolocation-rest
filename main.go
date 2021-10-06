package main

import (
	"encoding/json"
	"github.com/emincanozcan/go-maxmind-ip-geolocation-rest/services/search"
	"io/ioutil"
)

func main() {
	//data := ip_to_geolocation.Geolocation("192.168.1.1")
	//fmt.Println(data)

	locations := search.DataGenerator()
	file, _ := json.MarshalIndent(locations, "", "  ")
	_ = ioutil.WriteFile("data-src/data.json", file, 0777)
}
