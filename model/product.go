package model

import (
	"github.com/sunfmin/batchbuy/api"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"sort"
	"time"
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

func (product *Product) ToApi() (apiProduct *api.Product) {
	apiProduct = &api.Product{}
	apiProduct.Id = product.Id.Hex()
	apiProduct.Name = product.Name
	apiProduct.PhotoLink = product.PhotoLink
	apiProduct.Price = product.Price
	apiProduct.ValidFrom = product.ValidFrom.String()
	apiProduct.ValidTo = product.ValidTo.String()

	return
}

func availableCond(date time.Time) bson.M {
	emptyDate, err := time.Parse("2006-01-02", "0001-01-01")
	if err != nil {
		return nil
	}

	return bson.M{
		"$or": []M{
			{"validfrom": emptyDate, "validto": emptyDate},
			{"validfrom": M{"$lte": date}, "validto": emptyDate},
			{"validfrom": emptyDate, "validto": M{"$gte": date}},
			{"validfrom": M{"$lte": date}, "validto": M{"$gte": date}},
		},
	}
}

func Top3PopularProducts(date time.Time) (products []*Product, err error) {
	return PopularProductsFinder(nil, date)
}

func MyTop3FavouriteProducts(email string, date time.Time) (products []*Product, err error) {
	return PopularProductsFinder(bson.M{"userid": email}, date)
}

type PopularProductInfoDetails struct {
	Count      int
	OrderCount int
}

type PopularProductInfo struct {
	ProductId string `_id`
	Value     PopularProductInfoDetails
}

type PopularProductInfoSorter struct {
	Infos []*PopularProductInfo
}

func (this *PopularProductInfoSorter) Len() int {
	return len(this.Infos)
}

func (this *PopularProductInfoSorter) Swap(i, j int) {
	this.Infos[i], this.Infos[j] = this.Infos[j], this.Infos[i]
}

func (this *PopularProductInfoSorter) Less(i, j int) bool {
	return this.Infos[i].Value.OrderCount > this.Infos[j].Value.OrderCount
}

func PopularProductsFinder(query bson.M, date time.Time) (products []*Product, err error) {
	ppInfos := []*PopularProductInfo{}
	_, err = orderCol.Find(query).MapReduce(&mgo.MapReduce{
		Map:    `
			function() {
				emit(this.productid, {ordercount: 1, count: this.count});
			}`,
		Reduce: `
			function(key, details) {
				var reducedVal = {ordercount: 0, count: 0 };
				for (var i = 0; i < details.length; i++) {
					reducedVal.count += details[i].count;
					reducedVal.ordercount += details[i].ordercount;
				}
				return reducedVal;
			}
		`,
	}, &ppInfos)
	if err != nil {
		return
	}

	productIds := []bson.ObjectId{}
	if len(ppInfos) > 3 {
		sorter := PopularProductInfoSorter{ppInfos}
		sort.Sort(&sorter)
		ppInfos = sorter.Infos
		for _, ppInfo := range ppInfos[:3] {
			productIds = append(productIds, bson.ObjectIdHex(ppInfo.ProductId))
		}
	} else {
		for _, ppInfo := range ppInfos {
			productIds = append(productIds, bson.ObjectIdHex(ppInfo.ProductId))
		}
	}

	pQuery := availableCond(date)
	pQuery["_id"] = bson.M{"$in": productIds}
	err = productCol.Find(pQuery).All(&products)

	return
}
