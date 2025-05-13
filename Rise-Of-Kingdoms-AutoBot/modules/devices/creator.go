package devices

import (
	"Rise-Of-Kingdoms-AutoBot/models"
	"encoding/json"
	"os"
)

// func LoadBotConfig() ([]models.Bot, error) {
// 	config := []models.Bot{}
// 	file, err := os.Open("bots_config.json")
// 	if err != nil {
// 		return config, err
// 	}
// 	defer file.Close()

// 	decoder := json.NewDecoder(file)
// 	err = decoder.Decode(&config)
// 	if err != nil {
// 		return config, err
// 	}
// 	return config, nil
// }

// LoadDeviceConfig loads the device configuration from a JSON file
func LoadDeviceConfig() ([]models.Device, error) {
	config := []models.Device{}
	file, err := os.Open("devices_config.json")
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}
