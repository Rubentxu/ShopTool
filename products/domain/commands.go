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

// Static type check that the eventhorizon.Command interface is implemented.
var _ = eh.Command(&AddProductLang{})

// AddProductLang to Product
type AddProductLang struct {
	ProductID        eh.UUID `json:"id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	DescriptionShort string  `json:"description_short"`
	LinkRewrite      string  `json:"link_rewrite"`
	MetaDescription  string  `json:"meta_description"`
	MetaKeywords     string  `json:"meta_keywords"`
	MetaTitle        string  `json:"meta_title"`
	AvailableNow     string  `json:"available_now"`
	AvailableLater   string  `json:"available_later"`
	LangCode         string  `json:"lang_code"`
}

// AggregateType type for AddProductLang
func (c *AddProductLang) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type for AddProductLang
func (c *AddProductLang) AggregateID() eh.UUID { return c.ProductID }

// CommandType type for AddProductLang
func (c *AddProductLang) CommandType() eh.CommandType { return AddProductLangCommand }
