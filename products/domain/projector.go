package domain

import (
	"context"
	"errors"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/eventhandler/projector"
)

type ProductProjector struct{}

func (p *ProductProjector) ProjectorType() projector.Type {
	return projector.Type(string(AggregateProductType) + "_projector")
}

func (p *ProductProjector) Project(ctx context.Context, event eh.Event, entity eh.Entity) (eh.Entity, error) {
	model, ok := entity.(*Product)
	if !ok {
		return nil, errors.New("model is of incorrect typel7l")
	}

}
