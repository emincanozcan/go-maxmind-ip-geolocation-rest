package location_list

import (
	"encoding/json"
	"sync"
)

var lList *locations
var lListAsJsonBytes []byte
var once1 sync.Once
var once2 sync.Once

func GetLocations() *locations {
	once1.Do(func() {
		lList = dataGenerator()
	})
	return lList
}

func GetLocationsAsJsonBytes() []byte {
	once2.Do(func() {
		locations := GetLocations()
		j, _ := json.Marshal(locations)
		lListAsJsonBytes = j
	})
	return lListAsJsonBytes
}
