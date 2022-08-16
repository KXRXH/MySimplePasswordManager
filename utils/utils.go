package utils

import (
	"encoding/base64"
	"os"
	"path/filepath"
)

var Bool2Int = map[bool]int8{false: 0, true: 1}

func DecodeString(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func EncodeString(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func GetPwd() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
