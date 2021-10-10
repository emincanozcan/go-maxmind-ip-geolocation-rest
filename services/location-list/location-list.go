package location_list

import (
	"encoding/json"
	"fmt"
	"sync"
)

var lList *locations
var lListAsJsonBytes []byte
var once1 sync.Once
var once2 sync.Once

func GetLocations() *locations {
	once1.Do(func() {
		fmt.Println("once 2")
		lList = dataGenerator()
	})
	return lList
}

func GetLocationsAsJsonBytes() []byte {
	once2.Do(func() {
		fmt.Println("once 1")
		locations := GetLocations()
		j, _ := json.Marshal(locations)
		lListAsJsonBytes = j
	})
	return lListAsJsonBytes
}
