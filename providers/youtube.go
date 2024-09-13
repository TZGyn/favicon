package providers

import (
	"net/http"
	"tzgyn/favicon/utils"
)

func YoutubeChannel(channel string) (Avatar, error) {
	youtube := "https://www.youtube.com/@"

	url := youtube + channel

	resp, err := http.Get(url)
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
