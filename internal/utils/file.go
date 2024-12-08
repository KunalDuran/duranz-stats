package utils

import (
	"os"
	"strings"
)

func ListJSONFiles(path string) []string {
	var fileList []string

	files, err := os.ReadDir(path)
	if err != nil {
		return fileList
	}

	for _, f := range files {
		if strings.Contains(f.Name(), "json") {
			fileList = append(fileList, f.Name())
		}
	}
	return fileList
}

func ReadMatchJSON(f_path string) ([]byte, error) {
	body, err := os.ReadFile(f_path)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
