package support

import (
	"os"
	"path"
)

func AbsPath(name string) (string, error) {
	if path.IsAbs(name) {
		return name, nil
	}

	root, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return path.Join(root, name), nil
}

