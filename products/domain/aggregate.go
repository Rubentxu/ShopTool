package domain

import (
	"context"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
)

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
	productLangs []ProductLang
}

// HandleCommand implements the HandleCommand method of the
// eventhorizon.CommandHandler interface.
func (a *AggregateProduct) HandleCommand(ctx context.Context, cmd eh.Command) error {
	return nil
}

// ApplyEvent implements the ApplyEvent method of the
// eventhorizon.Aggregate interface.
func (a *AggregateProduct) ApplyEvent(ctx context.Context, event eh.Event) error {
	return nil
}
