package models

import (
	"errors"
	"sync"
)

var baseURL = "https://api.bitmon.io/img/elements/"

type ElementAttributes struct {
}

type Element struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Image       string              `json:"image"`
	Attributes  []ElementAttributes `json:"attributes"`
}

type ElementsModel struct {
	Elements map[string]Element
	lock     sync.RWMutex
}

func (m *ElementsModel) Get(id string) (Element, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	e, ok := m.Elements[id]
	if !ok {
		return Element{}, errors.New("element doesn't exist")
	}
	return e, nil
}

var Elements = ElementsModel{
	Elements: map[string]Element{
		"0": {
			Name:        "Free Ticket (0 USD)",
			Description: "The Free Ticket can be used to enter a free adventure to capture common bitmons",
			Image:       baseURL + "free_ticket.png",
		},
		"1": {
			Name:        "Basic Ticket (1 USD)",
			Description: "The Basic Ticket can be used to enter a basic adventure to capture common and rare bitmons",
			Image:       baseURL + "basic_ticket.png",
		},
		"2": {
			Name:        "Premium Ticket (10 USD)",
			Description: "The Premium Ticket can be used to enter a premium adventure to capture stronger common and rare bitmons",
			Image:       baseURL + "premium_ticket.png",
		},
		"3": {
			Name:        "Golden Ticket (50 USD)",
			Description: "The Golden Ticket can be used to enter a golden adventure to capture rare and legendary bitmons",
			Image:       baseURL + "golden_ticket.png",
		},
		"4": {
			Name:        "Ultimate Ticket (100 USD)",
			Description: "The Ultimate Ticket can be used to enter a ultimate adventure to capture strong rare and legendary bitmons",
			Image:       baseURL + "ultimate_ticket.png",
		},
		"5": {
			Name:        "Epic Ticket (1000 USD)",
			Description: "The Epic Ticket can be used to enter a epic adventure to capture strong legendary bitmons",
			Image:       baseURL + "epic_ticket.png",
		},
	},
}
