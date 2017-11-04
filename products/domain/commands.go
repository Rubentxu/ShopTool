package domain

import (
	eh "github.com/looplab/eventhorizon"
)

func init() {
	eh.RegisterCommand(func() eh.Command { return &AddProductLang{} })
}

// Command constants
const (
	AddProductLangCommand    = eh.CommandType("product:addProductLang")
	UpdateProductLangCommand = eh.CommandType("product:updateProductLang")
	RemoveProductLangCommand = eh.CommandType("product:removeProductLang")
)

// Static type check that the eventhorizon.Command interface is implemented.
var _ = eh.Command(&AddProductLang{})
var _ = eh.Command(&UpdateProductLang{})
var _ = eh.Command(&RemoveProductLang{})

// AddProductLang to Product
type AddProductLang struct {
	ProductLang `json:"productLang"`
	ProductID   eh.UUID `json:"id"`
}

// AggregateType type for AddProductLang
func (c *AddProductLang) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type for AddProductLang
func (c *AddProductLang) AggregateID() eh.UUID { return c.ProductID }

// CommandType type for AddProductLang
func (c *AddProductLang) CommandType() eh.CommandType { return AddProductLangCommand }

// UpdateProductLang to Product
type UpdateProductLang struct {
	*ProductLang
	ProductID eh.UUID `json:"id"`
}

// AggregateType type for UpdateProductLang
func (c *UpdateProductLang) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type for UpdateProductLang
func (c *UpdateProductLang) AggregateID() eh.UUID { return c.ProductID }

// CommandType type for UpdateProductLang
func (c *UpdateProductLang) CommandType() eh.CommandType { return AddProductLangCommand }

// RemoveProductLang with langcode
type RemoveProductLang struct {
	LangCode  string  `json:"lang_code"`
	ProductID eh.UUID `json:"id"`
}

// AggregateType type for RemoveProductLang
func (c *RemoveProductLang) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type for RemoveProductLang
func (c *RemoveProductLang) AggregateID() eh.UUID { return c.ProductID }

// CommandType type for RemoveProductLang
func (c *RemoveProductLang) CommandType() eh.CommandType { return AddProductLangCommand }
