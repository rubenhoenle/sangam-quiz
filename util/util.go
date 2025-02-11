package util

import (
	"github.com/rubenhoenle/sangam-quiz/model"
	"math/rand"
)

func GetRandomSangamItem(items []model.SangamItem) model.SangamItem {
	randomIndex := rand.Intn(len(items))
	return items[randomIndex]
}
