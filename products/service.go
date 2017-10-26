package products

import (
	"context"
	"errors"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Product type
type Product struct {
	ID                     bson.ObjectId `json:"id"`
	IDSupplier             bson.ObjectId `json:"id_supplier"`
	IDManufacturer         bson.ObjectId `json:"id_manufacturer"`
	IDCategoryDefault      bson.ObjectId `json:"id_category_default"`
	IDShopDefault          bson.ObjectId `json:"id_shop_default"`
	OnSale                 bool          `json:"on_sale"`
	OnlineOnly             bool          `json:"online_only"`
	Ean13                  string        `json:"ean13"`
	Isbn                   string        `json:"isbn"`
	Upc                    string        `json:"upc"`
	Ecotax                 float32       `json:"ecotax"`
	Quantity               int32         `json:"quantity"`
	MinimalQuantity        int32         `json:"minimal_quantity"`
	Price                  float32       `json:"price"`
	WholesalePrice         float32       `json:"wholesale_price"`
	Unity                  string        `json:"unity"`
	UnitPriceRation        float32       `json:"unit_price_ratio"`
	AdditionalShippingCost float32       `json:"additional_shipping_cost"`
	Reference              string        `json:"reference"`
	SupplierReference      string        `json:"supplier_reference"`
	Location               string        `json:"location"`
	Width                  float32       `json:"width"`
	Height                 float32       `json:"height"`
	Depth                  float32       `json:"depth"`
	Weight                 float32       `json:"weight"`
	OutOfStock             int32         `json:"out_of_stock"`
	QuantityDiscount       bool          `json:"quantity_discount"`
	Customizable           int32         `json:"customizable"`
	UploadableFiles        int32         `json:"uploadable_files"`
	TextFields             int32         `json:"text_fields"`
	Active                 int32         `json:"active"`
	Redirect               RedirectType  `json:"redirect_type"`
	IDTypeRedirected       int32         `json:"id_type_redirected"`
	AvailableForOrder      bool          `json:"available_for_order"`
	AvailableDate          time.Time     `json:"available_date"`
	ShowCondition          bool          `json:"show_condition"`
	Condition              Condition     `json:"condition"`
	ShowPrice              bool          `json:"show_price"`
	Indexed                bool          `json:"indexed"`
	Visibility             Visibility    `json:"visibility"`
	CacheIsPack            bool          `json:"cache_is_pack"`
	CacheHasAtachments     bool          `json:"cache_has_atachments"`
	IsVirtual              bool          `json:"is_virtual"`
	CacheDefaultAttribute  int32         `json:"cache_default_attribute"`
	DateAdd                time.Time     `json:"date_add"`
	DateUpd                time.Time     `json:"date_upd"`
	AdvanceStockManagement bool          `json:"advanced_stock_management"`
	PackStockType          bool          `json:"pack_stock_type"`
	State                  bool          `json:"state"`
	ProductsLang           []ProductLang `json:"productsLang"`
}

// ProductLang for products type
type ProductLang struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	DescriptionShort string `json:"description_short"`
	LinkRewrite      string `json:"link_rewrite"`
	MetaDescription  string `json:"meta_description"`
	MetaKeywords     string `json:"meta_keywords"`
	MetaTitle        string `json:"meta_title"`
	AvailableNow     string `json:"available_now"`
	AvailableLater   string `json:"available_later"`
	LangCode         string `json:"lang_code"`
}

// RedirectType in products
type RedirectType string

const (
	// None redirects
	None RedirectType = "NONE"
	// NotFound redirect// Service for products
	NotFound RedirectType = "404"
	// MovedPermanentlyProduct redirect
	MovedPermanentlyProduct RedirectType = "301PRODUCT"
	// FoundProduct redirect
	FoundProduct RedirectType = "302PRODUCT"
	// MovedPermanentlyCategory redirect
	MovedPermanentlyCategory RedirectType = "301CATEGORY"
	// FoundCategory redirect
	FoundCategory RedirectType = "302CATEGORY"
)

// Condition type
type Condition string

// Conditions products
const (
	New         Condition = "New"
	Used        Condition = "Used"
	Refurbished Condition = "Refurbished"
)

// Visibility type
type Visibility int

// Visibility products
const (
	Both = iota
	Catalog
	Search
	Hidden
)

// Errors from service
var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

// ProductService for products
type ProductService interface {
	GetProduct(ctx context.Context, id string) (Product, error)
	GetAllProducts(ctx context.Context) ([]Product, error)
	DeleteProduct(ctx context.Context, id string) error
	CreateProduct(ctx context.Context, p *Product) error
}

type productService struct {
	repo *ProductRepository
}

// NewProductService create service products
func NewProductService() (ProductService, error) {
	var p ProductService
	r, err := NewProductRepository()
	if err != nil {
		log.Fatalf("[Error Create ProductService]: %s\n", err)
		return p, err
	}
	return &productService{
		repo: &r,
	}, nil
}

func (s *productService) GetProduct(ctx context.Context, id string) (Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) GetAllProducts(ctx context.Context) ([]Product, error) {
	return s.repo.GetAll(), nil
}

func (s *productService) DeleteProduct(ctx context.Context, id string) error {
	return s.repo.Delete(id)
}

func (s *productService) CreateProduct(ctx context.Context, p *Product) error {
	return s.repo.Create(p)
}
