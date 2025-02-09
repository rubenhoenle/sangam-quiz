package model

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
)

// assert bubbletea listitem interface is implemented
var _ list.Item = (*SangamItem)(nil)

type SangamItem struct {
	Id    string `json:"short_name"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (i SangamItem) Title() string { return i.Name }
func (i SangamItem) Description() string {
	euros := i.Price / 100
	centsRemaining := i.Price % 100
	return fmt.Sprintf("%d,%02d€", euros, centsRemaining)
}
func (i SangamItem) FilterValue() string { return i.Name }
