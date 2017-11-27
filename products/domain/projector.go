package domain

import (
	"context"
	"errors"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/eventhandler/projector"

	"fmt"

)

// ProductProjector is a projector of product events on the Product read model.
type ProductProjector struct{}

// ProjectorType implements the ProjectorType method of the
// eventhorizon.Projector interface.
func (p *ProductProjector) ProjectorType() projector.Type {
	return projector.Type(string(AggregateProductType) + "_projector")
}

// Project implements the Project method of the eventhorizon.Projector interface.
func (p *ProductProjector) Project(ctx context.Context, event eh.Event, entity eh.Entity) (eh.Entity, error) {
	println("Projector project")
	model, ok := entity.(*Product)
	if !ok {
		return nil, errors.New("model is of incorrect type")
	}

	switch event.EventType() {
	case ProductCreated:
		println("Projector ProductCreated")
		data, ok := event.Data().(*CreateData)
		if !ok {
			return nil, errors.New("invalid event data ProductLangAdded")
		}

		// Set the ID when first created.
		model.ID = event.AggregateID()
		model.ProductLangs = []*ProductLang{} // Prevents "null" in JSON.
		model.Images = []Image{}
		model.Reference = data.Reference
		model.Upc = data.Upc
		model.Isbn = data.Isbn
		model.Ean13 = data.Ean13
		model.CreatedAt = TimeNow()
	case ProductDeleted:
		// Return nil as the entity to delete the model.
		return nil, nil
	case ProductLangAdded:
		println("Projector ProductLangAdded")
		data, ok := event.Data().(*ProductLangData)
		if !ok {
			return nil, errors.New("invalid event data ProductLangAdded")
		}
		productLang := &ProductLang{}
		*productLang = *&data.ProductLang
		model.ProductLangs = append(model.ProductLangs, productLang)
	case ProductLangUpdated:
		println("Projector ProductLangAdded")
		data, ok := event.Data().(*ProductLangData)
		if !ok {
			return nil, errors.New("invalid event data ProductLangAdded")
		}
		productLang := &ProductLang{}
		*productLang = *&data.ProductLang
		model.ProductLangs = append(model.ProductLangs, productLang)
	case ProductLangRemoved:
		println("Projector ProductLangAdded")
		data, ok := event.Data().(*ProductLangRemovedData)
		if !ok {
			return nil, errors.New("invalid event data ProductLangAdded")
		}
		atemp := model.ProductLangs
		for i, e := range atemp {
			if e.LangCode == data.LangCode {
				model.ProductLangs = atemp[:i+copy(atemp[i:], atemp[i+1:])]
			}
		}
	case AvailabilitySet:
		data, ok := event.Data().(*AvailabilityData)
		if !ok {
			return nil, errors.New("Invalid event data for ProductLangUpdated")
		}
		model.Availability = data.Availability
	case TransportSpecificationSet:
		data, ok := event.Data().(*TransportSpecificationData)
		if !ok {
			return nil, fmt.Errorf("Invalid event %s", event.EventType())
		}
		model.TransportSpecification = data.TransportSpecification
	case TransportAdded:
		data, ok := event.Data().(*TransporterData)
		if !ok {
			return nil, fmt.Errorf("Invalid event %s for transporter %s", event.EventType(), data.Id)
		}
		for _, e := range model.TransportSpecification.Transporters {

			if len(model.TransportSpecification.Transporters) > 0 && e.Id != "" && e.Id == data.Id {
				return nil, fmt.Errorf("transport %s for aggretate  %s exist -> ", e.Name, model.EntityID())

			}
		}
		model.TransportSpecification.Transporters = append(model.TransportSpecification.Transporters, data.Transporter)
	case TransportUpdated:
		data, ok := event.Data().(*TransporterData)
		if !ok {
			return nil, fmt.Errorf("Invalid event %s for transporter %s", event.EventType(), data.Id)
		}

		existProductLang := false
		if model.TransportSpecification.Transporters == nil {
			return nil, fmt.Errorf("Transporter for %s not exist", data.Id)
		}

		for _, e := range model.TransportSpecification.Transporters {
			if e.Id == data.Id {
				existProductLang = true
			}
		}

		if !existProductLang {
			return nil, fmt.Errorf("Transporter for ID %s not exist", data.Id)
		}
		model.TransportSpecification.Transporters = append(model.TransportSpecification.Transporters, data.Transporter)
	case TransportRemoved:
		data, ok := event.Data().(*TranporterRemovedData)
		if !ok {
			return nil, fmt.Errorf("Invalid event %s for transporter %s", event.EventType(), data.transportID)
		}
		if model.TransportSpecification.Transporters == nil {
			return nil, fmt.Errorf("Invalid event %s for transporter %s", event.EventType(), data.transportID)
		}

		removedProductLang := false
		atemp := model.TransportSpecification.Transporters
		for i, e := range atemp {
			if e.Id == data.transportID {
				model.TransportSpecification.Transporters = atemp[:i+copy(atemp[i:], atemp[i+1:])]
				removedProductLang = true
			}
		}
		if !removedProductLang {
			return nil, fmt.Errorf("Transporter for ID %s not exist", data.transportID)
		}
	case PricesSpecificationSet:
		data, ok := event.Data().(*PricesSpecificationData)
		if !ok {
			return nil, fmt.Errorf("Invalid event %s", event.EventType())
		}
		model.PricesSpecification = data.PricesSpecification
	case ImageAdded:
		data, ok := event.Data().(*ImageAddedData)
		if !ok {
			return nil, fmt.Errorf("Invalid event %s for image %s", event.EventType(), data.Name)
		}
		for _, e := range model.Images {
			if len(model.Images) > 0 && e.Name != "" && e.Name == data.Name {
				return nil, fmt.Errorf("image %s for aggretate %s exist -> ", e.Name, model.EntityID())
			}
		}
		model.Images = append(model.Images, data.Image)
	case ImageUpdated:
		data, ok := event.Data().(*ImageUpdatedData)
		if !ok {
			return nil, fmt.Errorf("Invalid event %s for image %s", event.EventType(), data.Name)
		}

		var imageItem Image
		if model.Images == nil {
			return nil, fmt.Errorf("Image for name %s not exist", data.Name)
		}

		for _, e := range model.Images {
			if e.Name == data.Name {
				imageItem = e
			}
		}

		if imageItem.Name == "" {
			return nil, fmt.Errorf("Image for name %s not exist", data.Name)
		}
		model.Images = append(model.Images, Image{
			Name:        imageItem.Name,
			Description: data.Description,
			Caption:     data.Description,
			Thumbnail:   imageItem.Thumbnail,
		})
	case ImageRemoved:
		data, ok := event.Data().(*ImageRemovedData)
		if !ok {
			return nil, fmt.Errorf("Invalid event %s for image %s", event.EventType(), data.Name)
		}
		if model.Images == nil {
			return nil, fmt.Errorf("Invalid event %s for image %s", event.EventType(), data.Name)
		}

		removedImage := false
		atemp := model.Images
		for i, e := range atemp {
			if e.Name == data.Name {
				model.Images = atemp[:i+copy(atemp[i:], atemp[i+1:])]
				removedImage = true
			}
		}
		if !removedImage {
			return nil, fmt.Errorf("Image %s not exist", data.Name)
		}
	}

	model.Version++
	model.UpdatedAt = TimeNow()
	return model, nil
}
