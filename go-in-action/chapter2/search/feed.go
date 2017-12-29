package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
