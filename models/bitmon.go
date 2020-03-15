package models

import "sync"

type BitmonAttributes struct {
}

type Bitmon struct {
	Name        string
	Description string
	Image       string
	Attributes  []BitmonAttributes
}

type BitmonModel struct {
	Bitmons map[string]Bitmon
	lock    sync.RWMutex
}

var Bitmons = BitmonModel{
	Bitmons: map[string]Bitmon{
		"0": {},
	},
}
