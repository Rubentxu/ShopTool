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
	created      bool
	productLangs []ProductLang
}

// HandleCommand implements the HandleCommand method of the
// eventhorizon.CommandHandler interface.
func (a *AggregateProduct) HandleCommand(ctx context.Context, cmd eh.Command) error {
	switch cmd := cmd.(type) {
	case *AddProductLang:
		productLangData := ProductLang{}
		productLangData = cmd.ProductLang
		a.StoreEvent(ProductLangAdded, &ProductLangAddedData{
			ProductLang: productLangData,
		}, TimeNow())
	case *UpdateProductLang:
		productLangData := &ProductLang{}
		*productLangData = *cmd.ProductLang
		a.StoreEvent(ProductLangUpdated, &ProductLangUpdatedData{
			ProductLang: productLangData,
		}, TimeNow())
	case *RemoveProductLang:
		a.StoreEvent(ProductLangRemove, &ProductLangRemoveData{
			LangCode: cmd.LangCode,
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
	case ProductLangAdded:
		data, ok := event.Data().(*ProductLangAddedData)
		if !ok {
			return errors.New("Invalid event data for ProductLangAdded")
		}
		if a.productLangs == nil {
			a.productLangs = []ProductLang{}
		} else {
			for _, e := range a.productLangs {
				if e.LangCode == data.LangCode {
					return errors.New("ProductLang for langCode exist")
				}
			}
		}
		productLangData := ProductLang{}
		productLangData = data.ProductLang
		a.productLangs = append(a.productLangs, productLangData)

	case ProductLangUpdated:
		data, ok := event.Data().(*ProductLangUpdatedData)
		if !ok {
			return errors.New("Invalid event data for ProductLangUpdated")
		}

		existProductLang := false
		if a.productLangs == nil {
			return errors.New("ProductLang for langCode not exist")
		}

		for _, e := range a.productLangs {
			if e.LangCode == data.LangCode {
				existProductLang = true
			}
		}

		if !existProductLang {
			return errors.New("ProductLang for langCode not exist")
		}
		productLangData := ProductLang{}
		productLangData = *data.ProductLang
		a.productLangs = append(a.productLangs, productLangData)

	case ProductLangRemove:
		data, ok := event.Data().(*ProductLangRemoveData)
		if !ok {
			return errors.New("Invalid event data for ProductLangUpdated")
		}
		if a.productLangs == nil {
			return errors.New("ProductLang for langCode not exist")
		}

		removedProductLang := false
		atemp := a.productLangs
		for i, e := range atemp {
			if e.LangCode == data.LangCode {
				a.productLangs = atemp[:i+copy(atemp[i:], atemp[i+1:])]
				removedProductLang = true
			}
		}
		if removedProductLang {
			return errors.New("ProductLang for langCode not exist")
		}

	default:
		return fmt.Errorf("could not apply event: %s", event.EventType())
	}
	return nil
}
