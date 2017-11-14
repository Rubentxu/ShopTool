package domain

import (
	"time"

	eh "github.com/looplab/eventhorizon"
)

// ProductLang for product
type ProductLang struct {
	Name             string `json:"name" bson:"name"`
	Description      string `json:"description" bson:"description"`
	DescriptionShort string `json:"description_short" bson:"description_short"`
	LinkRewrite      string `json:"link_rewrite" bson:"link_rewrite"`
	MetaDescription  string `json:"meta_description" bson:"meta_description"`
	MetaKeywords     string `json:"meta_keywords" bson:"meta_keywords"`
	MetaTitle        string `json:"meta_title" bson:"meta_title"`
	AvailableNow     string `json:"available_now" bson:"available_now"`
	AvailableLater   string `json:"available_later" bson:"available_later"`
	LangCode         string `json:"lang_code" bson:"lang_code"`
}

type Visibility int

const (
	BOTH Visibility = 1 + iota
	CATALOG
	SEARCH
	NONE
)

type Availability struct {
	Quantity          int        `json:"quantity" bson:"quantity"`
	MinimalQuantity   int        `json:"minimal_quantity" bson:"minimal_quantity"`
	OnlineOnly        bool       `json:"online_only" bson:"online_only"`
	OnSale            bool       `json:"on_sale" bson:"on_sale"`
	OutOfStock        bool       `json:"out_of_stock bson" bson:"out_of_stock"`
	Active            bool       `json:"active" bson:"active"`
	AvailableForOrder bool       `json:"available_for_order" bson:"available_for_order"`
	AvailableDate     time.Time  `json:"available_date" bson:"available_date"`
	Visibility        Visibility `json:"visibility" bson:"visibility"`
	DateAdd           time.Time  `json:"date_add" bson:"date_add"`
	DateUpd           time.Time  `json:"date_upd" bson:"date_upd"`
	QuantityDiscount  bool       `json:"quantity_discount" bson:"quantity_discount"`
}

// Product  is the read model for the product.
type Product struct {
	ID           eh.UUID        `json:"id" bson:"id"`
	Version      int            `json:"version" bson:"version"`
	ProductLangs []*ProductLang `json:"productLangs" bson:"productLangs"`
	Availability `json:"availability" bson:"availability"`
	Reference    string    `json:"reference" bson:"reference"`
	Ean13        string    `json:"ean_13" bson:"ean_13"`
	Isbn         string    `json:"isbn" bson:"isbn"`
	Upc          string    `json:"upc" bson:"upc"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}

var _ = eh.Entity(&Product{})
var _ = eh.Versionable(&Product{})

// EntityID implements the EntityID method of the eventhorizon.Entity interface.
func (p *Product) EntityID() eh.UUID {
	return p.ID
}

// AggregateVersion implements the AggregateVersion method of the
// eventhorizon.Versionable interface.
func (p *Product) AggregateVersion() int {
	return p.Version
}
