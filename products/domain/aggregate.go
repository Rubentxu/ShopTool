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
	nextItemID   int
	productLangs []ProductLang
}

// HandleCommand implements the HandleCommand method of the
// eventhorizon.CommandHandler interface.
func (a *AggregateProduct) HandleCommand(ctx context.Context, cmd eh.Command) error {
	switch cmd := cmd.(type) {
	case *AddProductLang:
		a.StoreEvent(ProductLangAdded, &ProductLangAddedData{
			ProductLang: &ProductLang{
				ID:               a.nextItemID,
				Name:             cmd.Name,
				Description:      cmd.Description,
				DescriptionShort: cmd.DescriptionShort,
				LinkRewrite:      cmd.LinkRewrite,
				MetaDescription:  cmd.MetaDescription,
				MetaKeywords:     cmd.MetaKeywords,
				MetaTitle:        cmd.MetaTitle,
				AvailableNow:     cmd.AvailableNow,
				AvailableLater:   cmd.AvailableLater,
				LangCode:         cmd.LangCode,
			},
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
		}
		a.productLangs = append(a.productLangs, ProductLang{
			ID:               data.ProductLang.ID,
			Name:             data.ProductLang.Name,
			Description:      data.ProductLang.Description,
			DescriptionShort: data.ProductLang.DescriptionShort,
			LinkRewrite:      data.ProductLang.LinkRewrite,
			MetaDescription:  data.ProductLang.MetaDescription,
			MetaKeywords:     data.ProductLang.MetaKeywords,
			MetaTitle:        data.ProductLang.MetaTitle,
			AvailableNow:     data.ProductLang.AvailableNow,
			AvailableLater:   data.ProductLang.AvailableLater,
			LangCode:         data.ProductLang.LangCode,
		})
		a.nextItemID++

	default:
		return fmt.Errorf("could not apply event: %s", event.EventType())
	}
	return nil
}
