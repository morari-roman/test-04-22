package asana

import (
	"encoding/json"
)

type AsanaEntity struct {
	Gid          string `json:"gid"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

type AsanaData struct {
	Data []AsanaEntity `json:"data"`
}

func New() *AsanaData {
	return &AsanaData{}
}

func (pr *AsanaData) DecodeData(bs []byte) error {
	err := json.Unmarshal(bs, &pr)

	return err
}
