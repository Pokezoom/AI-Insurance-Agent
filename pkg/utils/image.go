package utils

import (
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func GetImageBase64(path string) (string, string, error) {
	var imgByte []byte

	if strings.Contains(path, "http") {
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get(path)
		if err != nil {
			return "", "", err
		}
		defer resp.Body.Close()
		imgByte, _ = io.ReadAll(resp.Body)
	} else {
		file, err := os.Open(path)
		if err != nil {
			return "", "", err
		}
		defer file.Close()
		imgByte, _ = io.ReadAll(file)
	}

	mimeType := http.DetectContentType(imgByte)
	base64String := base64.StdEncoding.EncodeToString(imgByte)
	return base64String, mimeType, nil
}
