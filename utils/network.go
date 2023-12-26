package utils

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func Upload(zipData *bytes.Buffer, endpoint, unit_id string) error {
	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	zipFileWriter, err := multipartWriter.CreateFormFile("file", unit_id+".zip")
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

	resp, err := http.Post(endpoint, multipartWriter.FormDataContentType(), &requestBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status + string(body))
	}

	return nil
}

func Download(endpoint, path string) error {
	resp, err := http.Get(endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
