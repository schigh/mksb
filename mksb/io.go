package mksb

import (
	"bytes"
	"io/ioutil"
)

func GetLinesInFile(filePath string, delimiter int) ([][]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	data = bytes.TrimSpace(data)
	return bytes.Split(data, []byte{byte(delimiter)}), nil
}
