package products

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"ShopTool/products/domain" // import "github.com/Rubentxu/ShopTool/products/domain"

	"github.com/gorilla/websocket"
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

// Handler is a http.Handler for the ShopTool app.
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

// NewHandler sets up the full Event Horizon domain for the ShopTool app and
// returns a handler exposing some of the components.
func NewHandler(dbURL string) (*Handler, error) {

	// Create the event store.
	eventStore, err := eventstore.NewEventStore(dbURL, "")
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
	repo, err := repo.NewRepo(dbURL, "ShopTool", "product")
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
	eventBus.AddHandler(projector, domain.ProductLangRemoved)

	// Handle the API.
	h := http.NewServeMux()
	h.Handle("/api/events/", EventBusHandler(eventPublisher))
	h.Handle("/api/product/", httputils.QueryHandler(productRepo))
	h.Handle("/api/product/command/create", httputils.CommandHandler(loggingHandler, domain.CreateProductCommand))
	h.Handle("/api/product/command/remove", httputils.CommandHandler(loggingHandler, domain.DeleteProductCommand))
	h.Handle("/api/product/command/prodlang/add", httputils.CommandHandler(loggingHandler, domain.AddProductLangCommand))
	h.Handle("/api/product/command/prodlang/update", httputils.CommandHandler(loggingHandler, domain.UpdateProductLangCommand))
	h.Handle("/api/product/command/prodlang/remove", httputils.CommandHandler(loggingHandler, domain.RemoveProductLangCommand))
	h.Handle("/api/product/command/availabilityConfig", httputils.CommandHandler(loggingHandler, domain.SetAvailabilityCommand))
	h.Handle("/api/product/command/transportSpecificationConfig", httputils.CommandHandler(loggingHandler, domain.SetTransportSpecificationCommand))
	h.Handle("/api/product/command/transport/add", httputils.CommandHandler(loggingHandler, domain.AddTransportCommand))
	h.Handle("/api/product/command/transport/update", httputils.CommandHandler(loggingHandler, domain.UpdateTransportCommand))
	h.Handle("/api/product/command/transport/remove", httputils.CommandHandler(loggingHandler, domain.RemoveTransportCommand))
	h.Handle("/api/product/docs/", http.StripPrefix("/api/product/docs/", http.FileServer(http.Dir("swagger-ui/"))))

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

// Observer is a simple event handler for observing events.
type Observer struct {
	EventCh chan eh.Event
}

// Notify implements the Notify method of the EventObserver interface.
func (o *Observer) Notify(ctx context.Context, event eh.Event) error {
	select {
	case o.EventCh <- event:
	default:
		log.Println("missed event:", event)
	}
	return nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// EventDef definition event
type EventDef struct {
	eventType     string
	data          eh.EventData
	timestamp     string
	aggregateType string
	aggregateID   string
	version       string
}

func (e *EventDef) MarshalJSON() ([]byte, error) {
	m := make(map[string]string)
	m["eventType"] = e.eventType
	data, _ := json.Marshal(e.data)
	m["data"] = string(data)
	m["aggregateType"] = e.aggregateType
	m["aggregateID"] = e.aggregateID
	m["version"] = e.version
	return json.Marshal(m)
}

// NewEvenDef eventDef
func NewEvenDef(event eh.Event) *EventDef {
	//time, _ := event.Timestamp().MarshalJSON()
	return &EventDef{
		eventType:     string(event.EventType()),
		data:          event.Data(),
		aggregateType: string(event.AggregateType()),
		aggregateID:   string(event.AggregateID()),
		version:       string(event.Version()),
	}

}

// EventBusHandler is a Websocket handler for eventhorizon.Events. Events will
// be forwarded to all requests that have been upgraded to websockets.
func EventBusHandler(eventPublisher eh.EventPublisher) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()

		observer := &Observer{
			EventCh: make(chan eh.Event, 10),
		}
		eventPublisher.AddObserver(observer)

		if err := c.WriteMessage(websocket.TextMessage, []byte("conectado a websocket")); err != nil {
			log.Println("write:", err)
		}

		for event := range observer.EventCh {
			eventString, err := NewEvenDef(event).MarshalJSON()
			if err != nil {
				fmt.Printf("Error parseo evento %s\n", err.Error())
			}
			fmt.Printf("evento json %s\n", eventString)
			if err := c.WriteMessage(websocket.TextMessage, []byte(string(eventString))); err != nil {
				log.Println("write:", err)
				break
			}
		}
	})
}
