package service

import "github.com/rubenhoenle/sangam-quiz/model"

type SangamItemProvider interface {
	GetSangamItems() ([]model.SangamItem, error)
}
