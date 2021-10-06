package search

import (
	"encoding/csv"
	"os"
)

const CSVPath = "data-src/GeoLite2City.csv"

type Locations struct {
	Countries map[string]Country `json:"countries,omitempty"`
}

type Country struct {
	Name    string             `json:"name"`
	Iso     string             `json:"iso"`
	SubDivs map[string]SubDiv1 `json:"sub_divs,omitempty"`
}

type SubDiv1 struct {
	Name    string             `json:"name"`
	Iso     string             `json:"iso"`
	SubDivs map[string]SubDiv2 `json:"sub_divs"`
	Cities  map[string]City    `json:"cities,omitempty"`
}

type SubDiv2 struct {
	Name   string          `json:"name"`
	Iso    string          `json:"iso"`
	Cities map[string]City `json:"cities,omitempty"`
}

type City struct {
	Name string `json:"name"`
}

func DataGenerator() Locations {
	locations := Locations{
		Countries: make(map[string]Country),
	}
	f, _ := os.Open(CSVPath)
	defer f.Close()

	records, _ := csv.NewReader(f).ReadAll()

	for i := 0; i < len(records); i++ {
		countryIso := records[i][4]
		countryName := records[i][5]

		subDiv1Iso := records[i][6]
		subDiv1Name := records[i][7]

		subDiv2Iso := records[i][8]
		subDiv2Name := records[i][9]

		cityName := records[i][10]

		if _, ok := locations.Countries[countryIso]; !ok {

			if locations.Countries == nil {
				locations.Countries = make(map[string]Country)
			}

			locations.Countries[countryIso] = Country{
				Iso:     countryIso,
				Name:    countryName,
				SubDivs: make(map[string]SubDiv1),
			}
		}

		c := locations.Countries[countryIso]

		if _, ok := c.SubDivs[subDiv1Iso]; !ok {
			c.SubDivs[subDiv1Iso] = SubDiv1{
				Iso:     subDiv1Iso,
				Name:    subDiv1Name,
				SubDivs: make(map[string]SubDiv2),
				Cities:  make(map[string]City),
			}
		}

		s1 := c.SubDivs[subDiv1Iso]
		if len(subDiv2Name) > 0 {
			if _, ok := s1.SubDivs[subDiv2Iso]; !ok {
				s1.SubDivs[subDiv2Iso] = SubDiv2{
					Iso:    subDiv2Iso,
					Name:   subDiv2Name,
					Cities: make(map[string]City),
				}
			}

			s1.SubDivs[subDiv2Iso].Cities[cityName] = City{
				Name: cityName,
			}

		} else {
			s1.Cities[cityName] = City{
				Name: cityName,
			}
		}
	}
	return locations
}
