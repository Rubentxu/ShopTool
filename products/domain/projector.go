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
	model, ok := entity.(*Product)
	if !ok {
		return nil, errors.New("model is of incorrect typel7l")
	}

	switch event.EventType() {
	case ProductLangAdded:
		data, ok := event.Data().(*ProductLangAddedData)
		if !ok {
			return nil, errors.New("invalid event data ProductLangAdded")
		}
		model.ID = event.AggregateID()
		if model.ProductLangs == nil {
			model.ProductLangs = []*ProductLang{}
		}

		model.ProductLangs = append(model.ProductLangs, &data.ProductLang)
	}
	model.Version++
	model.UpdateAt = TimeNow()
	return model, nil
}
