package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
)

func GetSecretValue(secretName string) (string, error) {
	f, err := os.Open("/run/secrets/" + secretName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Invoke(function string, data []byte) ([]byte, error) {
	resp, err := http.Post("http://gateway:8080/function/"+function, "application/octet-stream", bytes.NewReader(data))

	body, _ := ioutil.ReadAll(resp.Body)

	return body, err
}

func InvokeAsync(function string, data []byte, callback string) error {
	_, err := http.Post("http://gateway:8080/async-function/"+function, "application/octet-stream", bytes.NewReader(data))
	if err != nil {
		return err
	}
	return nil
}
