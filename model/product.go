package model

import (
	// "strconv"
	"time"
	// "strings"
	"github.com/sunfmin/batchbuy/api"
	"labix.org/v2/mgo/bson"
)

var productTN = "products"
var productCol = db.C(productTN)

type Product struct {
	Id        bson.ObjectId "_id"
	Name      string
	PhotoLink string
	Price     int64
	ValidFrom time.Time
	ValidTo   time.Time
}

type ProductInput struct {
	Name      string
	PhotoLink string
	Price     int64
	ValidFrom time.Time
	ValidTo   time.Time
}

func (product *Product) Put(id string, input ProductInput) (err error) {
	if !isObjectIdHex(id) {
		id = bson.NewObjectId().Hex()
	}
	changeInfo, err := productCol.UpsertId(bson.ObjectIdHex(id), input)
	if err != nil {
		return
	}

	err = productCol.FindId(changeInfo.UpsertedId).One(product)

	return
}

func (Product) Remove(id string) (err error) {
	if isObjectIdHex(id) {
		err = productCol.RemoveId(bson.ObjectIdHex(id))
	}

	return
}

func ProductListOfDate(date time.Time) (product []Product, err error) {
	err = productCol.Find(M{"validfrom": M{"$lte": date}, "validto": M{"$gte": date}}).All(&product)
	if err != nil {
		return
	}

	otherProducts, err := unrestrainedProducts()
	product = append(product, otherProducts...)

	return
}

func unrestrainedProducts() (products []Product, err error) {
	emptyDate, err := time.Parse("2006-01-02", "0001-01-01")
	if err != nil {
		return
	}

	err = productCol.Find(M{"$or": []M{M{"validfrom": emptyDate}, {"validto": emptyDate}}}).All(&products)
	if err != nil {
		return
	}

	return
}

// TODO test it
func ProductListOfDateForApi(date time.Time) (products []*api.Product, err error) {
	modelProducts, err := ProductListOfDate(date)
	if err != nil {
		return
	}

	for _, modelProduct := range modelProducts {
		products = append(products, modelProduct.ToApi())
	}

	return
}

func AllProductsForApi() (products []*api.Product, err error) {
	modelProducts := []*Product{}
	err = productCol.Find(M{}).All(&modelProducts)
	if err != nil {
		return
	}

	for _, modelProduct := range modelProducts {
		products = append(products, modelProduct.ToApi())
	}

	return
}

func (product Product) ToApi() (apiProduct *api.Product) {
	apiProduct = &api.Product{}
	apiProduct.Id = product.Id.Hex()
	apiProduct.Name = product.Name
	apiProduct.PhotoLink = product.PhotoLink
	apiProduct.Price = product.Price
	apiProduct.ValidFrom = product.ValidFrom.String()
	apiProduct.ValidTo = product.ValidTo.String()

	return
}
