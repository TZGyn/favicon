package providers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"tzgyn/webatar/utils"
)

func Domain(domain string) (Avatar, error) {
	if !strings.HasPrefix(domain, "http://") && !strings.HasPrefix(domain, "https://") {
		domain = "https://" + domain
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", domain, nil)
	if err != nil {
		return Avatar{}, err
	}

	req.Header.Set("User-Agent", "Bot")

	resp, err := client.Do(req)
	if err != nil {
		return fallback(domain)
	}
	defer resp.Body.Close()

	imageSrc, err := utils.GetTagWithAttribute(resp.Body, "link[rel='icon']", "href")
	if err != nil {
		return fallback(domain)
	}

	if !strings.HasPrefix(imageSrc, "http://") && !strings.HasPrefix(imageSrc, "https://") {
		imageSrc = domain + imageSrc
	}

	image, err := utils.GetImage(imageSrc)

	if err != nil {
		return fallback(domain)
	}

	return Avatar{Data: image.Data, Data_type: image.Data_type}, nil

}

func fallback(domain string) (Avatar, error) {
	google := "https://www.google.com/s2/favicons?domain_url=%s&sz=128"

	url := fmt.Sprintf(google, domain)

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
