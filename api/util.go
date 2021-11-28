package api

import (
	"encoding/base64"
)

func b64decode(str string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	return string(b), nil
}


