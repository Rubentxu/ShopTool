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
	eh.RegisterCommand(func() eh.Command { return &SetAvailability{} })
	eh.RegisterCommand(func() eh.Command { return &SetTransportSpecification{} })
	eh.RegisterCommand(func() eh.Command { return &AddTransport{} })
	eh.RegisterCommand(func() eh.Command { return &UpdateTransport{} })
	eh.RegisterCommand(func() eh.Command { return &RemoveTransport{} })
	eh.RegisterCommand(func() eh.Command { return &SetPricesSpecification{} })

}

// Command constants
const (
	CreateProductCommand             = eh.CommandType("product:create")
	DeleteProductCommand             = eh.CommandType("product:delete")
	AddProductLangCommand            = eh.CommandType("product:addProductLang")
	UpdateProductLangCommand         = eh.CommandType("product:updateProductLang")
	RemoveProductLangCommand         = eh.CommandType("product:removeProductLang")
	SetAvailabilityCommand           = eh.CommandType("product:setAvailability")
	SetTransportSpecificationCommand = eh.CommandType("product:setTransportSpecification")
	AddTransportCommand              = eh.CommandType("product:transportSpecification:addTransport")
	UpdateTransportCommand           = eh.CommandType("product:transportSpecification:updateTransport")
	RemoveTransportCommand           = eh.CommandType("product:transportSpecification:removeTransport")
	SetPricesSpecificationCommand    = eh.CommandType("product:pricesSpecification")
)

// Static type check that the eventhorizon.Command interface is implemented.
var _ = eh.Command(&Create{})
var _ = eh.Command(&Delete{})
var _ = eh.Command(&AddProductLang{})
var _ = eh.Command(&UpdateProductLang{})
var _ = eh.Command(&RemoveProductLang{})
var _ = eh.Command(&SetAvailability{})
var _ = eh.Command(&SetTransportSpecification{})
var _ = eh.Command(&AddTransport{})
var _ = eh.Command(&UpdateTransport{})
var _ = eh.Command(&RemoveTransport{})
var _ = eh.Command(&SetPricesSpecification{})

// Create creates a new todo list.
type Create struct {
	Reference string `json:"reference"`
	Ean13     string `json:"ean_13"`
	Isbn      string `json:"isbn"`
	Upc       string `json:"upc"`
}

// AggregateType type for create
func (c *Create) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type for Create
func (c *Create) AggregateID() eh.UUID { return idGen() }

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
	ProductLang
	ProductID eh.UUID `json:"id"`
}

// AggregateType type for AddProductLang
func (c *AddProductLang) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type for AddProductLang
func (c *AddProductLang) AggregateID() eh.UUID { return c.ProductID }

// CommandType type for AddProductLang
func (c *AddProductLang) CommandType() eh.CommandType { return AddProductLangCommand }

// UpdateProductLang to Product
type UpdateProductLang struct {
	ProductLang
	ProductID eh.UUID `json:"id"`
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

// SetAvailability definición de disponibilidad de producto
type SetAvailability struct {
	Availability
	ProductID eh.UUID `json:"id"`
}

// AggregateType type para SetAvailability
func (c *SetAvailability) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type para SetAvailability
func (c *SetAvailability) AggregateID() eh.UUID { return c.ProductID }

// CommandType type para SetAvailability
func (c *SetAvailability) CommandType() eh.CommandType { return SetAvailabilityCommand }

// SetSetTransportSpecification definición de disponibilidad de producto
type SetTransportSpecification struct {
	TransportSpecification
	ProductID eh.UUID `json:"id"`
}

// AggregateType type para SetTransportSpecification
func (c *SetTransportSpecification) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type para SetTransportSpecification
func (c *SetTransportSpecification) AggregateID() eh.UUID { return c.ProductID }

// CommandType type para SetTransportSpecification
func (c *SetTransportSpecification) CommandType() eh.CommandType {
	return SetTransportSpecificationCommand
}

type AddTransport struct {
	Transporter
	ProductID eh.UUID `json:"id"`
}

// AggregateType type para TransporterItem
func (c *AddTransport) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type para TransporterItem
func (c *AddTransport) AggregateID() eh.UUID { return c.ProductID }

// CommandType type para TransporterItem
func (c *AddTransport) CommandType() eh.CommandType {
	return AddTransportCommand
}

type UpdateTransport struct {
	Transporter
	ProductID eh.UUID `json:"id" b`
}

// AggregateType type para TransporterItem
func (c *UpdateTransport) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type para TransporterItem
func (c *UpdateTransport) AggregateID() eh.UUID { return c.ProductID }

// CommandType type para TransporterItem
func (c *UpdateTransport) CommandType() eh.CommandType {
	return UpdateTransportCommand
}

type RemoveTransport struct {
	transportID eh.UUID `json:"transporter_id"`
	ProductID   eh.UUID `json:"id"`
}

// AggregateType type para TransporterItem
func (c *RemoveTransport) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type para TransporterItem
func (c *RemoveTransport) AggregateID() eh.UUID { return c.ProductID }

// CommandType type para TransporterItem
func (c *RemoveTransport) CommandType() eh.CommandType {
	return RemoveTransportCommand
}

type SetPricesSpecification struct {
	PricesSpecification
	ProductID   eh.UUID `json:"id"`
}

// AggregateType type para SetPricesSpecification
func (c *SetPricesSpecification) AggregateType() eh.AggregateType { return AggregateProductType }

// AggregateID type para SetPricesSpecification
func (c *SetPricesSpecification) AggregateID() eh.UUID { return c.ProductID }

// CommandType type para SetPricesSpecification
func (c *SetPricesSpecification) CommandType() eh.CommandType {
	return RemoveTransportCommand
}