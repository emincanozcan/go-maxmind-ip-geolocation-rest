package location_list

import (
	"encoding/csv"
	"os"
	"strconv"
)

const CSVPath = "data-src/GeoLite2City.csv"

type locations struct {
	Countries map[string]*country `json:"countries,omitempty"`
}
type country struct {
	Name    string              `json:"name"`
	SubDivs map[string]*subDiv1 `json:"sub_divs,omitempty"`
}
type subDiv1 struct {
	Name    string              `json:"name"`
	SubDivs map[string]*subDiv2 `json:"sub_divs"`
	Cities  map[uint32]string   `json:"cities,omitempty"`
}
type subDiv2 struct {
	Name   string            `json:"name"`
	Cities map[uint32]string `json:"cities,omitempty"`
}

func (l *locations) addCountry(iso string, name string) *country {
	if _, ok := l.Countries[iso]; !ok {
		if l.Countries == nil {
			l.Countries = make(map[string]*country)
		}

		l.Countries[iso] = &country{
			Name: name,
		}
	}

	return l.Countries[iso]
}

func (c *country) addSubDiv(iso string, name string) *subDiv1 {
	if _, ok := c.SubDivs[iso]; !ok {
		if c.SubDivs == nil {
			c.SubDivs = make(map[string]*subDiv1)
		}
		c.SubDivs[iso] = &subDiv1{
			Name: name,
		}
	}

	return c.SubDivs[iso]
}

func (s *subDiv1) addSubDiv(iso string, name string) *subDiv2 {
	if _, ok := s.SubDivs[iso]; !ok {

		if s.SubDivs == nil {
			s.SubDivs = make(map[string]*subDiv2)
		}
		s.SubDivs[iso] = &subDiv2{
			Name: name,
		}
	}

	return s.SubDivs[iso]
}

func (s *subDiv1) addCity(geoId uint32, name string) {
	if s.Cities == nil {
		s.Cities = make(map[uint32]string)
	}
	s.Cities[geoId] = name
}

func (s *subDiv2) addCity(geoId uint32, name string) {
	if s.Cities == nil {
		s.Cities = make(map[uint32]string)
	}
	s.Cities[geoId] = name
}

func dataGenerator() *locations {
	locations := locations{}

	f, _ := os.Open(CSVPath)
	defer f.Close()

	records, _ := csv.NewReader(f).ReadAll()
	for i := 1; i < len(records); i++ {
		tmp, _ := strconv.ParseUint(records[i][0], 10, 32)
		geoId := uint32(tmp)

		countryIso := records[i][4]
		countryName := records[i][5]
		subDiv1Iso := records[i][6]
		subDiv1Name := records[i][7]
		subDiv2Iso := records[i][8]
		subDiv2Name := records[i][9]
		cityName := records[i][10]

		if len(countryIso) < 1 {
			continue
		}
		c := locations.addCountry(countryIso, countryName)

		if len(subDiv1Iso) < 1 {
			continue
		}
		s1 := c.addSubDiv(subDiv1Iso, subDiv1Name)

		if len(subDiv2Name) > 0 {
			s2 := s1.addSubDiv(subDiv2Iso, subDiv2Iso)
			if len(cityName) > 0 {
				s2.addCity(geoId, cityName)
			}
		} else {
			if len(cityName) > 0 {
				s1.addCity(geoId, cityName)
			}
		}
	}
	return &locations
}
