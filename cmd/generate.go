package cmd

import (
	"encoding/json"
	"github.com/vinicius73/personal-gitmoji/pkg"
	"github.com/vinicius73/personal-gitmoji/support"
)

func GenerateCollection(source string) ([]byte, error) {
	source, err := support.AbsPath(source)

	if err != nil {
		return nil, err
	}

	collection, err := pkg.LoadGitmojiCollection(source)

	if err != nil {
		return nil, err
	}

	j, err := json.MarshalIndent(collection, "", "  ")

	if err != nil {
		return nil, err
	}

	return j, err
}
