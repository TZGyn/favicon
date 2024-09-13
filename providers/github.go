package providers

import (
	"fmt"
	"tzgyn/favicon/utils"
)

func GithubAccount(account string) (Avatar, error) {
	github := "https://github.com/%s.png"

	url := fmt.Sprintf(github, account)

	image, err := utils.GetImage(url)

	if err != nil {
		return Avatar{}, err
	}

	return Avatar{Data: image.Data, Data_type: image.Data_type}, nil
}
