package domain

import (
	"context"
	"errors"
	"fmt"
	"time"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
)

// TimeNow is a mockable version of time.Now.
var TimeNow = time.Now
var idGen = eh.NewUUID

func init() {
	eh.RegisterAggregate(func(id eh.UUID) eh.Aggregate {
		return &AggregateProduct{
			AggregateBase: events.NewAggregateBase(AggregateProductType, id),
		}
	})
}

// AggregateProductType is the aggregate type for the product
const AggregateProductType = eh.AggregateType("product")

// AggregateProduct  is an aggregate for a product.
type AggregateProduct struct {
	*events.AggregateBase
	availability           Availability
	created                bool
	productLangs           []ProductLang
	transportSpecification TransportSpecification
}

// HandleCommand implements the HandleCommand method of the
// eventhorizon.CommandHandler interface.
func (a *AggregateProduct) HandleCommand(ctx context.Context, cmd eh.Command) error {
	switch cmd.(type) {
	case *Create:
		// An aggregate can only be created once.
		if a.created {
			return errors.New("already created")
		}
	default:
		// All other events require the aggregate to be created.
		if !a.created {
			return errors.New("product not exist")
		}
	}

	switch cmd := cmd.(type) {
	case *Create:
		a.StoreEvent(ProductCreated, &CreateData{
			Reference: cmd.Reference,
			Ean13:     cmd.Ean13,
			Isbn:      cmd.Isbn,
			Upc:       cmd.Upc,
		}, TimeNow())
	case *Delete:
		a.StoreEvent(ProductDeleted, nil, TimeNow())
	case *AddProductLang:
		a.StoreEvent(ProductLangAdded, &ProductLangData{
			ProductLang: cmd.ProductLang,
		}, TimeNow())
	case *UpdateProductLang:
		a.StoreEvent(ProductLangUpdated, &ProductLangData{
			ProductLang: cmd.ProductLang,
		}, TimeNow())
	case *RemoveProductLang:
		a.StoreEvent(ProductLangRemoved, &ProductLangRemoveData{
			LangCode: cmd.LangCode,
		}, TimeNow())
	case *SetAvailability:
		a.StoreEvent(AvailabilitySet, &AvailabilityData{
			Availability: cmd.Availability,
		}, TimeNow())
	case *SetTransportSpecification:
		a.StoreEvent(TransportSpecificationSet, &TransportSpecificationData{
			TransportSpecification: cmd.TransportSpecification,
		}, TimeNow())
	case *AddTransport:
		a.StoreEvent(TransportAdded, &TransporterData{
			Transporter: cmd.Transporter,
		}, TimeNow())
	case *UpdateTransport:
		a.StoreEvent(TransportUpdated, &TransporterData{
			Transporter: cmd.Transporter,
		}, TimeNow())
	case *RemoveTransport:
		a.StoreEvent(TransportRemoved, &TranporterRemovedData{
			transportID: cmd.transportID,
		}, TimeNow())
	default:
		return fmt.Errorf("could not handle command: %s", cmd.CommandType())
	}
	return nil
}

// ApplyEvent implements the ApplyEvent method of the
// eventhorizon.Aggregate interface.
func (a *AggregateProduct) ApplyEvent(ctx context.Context, event eh.Event) error {
	switch event.EventType() {
	case ProductCreated:
		a.created = true

	case ProductDeleted:
		a.created = false

	case ProductLangAdded:
		data, ok := event.Data().(*ProductLangData)
		if !ok {
			return fmt.Errorf("Invalid event %s for productLang %s", event.EventType(), data.LangCode)
		}

		for _, e := range a.productLangs {
			if len(a.productLangs) > 0 && e.LangCode != "" && e.LangCode == data.LangCode {
				return fmt.Errorf("ProductLang for langCode %s exist. ", data.LangCode)
			}
		}
		a.productLangs = append(a.productLangs, data.ProductLang)

	case ProductLangUpdated:
		data, ok := event.Data().(*ProductLangData)
		if !ok {
			return fmt.Errorf("Invalid event %s for productLang %s", event.EventType(), data.LangCode)
		}

		existProductLang := false
		if a.productLangs == nil {
			return fmt.Errorf("Error Event %s , langCode %s not exist", event.EventType(), data.LangCode)
		}

		for _, e := range a.productLangs {
			if e.LangCode == data.LangCode {
				existProductLang = true
			}
		}

		if !existProductLang {
			return fmt.Errorf("Error Event %s , langCode %s not exist", event.EventType(), data.LangCode)
		}
		a.productLangs = append(a.productLangs, data.ProductLang)

	case ProductLangRemoved:
		data, ok := event.Data().(*ProductLangRemoveData)
		if !ok {
			return fmt.Errorf("Invalid event %s for productLang %s", event.EventType(), data.LangCode)
		}
		if a.productLangs == nil {
			return fmt.Errorf("Error Event %s , langCode %s not exist", event.EventType(), data.LangCode)
		}

		removedProductLang := false
		atemp := a.productLangs
		for i, e := range atemp {
			if e.LangCode == data.LangCode {
				a.productLangs = atemp[:i+copy(atemp[i:], atemp[i+1:])]
				removedProductLang = true
			}
		}
		if !removedProductLang {
			return fmt.Errorf("Error Event %s , langCode %s not exist", event.EventType(), data.LangCode)
		}
	case AvailabilitySet:
		data, ok := event.Data().(*AvailabilityData)
		if !ok {
			return fmt.Errorf("Invalid event %s", event.EventType())
		}
		a.availability = data.Availability
	case TransportSpecificationSet:
		data, ok := event.Data().(*TransportSpecificationData)
		if !ok {
			return fmt.Errorf("Invalid event %s", event.EventType())
		}
		a.transportSpecification = data.TransportSpecification
	case TransportAdded:
		data, ok := event.Data().(*TransporterData)
		if !ok {
			return fmt.Errorf("Invalid event %s for transporter %s", event.EventType(), data.Id)
		}
		for _, e := range a.transportSpecification.Transporters {

			if len(a.transportSpecification.Transporters) > 0 && e.Id != "" && e.Id == data.Id {
				return fmt.Errorf("transport %s for aggretate  %s exist -> ", e.Name, a.EntityID())

			}
		}
		a.transportSpecification.Transporters = append(a.transportSpecification.Transporters, data.Transporter)
	case TransportUpdated:
		data, ok := event.Data().(*TransporterData)
		if !ok {
			return fmt.Errorf("Invalid event %s for transporter %s", event.EventType(), data.Id)
		}

		existProductLang := false
		if a.transportSpecification.Transporters == nil {
			return fmt.Errorf("Transporter for %s not exist", data.Id)
		}

		for _, e := range a.transportSpecification.Transporters {
			if e.Id == data.Id {
				existProductLang = true
			}
		}

		if !existProductLang {
			return fmt.Errorf("Transporter for ID %s not exist", data.Id)
		}
		a.transportSpecification.Transporters = append(a.transportSpecification.Transporters, data.Transporter)
	case TransportRemoved:
		data, ok := event.Data().(*TranporterRemovedData)
		if !ok {
			return fmt.Errorf("Invalid event %s for transporter %s", event.EventType(), data.transportID)
		}
		if a.transportSpecification.Transporters == nil {
			return fmt.Errorf("Invalid event %s for transporter %s", event.EventType(), data.transportID)
		}

		removedProductLang := false
		atemp := a.transportSpecification.Transporters
		for i, e := range atemp {
			if e.Id == data.transportID {
				a.transportSpecification.Transporters = atemp[:i+copy(atemp[i:], atemp[i+1:])]
				removedProductLang = true
			}
		}
		if !removedProductLang {
			return fmt.Errorf("Transporter for ID %s not exist", data.transportID)
		}
	default:
		return fmt.Errorf("Could not apply event: %s", event.EventType())
	}
	return nil
}
