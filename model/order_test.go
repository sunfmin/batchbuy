package model

import (
	"testing"
	"fmt"
	"time"
	// "labix.org/v2/mgo/bson"
)

func TestCreateOrder(t *testing.T) {
	emptyDb()
	product := Product{}
	product.Put("new product", ProductInput{"pudding", "link to pudding", 100, time.Now().AddDate(0, 0, -10), time.Now().AddDate(0, 0, 10)})
	user := User{}
	user.Put("non-exited user", UserInput{"test", "test@test.com", "link to avatar"})
	
	orderInput := OrderInput{time.Now(), product.Id.Hex(), user.Email, 2}
	order := Order{}
	order.Put(time.Now(), "test@test.com", orderInput)
	
	result := Order{}
	orderCol.Find(M{}).One(&result)
	
	if order != result {
		fmt.Printf("Result: %s\n", result)
		fmt.Printf("Order: %s\n", order)
		t.Errorf("Can't Create New Order")
	}
	
	emptyDb()
}

func TestUpdateOrder(t *testing.T) {
	emptyDb()
	product := Product{}
	product.Put("new product", ProductInput{"pudding", "link to pudding", 100, time.Now().AddDate(0, 0, -10), time.Now().AddDate(0, 0, 10)})
	user := User{}
	user.Put("non-exited user", UserInput{"test", "test@test.com", "link to avatar"})
	
	orderInput := OrderInput{time.Now(), product.Id.Hex(), user.Email, 2}
	order := Order{}
	order.Put(time.Now(), "test@test.com", orderInput)
	
	orderInput.Count = 3
	order.Put(time.Now(), "test@test.com", orderInput)
	
	result := Order{}
	orderCol.FindId(order.Id).One(&result)
	
	if order != result && result.Count != 3 {
		t.Errorf("Can't Update Order")
		fmt.Printf("Result: %s\nOrder: %s\n", result, order.Id)
	}
	
	emptyDb()
}

func TestOrderListOfDate(t *testing.T) {
	emptyDb()
	product := Product{}
	product.Put("new product", ProductInput{"pudding", "link to pudding", 100, time.Now().AddDate(0, 0, -10), time.Now().AddDate(0, 0, 10)})
	user1, user2, user3 := User{}, User{}, User{}
	user1.Put("non-exited user", UserInput{"test1", "test1@test.com", "link to avatar"})
	user2.Put("non-exited user", UserInput{"test2", "test2@test.com", "link to avatar"})
	user3.Put("non-exited user", UserInput{"test3", "test3@test.com", "link to avatar"})
	
	order := Order{}
	order.Put(time.Now(), user1.Email, OrderInput{time.Now(), product.Id.Hex(), user1.Email, 2})
	order.Put(time.Now(), user2.Email, OrderInput{time.Now(), product.Id.Hex(), user2.Email, 2})
	order.Put(time.Now().AddDate(0, 0, -10), user3.Email, OrderInput{time.Now().AddDate(0, 0, -10), product.Id.Hex(), user3.Email, 2})
	
	orderList, _ := OrderListOfDate(time.Now())
	if len(orderList) != 2 {
		fmt.Printf("Get OrderList: %s\n", orderList)
		t.Errorf("Can't Get Order List Properly")
	}
	
	emptyDb()
}

func emptyDb() {
	orderCol.RemoveAll(M{})
	userCol.RemoveAll(M{})
	productCol.RemoveAll(M{})
}

func TestOrderListOfDateForApi(t *testing.T) {
	emptyDb()
	product := Product{}
	product.Put("new product", ProductInput{"pudding", "link to pudding", 100, time.Now().AddDate(0, 0, -10), time.Now().AddDate(0, 0, 10)})
	user1, user2, user3 := User{}, User{}, User{}
	user1.Put("non-exited user", UserInput{"test1", "test1@test.com", "link to avatar"})
	user2.Put("non-exited user", UserInput{"test2", "test2@test.com", "link to avatar"})
	user3.Put("non-exited user", UserInput{"test3", "test3@test.com", "link to avatar"})
	
	order := Order{}
	order.Put(time.Now(), user1.Email, OrderInput{time.Now(), product.Id.Hex(), user1.Email, 2})
	order.Put(time.Now(), user2.Email, OrderInput{time.Now(), product.Id.Hex(), user2.Email, 2})
	order.Put(time.Now().AddDate(0, 0, -10), user3.Email, OrderInput{time.Now().AddDate(0, 0, -10), product.Id.Hex(), user3.Email, 2})
	
	orderList := OrderListOfDateForApi(time.Now())
	if len(orderList) != 1 || orderList[0].Count != 4 {
		fmt.Printf("Get OrderList: %s\n", orderList)
		t.Errorf("Can't Get Order List Properly")
	}
	
	emptyDb()
}