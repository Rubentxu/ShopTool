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
type Unity int
type Condition int

const (
	BOTH Visibility = 1 + iota
	CATALOG
	SEARCH
	NONE
)

const (
	KILOS Unity = 1 + iota
	GRAMOS
	PIEZAS
)

const (
	NEW Condition = 1 + iota
	USED
	REFURBISHED
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

type Transporter struct {
	Id          eh.UUID `json:"transporter_id" bson:"transporter_id"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
}

type Type struct {
	Name            string    `json:"name" bson:"name"`
	Description     string    `json:"description" bson:"description"`
	Unity           Unity     `json:"unity" bson:"unity"`
	IsSimilar       bool      `json:"is_similar" bson:"is_similar"`
	Model           string    `json:"model" bson:"model"`
	Condition       Condition `json:"condition" bson:"condition"`
	Characteristics []Characteristic
}

type Characteristic struct {
	Name         string `json:"name" bson:"name"`
	Value        string `json:"value" bson:"value"`
	DefaultValue string `json:"default_value" bson:"default_value"`
}

type TransportSpecification struct {
	Width                  uint8         `json:"width" bson:"width"`
	Height                 uint8         `json:"height" bson:"height"`
	Depth                  uint8         `json:"depth" bson:"depth"`
	Weight                 uint8         `json:"weight" bson:"weight"`
	AdditionalShippingCost uint8         `json:"additional_shipping_cost" bson:"additional_shipping_cost"`
	Transporters           []Transporter `json:"transporters" bson:"transporters"`
}

type PricesSpecification struct {
	IdTaxtRulesGroup      eh.UUID   `json:"id_taxt_rules_group" bson:"id_taxt_rules_group"`
	Price                 int64     `json:"price" bson:"price"`
	WholeSalePrice        int64     `json:"whole_sale_price" bson:"whole_sale_price"`
	UnitPriceRatio        string    `json:"unit_price_ratio" bson:"unit_price_ratio"`
	ShowPrice             bool      `json:"show_price" bson:"show_price"`
	PrioritiesForRules    []string  `json:"priorities_for_rules" bson:"priorities_for_rules"`
	Ecotax                int       `json:"ecotax" bson:"ecotax"`
	PriceCurrency         string    `json:"price_currency" bson:"price_currency"`
	ValidFrom             time.Time `json:"valid_from" bson:"valid_from"`
	ValidThrough          time.Time `json:"valid_through" bson:"valid_through"`
	ValueAddedTaxIncluded bool      `json:"value_added_tax_included" bson:"value_added_tax_included"`
	MaxPrice              int       `json:"max_price" bson:"max_price"`
	MinPrice              int       `json:"min_price" bson:"min_price"`
}

type Image struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Caption     string `json:"caption" bson:"caption"`
	Thumbnail   string `json:"thumbnail" bson:"thumbnail"`
}

// Product  is the read model for the product.
type Product struct {
	ID                     eh.UUID        `json:"id" bson:"id"`
	Version                int            `json:"version" bson:"version"`
	Reference              string         `json:"reference" bson:"reference"`
	Ean13                  string         `json:"ean_13" bson:"ean_13"`
	Isbn                   string         `json:"isbn" bson:"isbn"`
	Upc                    string         `json:"upc" bson:"upc"`
	CreatedAt              time.Time      `json:"created_at" bson:"created_at"`
	UpdatedAt              time.Time      `json:"updated_at" bson:"updated_at"`
	ProductLangs           []*ProductLang `json:"productLangs" bson:"productLangs"`
	Availability           `json:"availability" bson:"availability"`
	TransportSpecification `json:"transport_specification" bson:"transport_specification"`
	PricesSpecification    `json:"price_specification" bson:"price_specification"`
	Images                 []Image
	ProductType            Type
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
