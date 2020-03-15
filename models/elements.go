package models

import "sync"

type ElementAttributes struct {
}

type Element struct {
	Name        string
	Description string
	Image       string
	Attributes  []ElementAttributes
}

type ElementsModel struct {
	Elements map[string]Element
	lock     sync.RWMutex
}

var Elements = ElementsModel{
	Elements: map[string]Element{
		"0": {},
	},
}
