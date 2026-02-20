package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	fileNames []string
	mu        sync.Mutex
)

type SuccessResponseType struct {
	Success    bool     `json:"success"`
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Data       []string `json:"data,omitempty"`
}

var (
	config ConfigType
	once   sync.Once
)

func FetchConfig() ConfigType {
	once.Do(func() {
		config = GetConfig()
	})
	return config
}

const PRODUCTION string = "PRODUCTION"
const DEV string = "DEV"

func CheckForDataFromServer() {

	for {

		// TBD - send token with Request
		resp, err := http.Get(FetchConfig().ServerHost + FetchConfig().ApiFileCheckPath)

		if err != nil {
			fmt.Println(err)
		}

		info, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
		}

		response := SuccessResponseType{}

		err = json.Unmarshal(info, &response)

		if err != nil {
			fmt.Println(err)
		}

		fileNames = append(fileNames, response.Data...)

		time.Sleep(time.Minute)
	}
}

func CheckForFilesAvailable() {
	for {
		file := popFile()
		if file == "" {
			time.Sleep(time.Second * 5)
			continue
		}
		fmt.Println("here 3")
		go fetchAndWriteFile(file)
	}
}

func popFile() string {
	mu.Lock()
	defer mu.Unlock()

	if len(fileNames) == 0 {
		return ""
	}

	file := fileNames[0]
	fileNames = fileNames[1:]
	return file
}

func fetchAndWriteFile(fileName string) {

	resp, err := http.Get(FetchConfig().ServerHost + FetchConfig().ApiDownloadFilePath + fileName)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	file, err := os.Create("./backups/" + fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	if GetConfig().Environment == PRODUCTION {
		go DeleteFileFromServer(fileName)
	}
}

func DeleteFileFromServer(filename string) {

	req, err := http.NewRequest(http.MethodDelete, FetchConfig().ServerHost+FetchConfig().ApiDeleteFilePath+filename, nil)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println(resp)
}
