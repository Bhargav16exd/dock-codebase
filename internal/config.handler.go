package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigType struct {
	ProductId         string `json:"product_id"`
	DockLicenseKeyOne string `json:"dock_license_key_one"`
	DockLicenseKeyTwo string `json:"dock_license_key_two"`
	IsDockActive      bool   `json:"is_dock_active"`
	ActivationApi     string `json:"activation_api"`
	MacId             string `json:"mac_id"`
}

func GetConfig() ConfigType {

	Config := ConfigType{}
	file, err := os.Open("./config/setup.json")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&Config)

	if err != nil {
		fmt.Println(err)
	}

	return Config
}

func updateConfig() error {

	config := GetConfig()

	config.IsDockActive = true

	file, err := os.OpenFile("./config/setup.json",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", "	")
	err = encoder.Encode(&config)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
