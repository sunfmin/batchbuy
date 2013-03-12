package main

import (
	"fmt"
	"github.com/sunfmin/batchbuy/api"
	"github.com/sunfmin/batchbuy/model"
	"time"
)

type Controller struct{}

func (Controller) PutProduct(id string, input api.ProductInput) (product *api.Product, err error) {
	if input.ValidFrom == "" {
		input.ValidFrom = "0001-01-01"
	}
	if input.ValidTo == "" {
		input.ValidTo = "0001-01-01"
	}

	validFromT, err := stringToTime(input.ValidFrom)
	if err != nil {
		return
	}
	validToT, err := stringToTime(input.ValidTo)
	if err != nil {
		return
	}

	modelProductInput := model.ProductInput{
		Name:      input.Name,
		PhotoLink: input.PhotoLink,
		Price:     input.Price,
		ValidFrom: validFromT,
		ValidTo:   validToT,
	}
	modelProduct := model.Product{}
	err = modelProduct.Put(id, modelProductInput)
	if err != nil {
		return
	}

	product = modelProduct.ToApi()

	return
}

const timeFmt = "2006-01-02"

func stringToTime(str string) (date time.Time, err error) {
	date, err = time.Parse(timeFmt, str)
	return
}

func (Controller) RemoveProduct(id string) (err error) {
	return
}

func (Controller) PutUser(email string, input api.UserInput) (user *api.User, err error) {
	userModel := model.User{}
	modelUserInput := model.UserInput{
		Name:       input.Name,
		Email:      input.Email,
		AvatarLink: input.AvatarLink,
	}
	fmt.Printf("%s\n", modelUserInput)
	err = userModel.Put(email, modelUserInput)
	if err != nil {
		return
	}

	user = userModel.ToApi()
	return
}

func (Controller) RemoveUser(email string) (err error) {
	err = model.RemoveUser(email)

	return
}

func (Controller) GetAllUsers() (users []*api.User, err error) {
	users, err = model.GetAllUsersInApi()

	return
}

func (Controller) PutOrder(date string, email string, productId string, count int) (order *api.Order, err error) {
	dateD, err := stringToTime(date)
	if err != nil {
		return
	}

	orderInput := model.OrderInput{dateD, productId, email, count}
	modelOrder := model.Order{}
	err = modelOrder.Put(dateD, email, orderInput)
	if err != nil {
		return
	}
	order = modelOrder.ToApi()

	return
}

func (Controller) RemoveOrder(date string, email string, productId string) (err error) {
	dateD, err := stringToTime(date)
	if err != nil {
		return
	}

	err = model.RemoveOrder(dateD, email, productId)

	return
}

func (Controller) AllProducts() (products []*api.Product, err error) {
	products, err = model.AllProductsForApi()
	return
}

func (Controller) ProductListOfDate(date string) (products []*api.Product, err error) {
	dateT, err := stringToTime(date)
	if err != nil {
		return
	}
	products, err = model.ProductListOfDateForApi(dateT)

	return
}

func (Controller) OrderListOfDate(date string) (orders []*api.Order, err error) {
	dateT, err := stringToTime(date)
	if err != nil {
		return
	}
	orders, err = model.OrderListOfDateForApi(dateT)

	return
}
