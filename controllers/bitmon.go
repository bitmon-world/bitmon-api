package controllers

import (
	"encoding/json"
	"errors"
	"github.com/bitmon-world/bitmon-api/models"
	"github.com/bitmon-world/bitmon-api/types"
)

type TestRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BitmonController struct {
	dbModel *models.BitmonDBModel
}

func (ctrl *BitmonController) GetMonList(params types.ReqParams) (interface{}, error) {
	data, err := ctrl.dbModel.GetGenMon(params.ID)
	if err != nil {
		return nil, errors.New("general monster information not found")
	}
	return data, nil
}

func (ctrl *BitmonController) GetMonInfo(params types.ReqParams) (interface{}, error) {
	data, err := ctrl.dbModel.GetGenMon(params.ID)
	if err != nil {
		return nil, errors.New("general monster information not found")
	}
	return data, nil
}

func (ctrl *BitmonController) AddMon(params types.ReqParams) (interface{}, error) {
	data, err := ctrl.dbModel.GetGenMon(params.ID)
	if err != nil {
		return nil, errors.New("general monster information not found")
	}
	return data, nil
}

func (ctrl *BitmonController) GetElementsList(params types.ReqParams) (interface{}, error) {
	data, err := ctrl.dbModel.GetElementsList()
	if err != nil {
		return nil, errors.New("unable to get items list")
	}
	return data, nil
}

func (ctrl *BitmonController) GetElement(params types.ReqParams) (interface{}, error) {
	data, err := ctrl.dbModel.GetElement(params.ID)
	if err != nil {
		return nil, errors.New("unable to item")
	}
	return data, nil
}

func (ctrl *BitmonController) GetImage(params types.ReqParams) (interface{}, error) {
	return types.Success{Success: true}, nil
}

func (ctrl *BitmonController) AddElement(params types.ReqParams) (interface{}, error) {
	var element types.Elements
	err := json.Unmarshal(params.Body, &element)
	if err != nil {
		return nil, err
	}
	err = ctrl.dbModel.AddElement(element)
	if err != nil {
		return nil, errors.New("unable to store element information")
	}
	return types.Success{Success: true}, nil
}

func (ctrl *BitmonController) CalcAdventure(params types.ReqParams) (interface{}, error) {
	return nil, nil
}

func NewBitmonController(dbUri string, dbName string) *BitmonController {
	model := models.NewDBModel(dbUri, dbName)
	return &BitmonController{dbModel: model}
}
