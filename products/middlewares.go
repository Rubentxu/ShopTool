package products

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(ProductService) ProductService

// LoggingMiddleware create
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next ProductService) ProductService {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   ProductService
	logger log.Logger
}

func (mw loggingMiddleware) GetProduct(ctx context.Context, id string) (p Product, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetProduct", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetProduct(ctx, id)
}

func (mw loggingMiddleware) GetAllProducts(ctx context.Context) (p []Product, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetAllProducts", "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetAllProducts(ctx)
}

func (mw loggingMiddleware) DeleteProduct(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "DeleteProduct", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.DeleteProduct(ctx, id)
}

func (mw loggingMiddleware) CreateProduct(ctx context.Context, p *Product) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "CreateProduct", "ProductID", p.ID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.CreateProduct(ctx, p)
}
