package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func MakeRequest(url, filePath, max string, insensitive bool) error {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return fmt.Errorf("open file failed: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("error getting file stat: %v", err)
	}

	fileUpload, err := writer.CreateFormFile("file", stat.Name())
	if err != nil {
		return fmt.Errorf("create form field failed: %v", err)
	}

	contentUpload, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("read upload file failed: %v", err)
	}
	fileUpload.Write(contentUpload)

	if insensitive {
		writer.WriteField("insensitive", "1")
	}
	writer.WriteField("max", max)
	writer.Close()

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return fmt.Errorf("create request error: %v", err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %v", err)
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println(string(content))
	return nil
}
