package model

import (
	"github.com/sunfmin/batchbuy/api"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
	// "fmt"
)

var orderTN = "orders"
var orderCol = db.C(orderTN)

type Order struct {
	Id        bson.ObjectId "_id"
	Date      time.Time
	ProductId string // => Product.Id.Hex()
	UserId    string // => User.Email
	Count     int
}

type OrderInput struct {
	Date      time.Time
	ProductId string // => Product.Id.Hex()
	UserId    string // => User.Email
	Count     int
}

func (order *Order) Put(date time.Time, email string, input OrderInput) (err error) {
	// Can't use Upsert here.
	conds := M{"userid": email, "date": getDayRangeCond(date), "productid": input.ProductId}
	count, err := orderCol.Find(conds).Count()
	if err != nil {
		return
	}

	if count == 0 {
		err = orderCol.Insert(input)
	} else {
		err = orderCol.Update(conds, &input)
	}
	if err != nil {
		return
	}

	err = orderCol.Find(conds).One(order)
	if err != nil {
		return
	}

	return
}

func RemoveOrder(date time.Time, email string, productId string) (err error) {
	err = orderCol.Remove(M{"userid": email, "date": getDayRangeCond(date), "productid": productId})
	return
}

func (order Order) GetProduct() (product *Product) {
	product = &Product{}
	productCol.FindId(bson.ObjectIdHex(order.ProductId)).One(product)

	return
}

func (order Order) GetUser() (user *User, err error) {
	user = &User{}
	err = userCol.Find(M{"email": order.UserId}).One(user)

	return
}

func (order Order) ToApi() (apiOrder *api.Order) {
	apiOrder = &api.Order{}
	apiOrder.Count = order.Count
	apiOrder.Date = order.Date.String()
	user, err := order.GetUser()
	if err != nil {
		panic(err)
	}
	apiOrder.Users = append(apiOrder.Users, user.ToApi())
	apiOrder.Product = order.GetProduct().ToApi()

	return
}

func getDayRangeCond(date time.Time) M {
	gteDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	lteDate := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())

	return M{"$gte": gteDate, "$lte": lteDate}
}

func OrderListOfDate(date time.Time) (orders []Order, err error) {
	allOrders := []Order{}
	err = orderCol.Find(M{"date": getDayRangeCond(date)}).Sort("productid").All(&allOrders)
	if err != nil {
		return
	}

	for _, order := range allOrders {
		_, err := order.GetUser()
		if err != mgo.ErrNotFound {
			orders = append(orders, order)
		} else {
			continue
		}
	}

	return
}

// Generate api.Order data from model.Order data
func OrderListOfDateForApi(date time.Time) (apiOrders []*api.Order, err error) {
	orders, err := OrderListOfDate(date)
	if err != nil {
		return
	}

	apiOrders = ordersToApi(orders)

	return
}

// weird error here: can't declare this func in this signature:
// func ordersToApi(orders []Order) (apiOrders []*api.Order)
func ordersToApi(orders []Order) []*api.Order {
	var newOrderf bool
	apiOrders := []*api.Order{}

	for _, order := range orders {
		newOrderf = true
		for _, apiOrder := range apiOrders {
			if apiOrder.Product.Id == order.ProductId {
				user, err := order.GetUser()
				if err != nil {
					panic(err)
				}

				apiOrder.Users = append(apiOrder.Users, user.ToApi())
				apiOrder.Count += order.Count

				newOrderf = false
				continue
			}
		}
		if newOrderf {
			apiOrders = append(apiOrders, order.ToApi())
		}
	}

	return apiOrders
}

func GetOrderCount(email string, productId string, date time.Time) (count int, err error) {
	order := Order{}
	err = orderCol.Find(M{"userid": email, "productid": productId, "date": getDayRangeCond(date)}).One(&order)
	if err != nil {
		return
	}
	count = order.Count

	return
}
