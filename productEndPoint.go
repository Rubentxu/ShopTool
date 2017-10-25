package shoptool

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them
// into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	// GetProductEndPoint endpoint.Endpoint
	GetAllProductsEndPoint endpoint.Endpoint
	// DeleteProductEndPoint  endpoint.Endpoint
	// CreateProductEndPoint  endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service. Useful in a profilesvc
// server.
func MakeServerEndpoints(s ProductService) Endpoints {
	return Endpoints{
		// GetProductEndPoint: MakeGetProductEndPoint(s),
		GetAllProductsEndPoint: MakeGetAllProductsEndPoint(s),
		// DeleteProductEndPoint:  MakeDeleteProductEndPoint(s),
		// CreateProductEndPoint:  MakeCreateProductEndPoint(s),
	}
}

type getProductRequest struct {
	ID string
}

type getProductResponse struct {
	Product Product `json:"product,omitempty"`
	Err     error   `json:"err,omitempty"`
}

type getAllProductResponse struct {
	Product []Product `json:"products,omitempty"`
	Err     error     `json:"err,omitempty"`
}

// MakeGetProductEndPoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeGetProductEndPoint(s ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getProductRequest)
		p, e := s.GetProduct(ctx, req.ID)
		return getProductResponse{Product: p, Err: e}, nil
	}
}

// MakeGetAllProductsEndPoint get all products
func MakeGetAllProductsEndPoint(s ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// req := request.(getAllProductRequest)
		p, e := s.GetAllProducts(ctx)
		return getAllProductResponse{Product: p, Err: e}, nil
	}
}
