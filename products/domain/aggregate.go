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
	switch cmd.(type) {
	case *Create:
		// An aggregate can only be created once.
		if a.created {
			return errors.New("already created")
		}
	default:
		// All other events require the aggregate to be created.
		if !a.created {
			return errors.New("not created")
		}
	}

	switch cmd := cmd.(type) {
	case *Create:
		a.StoreEvent(ProductCreated, nil, TimeNow())
	case *Delete:
		a.StoreEvent(ProductDeleted, nil, TimeNow())
	case *AddProductLang:
		a.StoreEvent(ProductLangAdded, &ProductLangAddedData{
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
		}, TimeNow())
	case *UpdateProductLang:
		a.StoreEvent(ProductLangUpdated, &ProductLangUpdatedData{
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
	case ProductCreated:
		a.created = true
		println("create product event")
	case ProductDeleted:
		a.created = false
		println("remove product event")
	case ProductLangAdded:
		data, ok := event.Data().(*ProductLangAddedData)
		if !ok {
			return errors.New("Invalid event data for ProductLangAdded")
		}

		for i, e := range a.productLangs {
			println("lancode " + string(i) + " " + e.LangCode)
			if len(a.productLangs) > 0 && e.LangCode != "" && e.LangCode == data.LangCode {
				return errors.New("ProductLang for langCode exist -> " + e.LangCode + "::" + data.LangCode + ":" + string(len(a.productLangs)))
			}
		}

		a.productLangs = append(a.productLangs, ProductLang{
			Name:             data.Name,
			Description:      data.Description,
			DescriptionShort: data.DescriptionShort,
			LinkRewrite:      data.LinkRewrite,
			MetaDescription:  data.MetaDescription,
			MetaKeywords:     data.MetaKeywords,
			MetaTitle:        data.MetaTitle,
			AvailableNow:     data.AvailableNow,
			AvailableLater:   data.AvailableLater,
			LangCode:         data.LangCode,
		})
		println("added productLang event")
	case ProductLangUpdated:
		println("updated productLang event")
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
		a.productLangs = append(a.productLangs, ProductLang{
			Name:             data.Name,
			Description:      data.Description,
			DescriptionShort: data.DescriptionShort,
			LinkRewrite:      data.LinkRewrite,
			MetaDescription:  data.MetaDescription,
			MetaKeywords:     data.MetaKeywords,
			MetaTitle:        data.MetaTitle,
			AvailableNow:     data.AvailableNow,
			AvailableLater:   data.AvailableLater,
			LangCode:         data.LangCode,
		})

	case ProductLangRemove:
		println("remove productLang event")
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
