package pkg

import (
	"encoding/json"
	"os"
)

func LoadPersonalGitmojiCollection(source string) (GitmojiCollection, error) {
	content, err := os.ReadFile(source)

	if err != nil {
		return GitmojiCollection{}, err
	}

	var collection GitmojiCollection

	err = json.Unmarshal(content, &collection)

	return collection, nil
}
