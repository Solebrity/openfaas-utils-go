package utils

import (
	"io/ioutil"
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
