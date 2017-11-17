package domain

import (
	"context"
	"reflect"
	"testing"
	"time"

	"errors"
	"github.com/kr/pretty"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
)

func TestAggregateHandleCommand(t *testing.T) {
	TimeNow = func() time.Time {
		return time.Date(2017, time.November, 20, 42, 0, 0, 0, time.Local)
	}
	idGen = func() eh.UUID {
		return "aaabbbbcccccdddd1234"
	}

	id := idGen()
	cases := map[string]struct {
		agg            *AggregateProduct
		cmd            eh.Command
		expectedEvents []eh.Event
		expectedErr    error
	}{
		"create": {
			&AggregateProduct{
				AggregateBase: events.NewAggregateBase(AggregateProductType, id),
			},
			&Create{
				Reference: "abcd1234",
				Ean13:     "123456789",
				Isbn:      "",
				Upc:       "",
			},
			[]eh.Event{
				eh.NewEventForAggregate(ProductCreated, &CreateData{
					Reference: "abcd1234",
					Ean13:     "123456789",
					Isbn:      "",
					Upc:       "",
				}, TimeNow(), AggregateProductType, id, 1),
			},
			nil,
		},
		"createError": {
			&AggregateProduct{
				AggregateBase: events.NewAggregateBase(AggregateProductType, id),
				created:       true,
			},
			&Create{
				Reference: "abcd1234",
				Ean13:     "123456789",
				Isbn:      "",
				Upc:       "",
			},
			nil,
			errors.New("already created"),
		},
		"delete": {
			&AggregateProduct{
				AggregateBase: events.NewAggregateBase(AggregateProductType, id),
				created:       true,
			},
			&Delete{
				ID: id,
			},
			[]eh.Event{
				eh.NewEventForAggregate(ProductDeleted, nil, TimeNow(), AggregateProductType, id, 1),
			},
			nil,
		},
		"deleteNotCreated": {
			&AggregateProduct{
				AggregateBase: events.NewAggregateBase(AggregateProductType, id),
				created:       false,
			},
			&Delete{
				ID: id,
			},
			nil,
			errors.New("product not exist"),
		},
		"addProductLang": {
			&AggregateProduct{
				AggregateBase: events.NewAggregateBase(AggregateProductType, id),
				created:       true,
			},
			&AddProductLang{
				ProductID: id,
				ProductLang: ProductLang{
					Name:             "testName",
					Description:      "testDescription",
					DescriptionShort: "testDescriptionShort",
					LinkRewrite:      "testLinkRewrite",
					MetaDescription:  "testMetaDescription",
					MetaKeywords:     "testMetaKeywords",
					MetaTitle:        "testMetaTitle",
					AvailableNow:     "testAvailableNow",
					AvailableLater:   "testAvailableLater",
					LangCode:         "testLangCode",
				},
			},
			[]eh.Event{
				eh.NewEventForAggregate(ProductLangAdded, &ProductLangData{
					ProductLang: ProductLang{
						Name:             "testName",
						Description:      "testDescription",
						DescriptionShort: "testDescriptionShort",
						LinkRewrite:      "testLinkRewrite",
						MetaDescription:  "testMetaDescription",
						MetaKeywords:     "testMetaKeywords",
						MetaTitle:        "testMetaTitle",
						AvailableNow:     "testAvailableNow",
						AvailableLater:   "testAvailableLater",
						LangCode:         "testLangCode",
					},
				}, TimeNow(), AggregateProductType, id, 1),
			},
			nil,
		},
		"updateProductLang": {
			&AggregateProduct{
				AggregateBase: events.NewAggregateBase(AggregateProductType, id),
				created:       true,
			},
			&UpdateProductLang{
				ProductID: id,
				ProductLang: ProductLang{
					Name:             "testName",
					Description:      "testDescription",
					DescriptionShort: "testDescriptionShort",
					LinkRewrite:      "testLinkRewrite",
					MetaDescription:  "testMetaDescription",
					MetaKeywords:     "testMetaKeywords",
					MetaTitle:        "testMetaTitle",
					AvailableNow:     "testAvailableNow",
					AvailableLater:   "testAvailableLater",
					LangCode:         "testLangCode",
				},
			},
			[]eh.Event{
				eh.NewEventForAggregate(ProductLangUpdated, &ProductLangUpdatedData{
					ProductLang: ProductLang{
						Name:             "testName",
						Description:      "testDescription",
						DescriptionShort: "testDescriptionShort",
						LinkRewrite:      "testLinkRewrite",
						MetaDescription:  "testMetaDescription",
						MetaKeywords:     "testMetaKeywords",
						MetaTitle:        "testMetaTitle",
						AvailableNow:     "testAvailableNow",
						AvailableLater:   "testAvailableLater",
						LangCode:         "testLangCode",
					},
				}, TimeNow(), AggregateProductType, id, 1),
			},
			nil,
		},
		"removeProductLang": {
			&AggregateProduct{
				AggregateBase: events.NewAggregateBase(AggregateProductType, id),
				created:       true,
			},
			&RemoveProductLang{
				LangCode:  "Es_es",
				ProductID: id,
			},
			[]eh.Event{
				eh.NewEventForAggregate(ProductLangRemoved, &ProductLangRemoveData{
					LangCode: "Es_es",
				}, TimeNow(), AggregateProductType, id, 1),
			},
			nil,
		},
		"removeProductLangNotCreated": {
			&AggregateProduct{
				AggregateBase: events.NewAggregateBase(AggregateProductType, id),
				created:       false,
			},
			&RemoveProductLang{
				LangCode:  "Es_es",
				ProductID: id,
			},
			nil,
			errors.New("product not exist"),
		},
		"setAvailability": {
			&AggregateProduct{
				AggregateBase: events.NewAggregateBase(AggregateProductType, id),
				created:       true,
			},
			&SetAvailability{
				Availability: Availability{
					Quantity:          0,
					MinimalQuantity:   0,
					OnlineOnly:        false,
					OnSale:            false,
					OutOfStock:        false,
					Active:            false,
					AvailableForOrder: false,
					AvailableDate:     TimeNow(),
					Visibility:        0,
					DateAdd:           TimeNow(),
					DateUpd:           TimeNow(),
					QuantityDiscount:  false,
				},
				ProductID: "",
			},
			[]eh.Event{
				eh.NewEventForAggregate(AvailabilitySet, &AvailabilityData{
					Availability: Availability{
						Quantity:          0,
						MinimalQuantity:   0,
						OnlineOnly:        false,
						OnSale:            false,
						OutOfStock:        false,
						Active:            false,
						AvailableForOrder: false,
						AvailableDate:     TimeNow(),
						Visibility:        0,
						DateAdd:           TimeNow(),
						DateUpd:           TimeNow(),
						QuantityDiscount:  false,
					},
				}, TimeNow(), AggregateProductType, id, 1),
			},
			nil,
		},
	}

	for name, tc := range cases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			err := tc.agg.HandleCommand(context.Background(), tc.cmd)
			if (err != nil && tc.expectedErr == nil) ||
				(err == nil && tc.expectedErr != nil) ||
				(err != nil && tc.expectedErr != nil && err.Error() != tc.expectedErr.Error()) {
				t.Errorf("test case '%s': incorrect error", name)
				t.Log("exp:", tc.expectedErr)
				t.Log("got:", err)
			}
			events := tc.agg.Events()
			if !reflect.DeepEqual(events, tc.expectedEvents) {
				t.Errorf("test case '%s': incorrect events", name)
				t.Log("exp:\n", pretty.Sprint(tc.expectedEvents))
				t.Log("got:\n", pretty.Sprint(events))
			}
		})
	}

}
