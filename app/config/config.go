package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ConfigService struct {
	Character   []string `json:"character"`
	Footwear    []string `json:"footwear"`
	Pants       []string `json:"pants"`
	Torso       []string `json:"torso"`
	Eyewear     []string `json:"eyewear"`
	Headwear    []string `json:"headwear"`
	WeaponRight []string `json:"weaponright"`
	WeaponLeft  []string `json:"weaponleft"`
}

func NewConfigService(configPath string) *ConfigService {
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var service ConfigService

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &service)

	return &service
}
