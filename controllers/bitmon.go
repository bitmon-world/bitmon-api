package controllers

import (
	"github.com/bitmon-world/bitmon-api/models"
	"github.com/bitmon-world/bitmon-api/types"
)

type TestRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BitmonController struct {
	Elements models.ElementsModel
	Bitmons  models.BitmonModel
}


func (c *BitmonController) GetMon(params types.ReqParams) (interface{}, error) {
	return nil, nil
}

func (c *BitmonController) GetElement(params types.ReqParams) (interface{}, error) {
	return c.Elements.Get(params.ID)
}