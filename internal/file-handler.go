package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
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

func CheckForDataFromServer() {

	for {

		// TBD - send token with Request
		resp, err := http.Get("http://localhost:3000/api/files/check")

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

		fetchAndWriteFile(file)
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

	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:9090")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer conn.Close()

	AuthData := Frame{
		ProductId:        GetConfig().ProductId,
		Token:            "token",
		FrameMessageType: MessageTypeAuth.String(),
	}

	authBytes, err := json.Marshal(AuthData)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// TBD - send auth request
	conn.Write(authBytes)
	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	//request file
	err = encoder.Encode(Frame{
		FrameMessageType: MessageTypeRequestFile.String(),
		FileMetaData: FileMetaData{
			FileName: fileName,
		},
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// keep connection alive if expecting response
	time.Sleep(2 * time.Second)

	resp := Frame{}
	decoder.Decode(&resp)

	if resp.FrameMessageType == MessageTypeFile.String() {

		f, err := os.OpenFile(
			"./backups/"+resp.FileMetaData.FileName,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0777,
		)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		_, err = f.Write(resp.Payload)

		if err != nil {
			fmt.Println(err)
		}
	}
}
