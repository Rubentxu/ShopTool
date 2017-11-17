package domain

import (
	eh "github.com/looplab/eventhorizon"
)

// Events to products
const (
	ProductCreated            = eh.EventType("product:created")
	ProductDeleted            = eh.EventType("product:deleted")
	ProductLangAdded          = eh.EventType("product:productLangAdded")
	ProductLangUpdated        = eh.EventType("product:productLangUpdated")
	ProductLangRemoved        = eh.EventType("product:productLangRemoved")
	AvailabilitySet           = eh.EventType("product:availability")
	TransportSpecificationSet = eh.EventType("product:transportSpecification")
	TransportAdded            = eh.EventType("product:transportSpecification:transportAdded")
	TransportUpdated          = eh.EventType("product:transportSpecification:transportUpdated")
	TransportRemoved          = eh.EventType("product:transportSpecification:transportRemoved")
)

func init() {
	eh.RegisterEventData(ProductCreated, func() eh.EventData {
		return &CreateData{}
	})
	eh.RegisterEventData(ProductLangAdded, func() eh.EventData {
		return &ProductLangData{}
	})
	eh.RegisterEventData(ProductLangUpdated, func() eh.EventData {
		return &ProductLangData{}
	})
	eh.RegisterEventData(ProductLangRemoved, func() eh.EventData {
		return &ProductLangRemoveData{}
	})
	eh.RegisterEventData(AvailabilitySet, func() eh.EventData {
		return &AvailabilityData{}
	})

	eh.RegisterEventData(TransportSpecificationSet, func() eh.EventData {
		return &TransportSpecificationData{}
	})

	eh.RegisterEventData(TransportAdded, func() eh.EventData {
		return &TransporterData{}
	})

	eh.RegisterEventData(TransportUpdated, func() eh.EventData {
		return &TransporterData{}
	})

	eh.RegisterEventData(TransportRemoved, func() eh.EventData {
		return &TranporterRemovedData{}
	})
}

// CreateData is the event data for the Product
type CreateData struct {
	Reference string `json:"reference" bson:"reference"`
	Ean13     string `json:"ean_13" bson:"ean_13"`
	Isbn      string `json:"isbn" bson:"isbn"`
	Upc       string `json:"upc" bson:"upc"`
}

// ProductLangData is the event data for the ProductLangAdded
type ProductLangData struct {
	ProductLang
}



// ProductLangRemoveData is the event data for the ProductLangRemoved
type ProductLangRemoveData struct {
	LangCode string `json:"lang_code" bson:"lang_code"`
}

type AvailabilityData struct {
	Availability
}

type TransportSpecificationData struct {
	 TransportSpecification
}

type TransporterData struct {
	Transporter
}

type TranporterRemovedData struct {
	transportID eh.UUID `json:"transport_id" bson:"transport_id"`
}