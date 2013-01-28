package model

import (
	"time"
	"github.com/sunfmin/batchbuy/api"
	"labix.org/v2/mgo/bson"
	// "fmt"
)

var orderTN = "orders"
var orderCol = db.C(orderTN)

type Order struct {
	Id bson.ObjectId "_id"
	Date      time.Time
	ProductId string  // => Product.Id.Hex()
	UserId    string  // => User.Email
	Count     int
}

type OrderInput struct {
	Date      time.Time
	ProductId string  // => Product.Id.Hex()
	UserId    string  // => User.Email
	Count     int
}

func (order *Order) Put(date time.Time, email string, input OrderInput) (err error) {
	// Can't use Upsert here.
	conds := M{"userid": email, "date": getDayRangeCond(date), "productid": input.ProductId}
	count, err := orderCol.Find(conds).Count()
	
	if count == 0 {
		orderCol.Insert(input)
	} else {
		orderCol.Update(conds, &input)
	}
	
	orderCol.Find(conds).One(order)
	
	return
}

func (order Order) Remove(date time.Time, email string) (err error) {
	conds := M{"userid": email, "date": getDayRangeCond(date)}
	err = orderCol.Remove(conds)
	return
}

func (order Order) GetProduct() (product *Product) {
	product = &Product{}
	productCol.FindId(bson.ObjectIdHex(order.ProductId)).One(product)
	
	return
}

func (order Order) GetUser() (user *User) {
	user = &User{}
	userCol.Find(M{"email": order.UserId}).One(user)
	
	return
}

func (order Order) ToApi() (apiOrder *api.Order) {
	apiOrder = &api.Order{}
	apiOrder.Count = order.Count
	apiOrder.Date = order.Date.String()
	apiOrder.Users = append(apiOrder.Users, order.GetUser().ToApi())
	apiOrder.Product = order.GetProduct().ToApi()
	
	return
}

func getDayRangeCond(date time.Time) M {
	gteDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	lteDate := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())

	return M{"$gte": gteDate, "$lte": lteDate}
}

func OrderListOfDate(date time.Time) (orders []Order, err error) {
	err = orderCol.Find(M{"date": getDayRangeCond(date)}).All(&orders)
	return
}

// Generate api.Order data from model.Order data
func OrderListOfDateForApi(date time.Time) (apiOrders []*api.Order) {
	orders, _ := OrderListOfDate(date)
	
	var newOrderf bool
	for _, order := range orders {
		newOrderf = true
		for _, apiOrder := range apiOrders {
			if apiOrder.Product.Id == order.ProductId {
				apiOrder.Users = append(apiOrder.Users, order.GetUser().ToApi())
				apiOrder.Count += order.Count

				newOrderf = false
				continue
			}
		}
		if newOrderf {
			apiOrders = append(apiOrders, order.ToApi())
		}
	}
	
	return
}

// func getTime(dateStr string) (date time.Time) {
// 	date, _ = time.Parse(time.RFC3339, dateStr)
// 	return 
// }