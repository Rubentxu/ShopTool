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

// Product  is the read model for the product.
type Product struct {
	ID           eh.UUID        `json:"id" bson:"id"`
	Version      int            `json:"version" bson:"version"`
	ProductLangs []*ProductLang `json:"productLangs" bson:"productLangs"`
	CreatedAt    time.Time      `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" bson:"updated_at"`
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
