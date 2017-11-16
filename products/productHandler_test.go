package products

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	//"github.com/looplab/eventhorizon/repo/mongodb"

	"golang.org/x/net/context"

	//eh "github.com/looplab/eventhorizon"

	"ShopTool/products/domain"
	"encoding/json"
	"github.com/looplab/eventhorizon/repo/mongodb"

)

//func TestStaticFiles(t *testing.T) {
//	h, err := NewHandler("localhost:27017")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	r := httptest.NewRequest("GET", "/api/product/docs/", nil)
//	w := httptest.NewRecorder()
//	h.ServeHTTP(w, r)
//	if w.Code != http.StatusOK {
//		t.Error(err)
//		println(w.Code)
//	}
//}

func TestGetAll(t *testing.T) {
	domain.TimeNow = func() time.Time {
		return time.Date(2017, time.July, 10, 23, 0, 0, 0, time.Local)
	}

	h, _ := NewHandler("localhost:27017")

	repo, ok := h.Repo.Parent().(*mongodb.Repo)
	if !ok {
		t.Fatal("incorrect repo type")
	}
	println("contexto")
	println(context.Background())
	println("contexto2")
	if err := repo.Clear(context.Background()); err != nil {
		t.Log("could not clear DB:", err)
	}

	r := httptest.NewRequest("GET", "/api/product/", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Error("the status should be correct:", w.Code)
	}
	if string(w.Body.Bytes()) != `[]` {
		t.Error("the body should be correct:", string(w.Body.Bytes()))
	}

	if err := h.CommandHandler.HandleCommand(context.Background(), &domain.Create{
		Reference: "12345",
		Ean13:     "123",
		Isbn:      "331",
		Upc:       "11",
	}); err != nil {
		t.Error("there should be no error:", err)
	}

	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Error("the status should be correct:", w.Code)
	}

	var product []domain.Product
	if err := json.Unmarshal(w.Body.Bytes(), &product); err != nil {
		t.Error(err)
	}

	if product[0].Version != 1 ||
		product[0].Reference != "12345" ||
		product[0].Upc != "11" ||
		product[0].Isbn != "331" ||
		product[0].Ean13 != "123" {
		t.Error("the body should be correct:", string(w.Body.Bytes()))
	}

	if err := h.CommandHandler.HandleCommand(context.Background(), &domain.AddProductLang{
		ProductLang: domain.ProductLang{
			Name:             "Producto molon",
			Description:      "Producto molon descripcion",
			DescriptionShort: "",
			LinkRewrite:      "",
			MetaDescription:  "",
			MetaKeywords:     "",
			MetaTitle:        "",
			AvailableNow:     "",
			AvailableLater:   "",
			LangCode:         "es-Es",
		},
		ProductID: product[0].ID,
	}); err != nil {
		t.Error("there should be no error:", err)
	}

	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Error("the status should be correct:", w.Code)
	}

	var product2 []domain.Product
	if err := json.Unmarshal(w.Body.Bytes(), &product2); err != nil {
		t.Error(err)
	}
	if product, lang := product2[0], product2[0].ProductLangs[0] ; product.Version != 2 ||
		lang.Name != "Producto molon" ||
		lang.Description != "Producto molon descripcion" ||
		lang.LangCode != "es-Es" {
		t.Error("the body should be correct:", string(w.Body.Bytes()))

	}

	if err := h.CommandHandler.HandleCommand(context.Background(), &domain.AddProductLang{
		ProductLang: domain.ProductLang{
			Name:             "Producto molon",
			Description:      "Producto molon descripcion",
			DescriptionShort: "",
			LinkRewrite:      "",
			MetaDescription:  "",
			MetaKeywords:     "",
			MetaTitle:        "",
			AvailableNow:     "",
			AvailableLater:   "",
			LangCode:         "es-Es",
		},
		ProductID: product[0].ID,
	}); err == nil {
		t.Error("there should be no error:", err)
	}

	w = httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Error("the status should be correct:", w.Code)
	}

	var product3 []domain.Product
	if err := json.Unmarshal(w.Body.Bytes(), &product3); err != nil {
		t.Error(err)
	}
	if product, lang := product3[0], product3[0].ProductLangs[0] ; product.Version != 2 ||
		lang.Name != "Producto molon" ||
		lang.Description != "Producto molon descripcion" ||
		lang.LangCode != "es-Es" {
		t.Error("the body should be correct:", string(w.Body.Bytes()))

	}

}
