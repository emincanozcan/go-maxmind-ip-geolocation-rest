package ip_to_geolocation

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

type GeolocationData struct {
	countryIsoCode string
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

	return GeolocationData{
		countryIsoCode: record.Country.IsoCode,
	}
}
