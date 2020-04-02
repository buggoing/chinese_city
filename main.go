package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type City struct {
	Name       string `json:"name"`
	Admaster   string `json:"admaster"`
	NameEn     string `json:"name_en"`
	Province   string `json:"province"`
	ProvinceEn string `json:"province_en"`
	Tier       string `json:"tier"`
	Longitude  string `json:"lng"`
	Latitude   string `json:"lat"`
	Area       string `json:"area"`
	City       string `josn:"city"`
}

func main() {
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	cities := make(map[string]City, 0)
	if err := json.Unmarshal([]byte(file), &cities); err != nil {
		log.Fatal(err)
	}
	file, err = ioutil.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	citiesGeo := make([]City, 0)
	if err := json.Unmarshal([]byte(file), &citiesGeo); err != nil {
		log.Fatal(err)
	}
	fmt.Println("city_geo0: ", citiesGeo[0])

	res := make(map[string]City, 0)
	for k, city := range cities {
		for _, cityGeo := range citiesGeo {
			if city.Admaster == cityGeo.Area || city.Admaster == cityGeo.City+cityGeo.Area || city.Admaster == cityGeo.Province+cityGeo.Area {
				city.Longitude = cityGeo.Longitude
				city.Latitude = cityGeo.Latitude
				cities[k] = city
				res[city.Name] = city
			}
		}
		fmt.Printf("city: %+v\n", city)
	}
	jsonData, err := json.MarshalIndent(cities, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	jsonFile, err := os.Create("city_geo.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)

	// fmt.Printf("city: %v", cities)
}
