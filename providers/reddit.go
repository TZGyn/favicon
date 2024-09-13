package providers

import (
	"net/http"
	"tzgyn/favicon/utils"
)

func RedditAccount(account string) (Avatar, error) {
	reddit := "https://www.reddit.com/user/"

	url := reddit + account

	resp, err := http.Get(url)
	if err != nil {
		return Avatar{}, err
	}
	defer resp.Body.Close()

	imageSrc, err := utils.GetTagWithAttribute(resp.Body, "img[alt*='avatar']", "src")
	if err != nil {
		return Avatar{}, err
	}

	image, err := utils.GetImage(imageSrc)

	if err != nil {
		return Avatar{}, err
	}

	return Avatar{Data: image.Data, Data_type: image.Data_type}, nil
}
