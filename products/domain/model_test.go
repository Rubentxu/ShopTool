package domain

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	eh "github.com/looplab/eventhorizon"
)

func TestModelProductJSON(t *testing.T) {
	id := eh.NewUUID()
	now := time.Now()

	// Don't use keys for init, we want to get compiler warnings if we haven't
	// used some fields.
	p := &Product{
		ID:      id,
		Version: 1,
		Reference: "1234565¡a",
		Ean13: "ean13",
		Isbn: "978-84",
		Upc: "",
		ProductLangs: []*ProductLang{
			&ProductLang{
				Name:             "Pantalones Monoles",
				Description:      "Estos pantalones son lo ultimo",
				DescriptionShort: "pantalones guays",
				LinkRewrite:      "link",
				MetaDescription:  "meta descripcion",
				MetaKeywords:     "meta keywords",
				MetaTitle:        "meta titulo",
				AvailableNow:     "",
				AvailableLater:   "",
				LangCode:         "Es_es",
			},
		},

		Availability: Availability{
			Quantity:          5,
			MinimalQuantity:   2,
			OnlineOnly:        true,
			OnSale:            false,
			OutOfStock:        false,
			Active:            true,
			AvailableForOrder: true,
			AvailableDate:     now,
			Visibility:        0,
			DateAdd:           now,
			DateUpd:           now,
			QuantityDiscount:  false,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	var expectedJSONStr = []byte(`
{
	"id": "` + id.String() + `",
   "version":1,
   "productLangs":[
      {
         "name":"Pantalones Monoles",
         "description":"Estos pantalones son lo ultimo",
         "description_short":"pantalones guays",
         "link_rewrite":"link",
         "meta_description":"meta descripcion",
         "meta_keywords":"meta keywords",
         "meta_title":"meta titulo",
         "available_now":"",
         "available_later":"",
         "lang_code":"Es_es"
      }
   ],
   "availability":{
      "quantity":5,
      "minimal_quantity":2,
      "online_only":true,
      "on_sale":false,
      "out_of_stock bson":false,
      "active":true,
      "available_for_order":true,
      "available_date": "` + now.Format(time.RFC3339Nano) + `",
      "visibility":0,
      "date_add":"` + now.Format(time.RFC3339Nano) + `",
      "date_upd":"` + now.Format(time.RFC3339Nano) + `",
      "quantity_discount":false
   },
   "reference":"1234565¡a",
   "ean_13":"ean13",
   "isbn":"978-84",
   "upc":"",
   "created_at":"` + now.Format(time.RFC3339Nano) + `",
   "updated_at":"` + now.Format(time.RFC3339Nano) + `"
}`)
	expectedJSON := new(bytes.Buffer)
	if err := json.Compact(expectedJSON, expectedJSONStr); err != nil {
		t.Error(err)
	}

	js, err := json.Marshal(p)
	if err != nil {
		t.Error(err)
	}

	if string(js) != expectedJSON.String() {
		t.Error("the JSON should be correct:")
		t.Log("exp:", expectedJSON.String())
		t.Log("got:", string(js))
	}
}
