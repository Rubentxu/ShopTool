package products

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// NewProductRepository for mongo
func NewProductRepository() (p ProductRepository, err error) {
	var repo ProductRepository
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:   []string{"127.0.0.1"},
		Timeout: 60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[MongoDB Session]: %s\n", err)
		return repo, err
	}
	collection := session.DB("shoptool").C("products")
	collection.RemoveAll(nil)
	repo = ProductRepository{
		C: collection,
	}
	return repo, nil
}

// ProductRepository collection mongo
type ProductRepository struct {
	C *mgo.Collection
}

// Create one product into mongo collection
func (repo ProductRepository) Create(p *Product) error {
	p.ID = bson.NewObjectId()
	err := repo.C.Insert(p)
	return err
}

// // Update modifies an existeing value of a collection
// func(repo ProductRepository) Update(p *Product) error {
// 	err := repo.C.Update(bson.M{"_id": b.ID},
// 		bson.M{"$set": bson.M{
// 			"name": p.N
// 		}}
// 	)
// }

// Delete removes an existing value from the collection.
func (repo ProductRepository) Delete(id string) error {
	err := repo.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// GetAll returns all documents from the collection.
func (repo ProductRepository) GetAll() []Product {
	var b []Product
	iter := repo.C.Find(nil).Sort("dateadd").Iter()
	result := Product{}
	for iter.Next(&result) {
		b = append(b, result)
	}
	return b
}

// GetByID returns single document from the collection.
func (repo ProductRepository) GetByID(id string) (Product, error) {
	var b Product
	err := repo.C.FindId(bson.ObjectIdHex(id)).One(&b)
	return b, err
}

// GetLangCode returns all documents from the collection filtering by LangCode.
func (repo ProductRepository) GetLangCode(langcode []string) []Product {
	var b []Product
	iter := repo.C.Find(bson.M{"productslang.langcode": bson.M{"$in": langcode}}).Sort("dateadd").Iter()
	result := Product{}
	for iter.Next(&result) {
		b = append(b, result)
	}
	return b
}
