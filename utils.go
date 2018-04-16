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
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	return body, err
}

func InvokeAsync(function string, data []byte, callback string) error {
	post, err := http.NewRequest("POST", "http://gateway:8080/async-function/"+function, bytes.NewReader(data))
	if err != nil {
		return err
	}

	post.Header.Add("content-type", "application/octet-stream")
	if callback != "" {
		post.Header.Add("X-Callback-URL", callback)
	}

	_, err = http.DefaultClient.Do(post)

	if err != nil {
		return err
	}
	return nil
}
