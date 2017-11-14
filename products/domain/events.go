package domain

import (
	eh "github.com/looplab/eventhorizon"
)

// Events to products
const (
	ProductCreated     = eh.EventType("product:created")
	ProductDeleted     = eh.EventType("product:deleted")
	LangAdded   = eh.EventType("product:langAdded")
	LangUpdated = eh.EventType("product:langUpdated")
	LangRemove  = eh.EventType("product:langRemove")
	AvailabilitySet    = eh.EventType("product:availability")
)

func init() {
	eh.RegisterEventData(ProductCreated, func() eh.EventData {
		return &CreateData{}
	})
	eh.RegisterEventData(LangAdded, func() eh.EventData {
		return &LangAddedData{}
	})
	eh.RegisterEventData(LangUpdated, func() eh.EventData {
		return &LangUpdatedData{}
	})
	eh.RegisterEventData(LangRemove, func() eh.EventData {
		return &LangRemoveData{}
	})
	eh.RegisterEventData(AvailabilitySet, func() eh.EventData {
		return &AvailabilityData{}
	})

}

// CreateData is the event data for the Product
type CreateData struct {
	Reference string `json:"reference" bson:"reference"`
	Ean13     string `json:"ean_13" bson:"ean_13"`
	Isbn      string `json:"isbn" bson:"isbn"`
	Upc       string `json:"upc" bson:"upc"`
}

// LangAddedData is the event data for the LangAdded
type LangAddedData struct {
	ProductLang
}

// LangUpdatedData is the event data for the ProductLangUpdate
type LangUpdatedData struct {
	ProductLang
}

// LangRemoveData is the event data for the LangRemove
type LangRemoveData struct {
	LangCode string `json:"lang_code" bson:"lang_code"`
}

type AvailabilityData struct {
	Availability
}
