package jsonprovider

import (
	_ "embed"
	"encoding/json"
	"github.com/rubenhoenle/sangam-quiz/model"
	"github.com/rubenhoenle/sangam-quiz/service"
)

//go:embed sangam.json
var jsonData []byte

type jsonFileContent struct {
	Name  string             `json:"name"`
	Items []model.SangamItem `json:"items"`
}

// assert the SangamItemProvider interface is implemented
var _ service.SangamItemProvider = (*JsonSangamItemProvider)(nil)

type JsonSangamItemProvider struct{}

func (p JsonSangamItemProvider) GetSangamItems() ([]model.SangamItem, error) {
	var content jsonFileContent
	err := json.Unmarshal(jsonData, &content)
	if err != nil {
		return []model.SangamItem{}, err
	}
	return content.Items, nil
}
