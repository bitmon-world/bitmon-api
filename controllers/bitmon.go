package controllers

import (
	"github.com/bitmon-world/bitmon-api/models"
)

type TestRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BitmonController struct {
	Elements models.ElementsModel
	Bitmons  models.BitmonModel
}
