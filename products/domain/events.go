package domain

import (
	eh "github.com/looplab/eventhorizon"
)

// Events to products
const (
	ProductCreated     = eh.EventType("product:created")
	ProductDeleted     = eh.EventType("product:deleted")
	ProductLangAdded   = eh.EventType("product:productLangAdded")
	ProductLangUpdated = eh.EventType("product:productLangUpdated")
	ProductLangRemove  = eh.EventType("product:productLangRemove")
	AvailabilitySet    = eh.EventType("product:availability")
)

func init() {
	eh.RegisterEventData(ProductCreated, func() eh.EventData {
		return &CreateData{}
	})
	eh.RegisterEventData(ProductLangAdded, func() eh.EventData {
		return &ProductLangAddedData{}
	})
	eh.RegisterEventData(ProductLangUpdated, func() eh.EventData {
		return &ProductLangUpdatedData{}
	})
	eh.RegisterEventData(ProductLangRemove, func() eh.EventData {
		return &ProductLangRemoveData{}
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

// ProductLangAddedData is the event data for the ProductLangAdded
type ProductLangAddedData struct {
	ProductLang
}

// ProductLangUpdatedData is the event data for the ProductProductLangUpdate
type ProductLangUpdatedData struct {
	ProductLang
}

// ProductLangRemoveData is the event data for the ProductLangRemove
type ProductLangRemoveData struct {
	LangCode string `json:"lang_code" bson:"lang_code"`
}

type AvailabilityData struct {
	Availability
}
