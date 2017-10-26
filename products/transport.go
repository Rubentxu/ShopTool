package products

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler.
// Useful in a profilesvc server.
func MakeHTTPHandler(s ProductService, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// POST    /product/                          adds another profile
	// GET     /product/:id                       retrieves the given profile by id
	// PUT     /product/:id                       post updated profile information about the profile
	// PATCH   /product/:id                       partial updated profile information
	// DELETE  /product/:id                       remove the given profile
	// GET     /product/:id/addresses/            retrieve addresses associated with the profile
	// GET     /product/:id/addresses/:addressID  retrieve a particular profile address
	// POST    /product/:id/addresses/            add a new address
	// DELETE  /product/:id/addresses/:addressID  remove an address

	// r.Methods("POST").Path("/product/").Handler(httptransport.NewServer(
	// 	e.PostProfileEndpoint,
	// 	decodePostProfileRequest,
	// 	encodeResponse,
	// 	options...,
	// ))
	r.Methods("GET").Path("/product").Handler(httptransport.NewServer(
		e.GetAllProductsEndPoint,
		decodeGetProductRequest,
		encodeResponse,
		options...,
	))
	// r.Methods("GET").Path("/product/{id}").Handler(httptransport.NewServer(
	// 	e.GetProfileEndpoint,
	// 	decodeGetProfileRequest,
	// 	encodeResponse,
	// 	options...,
	// ))
	// r.Methods("DELETE").Path("/product/{id}").Handler(httptransport.NewServer(
	// 	e.DeleteProfileEndpoint,
	// 	decodeDeleteProfileRequest,
	// 	encodeResponse,
	// 	options...,
	// ))
	// r.Handle("/api/docs", http.FileServer(http.Dir("swagger-ui")))
	// r.PathPrefix("/api/docs").Handler(http.StripPrefix("/api/docs", )))
	//r.Methods("GET").Path("/api").Handler(http.FileServer(http.Dir("../swagger-ui")))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("swagger-ui/")))
	return r
}

type getAllProductRequest struct{}

func decodeGetProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	// vars := mux.Vars(r)
	// id, ok := vars["id"]
	// if !ok {
	// 	return nil, ErrBadRouting
	// }
	return getAllProductRequest{}, nil
}

type errorer interface {
	error() error
}

// encodeResponse is the common method to encode all response types to the
// client. I chose to do it this way because, since we're using JSON, there's no
// reason to provide anything more specific. It's certainly possible to
// specialize on a per-response (per-method) basis.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
