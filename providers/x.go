package providers

import (
	"net/http"
	"tzgyn/webatar/utils"
)

func XAccount(account string) (Avatar, error) {
	twitter := "https://www.x.com/"

	url := twitter + account

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Avatar{}, err
	}

	req.Header.Set("User-Agent", "Bot")

	resp, err := client.Do(req)
	if err != nil {
		return Avatar{}, err
	}
	defer resp.Body.Close()

	OGImage, err := utils.GetOGImage(resp.Body)
	if err != nil {
		return Avatar{}, err
	}

	image, err := utils.GetImage(OGImage)

	if err != nil {
		return Avatar{}, err
	}

	return Avatar{Data: image.Data, Data_type: image.Data_type}, nil
}
