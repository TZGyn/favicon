package providers

import (
	"fmt"
	"io"
	"net/http"
)

func Domain(domain string) (Avatar, error) {
	duckduckgo := "https://icons.duckduckgo.com/ip3/%s.ico"

	url := fmt.Sprintf(duckduckgo, domain)

	resp, err := http.Get(url)
	if err != nil {
		return Avatar{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Avatar{}, err
	}

	result_type := resp.Header.Get("Content-Type")

	return Avatar{Data: string(data), Data_type: result_type}, nil
}
