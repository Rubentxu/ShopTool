package domain

import (
	eh "github.com/looplab/eventhorizon"
)

func init() {
	eh.RegisterCommand(func() eh.Command { return &Create{} })
	eh.RegisterCommand(func() eh.Command { return &Delete{} })
	eh.RegisterCommand(func() eh.Command { return &AddProductLang{} })
	eh.RegisterCommand(func() eh.Command { return &UpdateProductLang{} })
	eh.RegisterCommand(func() eh.Command { return &RemoveProductLang{} })
}

// Command constants
const (
	CreateProductCommand     = eh.CommandType("product:create")
	DeleteProductCommand     = eh.CommandType("product:delete")
	AddProductLangCommand    = eh.CommandType("product:addProductLang")
	UpdateProductLangCommand = eh.CommandType("product:updateProductLang")
	RemoveProductLangCommand = eh.CommandType("product:removeProductLang")
)

// Static type check that the eventhorizon.Command interface is implemented.
var _ = eh.Command(&Create{})
var _ = eh.Command(&Delete{})
var _ = eh.Command(&AddProductLang{})
var _ = eh.Command(&UpdateProductLang{})
var _ = eh.Command(&RemoveProductLang{})

// Create creates a new todo list.
type Create struct {
	ID eh.UUID `json:"id"`
}

// AggregateType type for create
func (c *Create) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type for Create
func (c *Create) AggregateID() eh.UUID { return c.ID }

// CommandType type for Create
func (c *Create) CommandType() eh.CommandType { return CreateProductCommand }

// Delete deletes a todo list.
type Delete struct {
	ID eh.UUID `json:"id"`
}

// AggregateType type for delete
func (c *Delete) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type for Delete
func (c *Delete) AggregateID() eh.UUID { return c.ID }

// CommandType type for Delete
func (c *Delete) CommandType() eh.CommandType { return DeleteProductCommand }

// AddProductLang to Product
type AddProductLang struct {
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
	ProductID        eh.UUID `json:"id"`
}

// AggregateType type for AddProductLang
func (c *AddProductLang) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type for AddProductLang
func (c *AddProductLang) AggregateID() eh.UUID { return c.ProductID }

// CommandType type for AddProductLang
func (c *AddProductLang) CommandType() eh.CommandType { return AddProductLangCommand }

// UpdateProductLang to Product
type UpdateProductLang struct {
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
	ProductID        eh.UUID `json:"id"`
}

// AggregateType type for UpdateProductLang
func (c *UpdateProductLang) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type for UpdateProductLang
func (c *UpdateProductLang) AggregateID() eh.UUID { return c.ProductID }

// CommandType type for UpdateProductLang
func (c *UpdateProductLang) CommandType() eh.CommandType { return UpdateProductLangCommand }

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
func (c *RemoveProductLang) CommandType() eh.CommandType { return RemoveProductLangCommand }
