package ip_to_geolocation

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

type GeolocationData struct {
	CountryIso        string `json:"country_iso"`
	SubDiv1Iso        string `json:"sub_div_1_iso"`
	SubDiv2Iso        string `json:"sub_div_2_iso"`
	CityGeolocationId uint   `json:"city_geolocation_id"`
}

const MMDBPath string = "data-src/GeoLite2City.mmdb"

func Geolocation(ipAddr string) GeolocationData {
	// TODO: measure performance, opening file for every request probably a bad idea.
	db, err := geoip2.Open(MMDBPath)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	ip := net.ParseIP(ipAddr)
	record, err := db.City(ip)

	if err != nil {
		log.Fatal(err)
	}

	data := GeolocationData{
		CountryIso:        record.Country.IsoCode,
		CityGeolocationId: record.City.GeoNameID,
	}

	if len(record.Subdivisions) > 0 {
		data.SubDiv1Iso = record.Subdivisions[0].IsoCode
	}

	if len(record.Subdivisions) > 1 {
		data.SubDiv2Iso = record.Subdivisions[1].IsoCode
	}

	return data
}
