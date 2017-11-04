package domain

import (
	eh "github.com/looplab/eventhorizon"
)

// Events to products
const (
	ProductLangAdded   = eh.EventType("product:productlang_added")
	ProductLangUpdated = eh.EventType("product:productlang_updated")
	ProductLangRemove  = eh.EventType("product:productlang_remove")
)

func init() {
	eh.RegisterEventData(ProductLangAdded, func() eh.EventData {
		return &ProductLangAddedData{}
	})
	eh.RegisterEventData(ProductLangUpdated, func() eh.EventData {
		return &ProductLangUpdatedData{}
	})
	eh.RegisterEventData(ProductLangRemove, func() eh.EventData {
		return &ProductLangRemoveData{}
	})
}

// ProductLangAddedData is the event data for the ProductLangAdded
type ProductLangAddedData struct {
	ProductLang
}

// ProductLangUpdatedData is the event data for the ProductLangUpdate
type ProductLangUpdatedData struct {
	*ProductLang
}

// ProductLangRemoveData is the event data for the ProductLangRemove
type ProductLangRemoveData struct {
	LangCode string `json:"lang_code" bson:"lang_code"`
}
