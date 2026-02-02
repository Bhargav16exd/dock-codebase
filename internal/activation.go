package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ActivateDockInput struct {
	ProductId  string `json:"product_id"`
	LicenseKey string `json:"license_key"`
	MacId      string `json:"mac_id"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var APPLICATION_JSON string = "application/json"

type ResponseType struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       Tokens `json:"data"`
}

func ActivateDock() {

	//fetch keys , send keys to activate api
	configs := GetConfig()

	payload := ActivateDockInput{
		ProductId:  configs.ProductId,
		LicenseKey: configs.DockLicenseKeyOne,
		MacId:      configs.MacId,
	}

	encodedPayload, err := json.Marshal(payload)

	if err != nil {
		fmt.Println(err)
	}

	bufferedPayload := bytes.NewBuffer(encodedPayload)

	encodedRes, err := http.Post(configs.ServerHost+configs.ApiActivationPath, APPLICATION_JSON, bufferedPayload)

	if err != nil {
		fmt.Println(err)
	}

	defer encodedRes.Body.Close()

	response := ResponseType{}

	json.NewDecoder(encodedRes.Body).Decode(&response)

	saveTokens(response.Data)
	updateConfig()
}

func saveTokens(tokens Tokens) error {

	file, err := os.OpenFile("./config/token.json",
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
	err = encoder.Encode(&tokens)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
