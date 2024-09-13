package utils

import (
	"io"
	"net/http"
)

type Image struct {
	Data      string
	Data_type string
}

func GetImage(url string) (Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Image{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Image{}, err
	}

	defer resp.Body.Close()

	result_type := resp.Header.Get("Content-Type")

	return Image{Data: string(data), Data_type: result_type}, nil
}
