package domain

import (
	eh "github.com/looplab/eventhorizon"
)

// Events to products
const (
	ProductCreated     = eh.EventType("product:created")
	ProductDeleted     = eh.EventType("product:deleted")
	ProductLangAdded   = eh.EventType("product:productlang_added")
	ProductLangUpdated = eh.EventType("product:productlang_updated")
	ProductLangRemove  = eh.EventType("product:productlang_remove")
)

func init() {
	eh.RegisterEventData(ProductLangAdded, func() eh.EventData {
		return &ProductLangAddedData{}
	})
	eh.RegisterEventData(ProductLangUpdated, func() eh.EventData {
		return &ProductLangUpdatedData{}
	})
	eh.RegisterEventData(ProductLangRemove, func() eh.EventData {
		return &ProductLangRemoveData{}
	})
}

// ProductLangAddedData is the event data for the ProductLangAdded
type ProductLangAddedData struct {
	Name             string `json:"name" bson:"name"`
	Description      string `json:"description" bson:"description"`
	DescriptionShort string `json:"description_short" bson:"description_short"`
	LinkRewrite      string `json:"link_rewrite" bson:"link_rewrite"`
	MetaDescription  string `json:"meta_description" bson:"meta_description"`
	MetaKeywords     string `json:"meta_keywords" bson:"meta_keywords"`
	MetaTitle        string `json:"meta_title" bson:"meta_title"`
	AvailableNow     string `json:"available_now" bson:"available_now"`
	AvailableLater   string `json:"available_later" bson:"available_later"`
	LangCode         string `json:"lang_code" bson:"lang_code"`
}

// ProductLangUpdatedData is the event data for the ProductLangUpdate
type ProductLangUpdatedData struct {
	Name             string `json:"name" bson:"name"`
	Description      string `json:"description" bson:"description"`
	DescriptionShort string `json:"description_short" bson:"description_short"`
	LinkRewrite      string `json:"link_rewrite" bson:"link_rewrite"`
	MetaDescription  string `json:"meta_description" bson:"meta_description"`
	MetaKeywords     string `json:"meta_keywords" bson:"meta_keywords"`
	MetaTitle        string `json:"meta_title" bson:"meta_title"`
	AvailableNow     string `json:"available_now" bson:"available_now"`
	AvailableLater   string `json:"available_later" bson:"available_later"`
	LangCode         string `json:"lang_code" bson:"lang_code"`
}

// ProductLangRemoveData is the event data for the ProductLangRemove
type ProductLangRemoveData struct {
	LangCode string `json:"lang_code" bson:"lang_code"`
}
