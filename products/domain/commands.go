package domain

import (
	eh "github.com/looplab/eventhorizon"
)

func init() {
	eh.RegisterCommand(func() eh.Command { return &AddProductLang{} })
}

// Command constants
const (
	AddProductLangCommand = eh.CommandType("product:addProductLang")
)

// AddProductLang to Product
type AddProductLang struct {
	ProductID        eh.UUID `json:"id" bson:"id"`
	Name             string  `json:"name" bson:"name"`
	Description      string  `json:"description" bson:"description"`
	DescriptionShort string  `json:"description_short" bson:"description_short"`
	LinkRewrite      string  `json:"link_rewrite" bson:"link_rewrite"`
	MetaDescription  string  `json:"meta_description" bson:"meta_description"`
	MetaKeywords     string  `json:"meta_keywords" bson:"meta_keywords"`
	MetaTitle        string  `json:"meta_title" bson:"meta_title"`
	AvailableNow     string  `json:"available_now" bson:"available_now"`
	AvailableLater   string  `json:"available_later" bson:"available_later"`
	LangCode         string  `json:"lang_code" bson:"lang_code"`
}

// AggregateType type for AddProductLang
func (c *AddProductLang) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type for AddProductLang
func (c *AddProductLang) AggregateID() eh.UUID { return c.ProductID }

// CommandType type for AddProductLang
func (c *AddProductLang) CommandType() eh.CommandType { return AddProductLangCommand }
