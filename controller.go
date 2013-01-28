package main

import (
	"github.com/sunfmin/batchbuy/model"
	"github.com/sunfmin/batchbuy/api"
	"time"
	"fmt"
)

type Controller struct {}

func (Controller) PutProduct(id string, input api.ProductInput) (product *api.Product, err error) {
	modelProductInput := model.ProductInput{
		Name: input.Name,
		PhotoLink: input.PhotoLink,
		Price: input.Price,
		ValidFrom: stringToTime(input.ValidFrom),
		ValidTo: stringToTime(input.ValidTo),
	}
	modelProduct := model.Product{}
	modelProduct.Put(id, modelProductInput)
	
	product = modelProduct.ToApi()
	
	return
}

const timeFmt = "2006-01-02"

func stringToTime(str string) time.Time {
	date, _ := time.Parse(timeFmt, str)
	return date
}

func (Controller) RemoveProduct(id string) (err error) {
	return
}

func (Controller) PutUser(email string, input api.UserInput) (user *api.User, err error) {
	userModel := model.User{}
	modelUserInput := model.UserInput{
		Name: input.Name,
		Email: input.Email,
		AvatarLink: input.AvatarLink,
	}
	fmt.Printf("%s\n", modelUserInput)
	userModel.Put(email, modelUserInput)
	
	user = userModel.ToApi()
	return
}

func (Controller) RemoveUser(email string) (err error) {
	return
}

func (Controller) PutOrder(date string, email string, productId string, count int) (order *api.Order, err error) {
	// order = &api.Order{}
	dateD := stringToTime(date)
	orderInput := model.OrderInput{dateD, productId, email, count}
	modelOrder := model.Order{}
	modelOrder.Put(dateD, email, orderInput)
	order = modelOrder.ToApi()
	
	return
}

func (Controller) RemoveOrder(date string, email string) {
	return
}

func (Controller) AllProducts() (products []*api.Product, err error) {
	products, err = model.AllProductsForApi()
	return
}

func (Controller) ProductListOfDate(date string) (products []*api.Product, err error) {
	products, err = model.ProductListOfDateForApi(stringToTime(date))
	
	return
}

func (Controller) OrderListOfDate(date string) (orders []*api.Order, err error) {
	orders = model.OrderListOfDateForApi(stringToTime(date))
	
	return
}