package domain

import (
	eh "github.com/looplab/eventhorizon"
)

// Events to products
const (
	ProductLangAdded = eh.EventType("product:productlang_added")
)

func init() {
	eh.RegisterEventData(ProductLangAdded, func() eh.EventData {
		return &ProductLangAddedData{}
	})
}

// ProductLangAddedData is the event data for the ProductLangAdded
type ProductLangAddedData struct {
	*ProductLang
}
