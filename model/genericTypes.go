package model

import (
	"github.com/looplab/eventhorizon"
	"image/color"
	"time"
	"net/url"
)

type Thing struct {
	Name          string
	AlternateName string
	Description   string
}

type QuantitativeValue struct {
	UnitCode string
	UnitText string
	Value    int64
	MaxValue int64
	MinValue int64
}

type PropertyValue struct {
	PropertyID string
	UnitCode   string
	UnitText   string
	Value      string
}

type Rating struct {
	Author      string
	BestRating  int8
	RatingValue int8
	WorstRating int8
}

type AggregateRating struct {
	Rating
	ItemReviewed Thing
	RatingCount  int
	ReviewCount  int
}

type ImageObject struct {
	Thing
	Caption        string
	Thumbnail      ImageObject
	ContentSize    string
	EncodingFormat string
	Height         int32
	Width          int32
	Source         string
}

type Audience struct {
	AudienceType   string
	GeographicArea Place
}

type Brand struct {
	AggregateRating
	Logo ImageObject
}

type ProductId string

type OrganizationId string

type ProductModel struct {
	Product
	IsVariantOf   ProductId
	PredecessorOf ProductId
	SuccessorOf   ProductId
}

type DeliveryMethod int

type OfferItemCondition int

type ItemAvailability int

type BusinessFunction int

type PhysicalActivityCategory int

type BusinessEntityType int

type TypeAndQuantityNode struct {
	AmountOfThisGood int
	BusinessFunction
	TypeOfGood ProductId
	UnitCode   string
	UnitText   string
}

type WarrantyScope int

type WarrantyPromise struct {
	DurationOfWarranty time.Duration
	WarrantyScope
}

type Offer struct {
	Thing
	AddOn                     Offer
	AdvanceBookingRequirement QuantitativeValue
	AggregateRating
	AreaServed              Place
	Availability            ItemAvailability
	AvailabilityStarts      time.Time
	AvailabilityEnds        time.Time
	AvailableAtOrFrom       Place
	AvailableDeliveryMethod DeliveryMethod
	BusinessFunction
	Category                  PhysicalActivityCategory
	DeliveryLeadTime          QuantitativeValue
	EligibleCustomerType      BusinessEntityType
	EligibleDuration          QuantitativeValue
	EligibleQuantity          QuantitativeValue
	EligibleRegion            Place
	EligibleTransactionVolume PriceSpecification
	Gtin8                     string
	Gtin12                    string
	Gtin13                    string
	Gtin14                    string
	IncludesObject            TypeAndQuantityNode
	IneligibleRegion          Place
	InventoryLevel            QuantitativeValue
	ItemCondition             OfferItemCondition
	ItemOffered               ProductId
	MPN                       string
	OfferedBy                 OrganizationId
	Price                     int
	PriceCurrency             string
	PriceSpecification
	PriceValidUntil time.Time
	SerialNumber    string
	Sku             string
	ValidFrom       time.Time
	ValidThrough    time.Time
	Warranty        WarrantyPromise
}

type PriceSpecification struct {
	Thing
	EligibleQuantity          QuantitativeValue
	EligibleTransactionVolume PriceSpecification
	MaxPrice                  int
	MinPrice                  int
	Price                     int
	PriceCurrency             string
	ValidFrom                 time.Time
	ValidThrough              time.Time
	ValueAddedTaxIncluded     bool
}

type Product struct {
	ID ProductId
	Thing
	AggregateRating
	Audience
	Brand
	AdditionalProperty        PropertyValue
	Award                     string
	Category                  Thing
	Gtin8                     string
	Gtin12                    string
	Gtin13                    string
	Gtin14                    string
	IsAccessoryOrSparePartFor ProductId
	IsConsumableFor           ProductId
	IsSimilarTo               ProductId
	ItemCondition             string
	Logo                      ImageObject
	Manufacturer              OrganizationId
	Material                  string
	Model                     ProductModel
	MPN                       string
	Offers                    []Offer
	ProductDate               time.Time
	PurchaseDate              time.Time
	ReleaseDate               time.Time
	Sku                       string
	Weight                    QuantitativeValue
	Width                     QuantitativeValue
	Color                     color.Color
	Depth                     QuantitativeValue
	Height                    QuantitativeValue
}

type Day int

const (
	MONDAY Day = 1 + iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

type OpeningHoursSpecification struct {
	Opens     time.Time
	Closes    time.Time
	DayOfWeek Day
}

type ContactPoint struct {
	AvailableLanguage string
	ContactType       string
	Email             string
	FaxNumber         string
	HoursAvailable    OpeningHoursSpecification
	Telephone         string
}

type PostalAddress struct {
	AddressCountry      string
	AddressLocality     string
	AddressRegion       string
	PostOfficeBoxNumber string
	PostalCode          string
	StreetAddress       string
}

type LocationFeatureSpecification struct {
	HoursAvailable OpeningHoursSpecification
	ValidFrom      time.Time
	ValidThrough   time.Time
}

type GeoCoordinates struct {
	Elevation int
	Latitude  float32
	Longitude float32
}

type PlaceId string

type Place struct {
	ID               PlaceId
	Address          PostalAddress
	AmenityFeature   LocationFeatureSpecification
	BranchCode       string
	ContainedInPlace PlaceId
	ContainsPlace    []PlaceId
	Geo              GeoCoordinates
}

type User struct {

}

type DeliveryEvent struct {
	AccessCode string
	AvailableFrom time.Time
	AvailableThrough time.Time
	HasDeliveryMethod DeliveryMethod
}

type ParcelDelivery struct {
	DeliveryAddress PostalAddress
	DeliveryStatus DeliveryEvent
	ExpectedArrivalFrom time.Time
	ExpectedArrivalUntil time.Time
	HasDeliveryMethod DeliveryMethod
	ItemShipped ProductId
	OriginalAddress PostalAddress
	PartOfOrder Order
	Provider OrganizationId
	TrackingNumber string
	TrackingUrl url.URL

}

type OrderStatus int

type OrderItem struct {
	OrderDelivery ParcelDelivery
	OrderItemNumber string
	OrderItemStatus OrderStatus
	OrderQuantity int
	OrderedItem ProductId
}

type Order struct {
	Thing
	AcceptedOffer Offer
	BillingAddress PostalAddress
	Broker OrganizationId
	ConfirmationNumber string
	Customer User
	Discount int
	DiscountCode string
	DiscountCurrency string
	IsGift bool
	OrderDate time.Time
	OrderDelivery ParcelDelivery
	OrderNumber string
	OrderStatus
	OrderedItem OrderItem
}

