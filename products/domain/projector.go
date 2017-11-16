package domain

import (
	"context"
	"errors"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/eventhandler/projector"

)

// ProductProjector is a projector of product events on the Product read model.
type ProductProjector struct{}

// ProjectorType implements the ProjectorType method of the
// eventhorizon.Projector interface.
func (p *ProductProjector) ProjectorType() projector.Type {
	return projector.Type(string(AggregateProductType) + "_projector")
}

// Project implements the Project method of the eventhorizon.Projector interface.
func (p *ProductProjector) Project(ctx context.Context, event eh.Event, entity eh.Entity) (eh.Entity, error) {
	println("Projector project")
	model, ok := entity.(*Product)
	if !ok {
		return nil, errors.New("model is of incorrect type")
	}

	switch event.EventType() {
	case ProductCreated:
		println("Projector ProductCreated")
		data, ok := event.Data().(*CreateData)
		if !ok {
			return nil, errors.New("invalid event data ProductLangAdded")
		}

		// Set the ID when first created.
		model.ID = event.AggregateID()
		model.ProductLangs = []*ProductLang{} // Prevents "null" in JSON.
		model.Reference = data.Reference
		model.Upc = data.Upc
		model.Isbn = data.Isbn
		model.Ean13 = data.Ean13
		model.CreatedAt = TimeNow()
	case ProductDeleted:
		// Return nil as the entity to delete the model.
		return nil, nil
	case ProductLangAdded:
		println("Projector ProductLangAdded")
		data, ok := event.Data().(*ProductLangAddedData)
		if !ok {
			return nil, errors.New("invalid event data ProductLangAdded")
		}
		productLang := &ProductLang{}
		*productLang = *&data.ProductLang
		model.ProductLangs = append(model.ProductLangs, productLang)
	case ProductLangUpdated:
		println("Projector ProductLangAdded")
		data, ok := event.Data().(*ProductLangUpdatedData)
		if !ok {
			return nil, errors.New("invalid event data ProductLangAdded")
		}
		productLang := &ProductLang{}
		*productLang = *&data.ProductLang
		model.ProductLangs = append(model.ProductLangs, productLang)
	case ProductLangRemove:
		println("Projector ProductLangAdded")
		data, ok := event.Data().(*ProductLangRemoveData)
		if !ok {
			return nil, errors.New("invalid event data ProductLangAdded")
		}
		atemp := model.ProductLangs
		for i, e := range atemp {
			if e.LangCode == data.LangCode {
				model.ProductLangs = atemp[:i+copy(atemp[i:], atemp[i+1:])]
			}
		}
	case AvailabilitySet:
		data, ok := event.Data().(*AvailabilityData)
		if !ok {
			return nil, errors.New("Invalid event data for ProductLangUpdated")
		}
		model.Availability = data.Availability
	}
	model.Version++
	model.UpdatedAt = TimeNow()
	return model, nil
}

