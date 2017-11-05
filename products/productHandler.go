package products

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"shopTool/products/domain"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/commandhandler/aggregate"
	eventbus "github.com/looplab/eventhorizon/eventbus/local"
	"github.com/looplab/eventhorizon/eventhandler/projector"
	eventstore "github.com/looplab/eventhorizon/eventstore/mongodb"
	"github.com/looplab/eventhorizon/httputils"
	eventpublisher "github.com/looplab/eventhorizon/publisher/local"
	repo "github.com/looplab/eventhorizon/repo/mongodb"
	"github.com/looplab/eventhorizon/repo/version"
)

// Handler is a http.Handler for the shoptool app.
type Handler struct {
	http.Handler

	CommandHandler eh.CommandHandler
	Repo           eh.ReadWriteRepo
}

// Logger is a simple event handler for logging all events.
type Logger struct{}

// Notify implements the Notify method of the EventObserver interface.
func (l *Logger) Notify(ctx context.Context, event eh.Event) error {
	log.Printf("EVENT %s", event)
	return nil
}

// UUIDHandler Get productid
func UUIDHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "unsuported method: "+r.Method, http.StatusMethodNotAllowed)
			return
		}
		uuid := eh.NewUUID()
		b, err := json.Marshal(uuid)
		if err != nil {
			http.Error(w, "could not encode result: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(b)
	})
}

// NewHandler sets up the full Event Horizon domain for the shoptool app and
// returns a handler exposing some of the components.
func NewHandler(dbURL string) (*Handler, error) {

	// Create the event store.
	eventStore, err := eventstore.NewEventStore(dbURL, "shoptool")
	if err != nil {
		return nil, fmt.Errorf("could not create event store: %s", err)
	}

	// Create the event bus that distributes events.
	eventBus := eventbus.NewEventBus()
	eventPublisher := eventpublisher.NewEventPublisher()
	eventPublisher.AddObserver(&Logger{})
	eventBus.SetPublisher(eventPublisher)

	// Create the aggregate repository.
	aggregateStore, err := events.NewAggregateStore(eventStore, eventBus)
	if err != nil {
		return nil, fmt.Errorf("could not create aggregate store: %s", err)
	}

	// Create the aggregate command handler.
	commandHandler, err := aggregate.NewCommandHandler(domain.AggregateProductType, aggregateStore)
	if err != nil {
		return nil, fmt.Errorf("could not create command handler: %s", err)
	}

	// Create a tiny logging middleware for the command handler.
	loggingHandler := eh.CommandHandlerFunc(func(ctx context.Context, cmd eh.Command) error {
		log.Printf("CMD %#v", cmd)
		return commandHandler.HandleCommand(ctx, cmd)
	})
	// Create the repository and wrap in a version repository.
	repo, err := repo.NewRepo(dbURL, "shoptool", "product")
	if err != nil {
		return nil, fmt.Errorf("could not create invitation repository: %s", err)
	}
	repo.SetEntityFactory(func() eh.Entity { return &domain.Product{} })
	productRepo := version.NewRepo(repo)

	// Create the read model projector.
	projector := projector.NewEventHandler(&domain.ProductProjector{}, productRepo)
	projector.SetEntityFactory(func() eh.Entity { return &domain.Product{} })
	eventBus.AddHandler(projector, domain.ProductCreated)
	eventBus.AddHandler(projector, domain.ProductDeleted)
	eventBus.AddHandler(projector, domain.ProductLangAdded)
	eventBus.AddHandler(projector, domain.ProductLangUpdated)
	eventBus.AddHandler(projector, domain.ProductLangRemove)

	// Handle the API.
	h := http.NewServeMux()
	h.Handle("/api/events/", httputils.EventBusHandler(eventPublisher))
	h.Handle("/api/product/", httputils.QueryHandler(productRepo))
	h.Handle("/api/product/create", httputils.CommandHandler(loggingHandler, domain.CreateProductCommand))
	h.Handle("/api/product/remove", httputils.CommandHandler(loggingHandler, domain.DeleteProductCommand))
	h.Handle("/api/product/prodlang/add", httputils.CommandHandler(loggingHandler, domain.AddProductLangCommand))
	h.Handle("/api/product/prodlang/update", httputils.CommandHandler(loggingHandler, domain.UpdateProductLangCommand))
	h.Handle("/api/product/prodlang/remove", httputils.CommandHandler(loggingHandler, domain.RemoveProductLangCommand))

	logger := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		h.ServeHTTP(w, r)
	})

	return &Handler{
		Handler:        logger,
		CommandHandler: loggingHandler,
		Repo:           productRepo,
	}, nil

}
