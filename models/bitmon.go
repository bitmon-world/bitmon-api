package models

import (
	"errors"
	"sync"
)

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

func (m *BitmonModel) Get(id string) (Bitmon, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	e, ok := m.Bitmons[id]
	if !ok {
		return Bitmon{}, errors.New("bitmons doesn't exist")
	}
	return e, nil
}
var Bitmons = BitmonModel{
	Bitmons: map[string]Bitmon{
		"0": {},
	},
}

