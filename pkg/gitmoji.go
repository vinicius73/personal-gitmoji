package pkg

import (
	"github.com/go-resty/resty/v2"
	_ "github.com/go-resty/resty/v2"
)

const GITMOJI_API = "https://gitmoji.dev/api/gitmojis"

func LoadDefaultGitmojiCollection() (GitmojiCollection, error) {
	var bodyResponse GitmojiCollection

	_, err := resty.New().R().
		SetResult(&bodyResponse).
		Get(GITMOJI_API)

	if err != nil {
		return GitmojiCollection{}, err
	}

	return bodyResponse, nil
}
