package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadAPI(zipData *bytes.Buffer, endpoint, unitID, apiKey string) error {
	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	zipFileWriter, err := multipartWriter.CreateFormFile("file", unitID+".zip")
	if err != nil {
		return err
	}

	_, err = io.Copy(zipFileWriter, zipData)
	if err != nil {
		return err
	}

	err = multipartWriter.Close()
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", endpoint, &requestBody)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", apiKey)
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		return errors.New(response.Status + string(body))
	}

	return nil
}

func DownloadAPI(endpoint, path, api_key string) error {
	request, err := http.NewRequest("GET", endpoint, nil)
	request.Header.Set("Authorization", api_key)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	return err
}

func GetRemoteUnits(ip, port, api_key string) ([]Unit, error) {
	var units []Unit

	endpoint := "http://" + ip + ":" + port + "/units"

	request, err := http.NewRequest("GET", endpoint, nil)
	request.Header.Set("Authorization", api_key)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return units, errors.New("Error making GET request")
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return units, errors.New("Error reading response body:" + err.Error())
	}

	if response.StatusCode != http.StatusOK {
		return units, errors.New("Error:" + response.Status)
	}

	err = json.Unmarshal(body, &units)
	if err != nil {
		return units, errors.New("Error unmarshaling JSON:" + err.Error())
	}

	return units, nil
}

type Payload struct {
	Version int
}

func GetUnitVersion(ip, port, api_key string, unit_id string) (int, error) {
	var payload Payload
	endpoint := "http://" + ip + ":" + port + "/version/" + unit_id

	request, err := http.NewRequest("GET", endpoint, nil)
	request.Header.Set("Authorization", api_key)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return payload.Version, errors.New("Error making GET request")
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return payload.Version, errors.New("Error reading response body:" + err.Error())
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		return payload.Version, errors.New("Error unmarshaling JSON:" + err.Error())
	}

	return payload.Version, nil
}
