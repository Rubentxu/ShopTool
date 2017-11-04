package domain

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/kr/pretty"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
)

func TestAggregateHandleCommand(t *testing.T) {
	TimeNow = func() time.Time {
		return time.Date(2017, time.November, 20, 42, 0, 0, 0, time.Local)
	}

	id := eh.NewUUID()
	cases := map[string]struct {
		agg            *AggregateProduct
		cmd            eh.Command
		expectedEvents []eh.Event
		expectedErr    error
	}{
		"addProductLang": {
			&AggregateProduct{
				AggregateBase: events.NewAggregateBase(AggregateProductType, id),
			},
			&AddProductLang{
				ProductID: id,
				ProductLang: &ProductLang{
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
				eh.NewEventForAggregate(ProductLangAdded, &ProductLangAddedData{
					ProductLang: &ProductLang{
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
