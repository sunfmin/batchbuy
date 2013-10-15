package services

import (
	"fmt"
	"github.com/sunfmin/batchbuy/api"
	"github.com/sunfmin/batchbuy/model"
	"time"
)

type ServiceImpl struct {
}

var DefaultService = ServiceImpl{}

func (serv ServiceImpl) PutProduct(id string, input api.ProductInput) (product *api.Product, err error) {
	if input.ValidFrom == "" {
		input.ValidFrom = "0001-01-01"
	}
	if input.ValidTo == "" {
		input.ValidTo = "0001-01-01"
	}

	validFromT, err := StringToTime(input.ValidFrom)
	if err != nil {
		return
	}
	validToT, err := StringToTime(input.ValidTo)
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

const TimeFmt = "2006-01-02"

func StringToTime(str string) (date time.Time, err error) {
	date, err = time.Parse(TimeFmt, str)
	return
}

func (serv ServiceImpl) RemoveProduct(id string) (err error) {
	return
}

func (serv ServiceImpl) PutUser(email string, input api.UserInput) (user *api.User, err error) {
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

func (serv ServiceImpl) RemoveUser(email string) (err error) {
	err = model.RemoveUser(email)

	return
}

func (serv ServiceImpl) GetAllUsers() (users []*api.User, err error) {
	users, err = model.GetAllUsersInApi()

	return
}

func (serv ServiceImpl) PutOrder(date string, email string, productId string, count int) (order *api.Order, err error) {
	dateD, err := StringToTime(date)
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

func (serv ServiceImpl) RemoveOrder(date string, email string, productId string) (err error) {
	dateD, err := StringToTime(date)
	if err != nil {
		return
	}

	err = model.RemoveOrder(dateD, email, productId)

	return
}

func (serv ServiceImpl) AllProducts() (products []*api.Product, err error) {
	products, err = model.AllProductsForApi()
	return
}

func (serv ServiceImpl) MyAvaliableProducts(date string, email string) (products []*api.Product, err error) {
	u := model.User{Email: email}
	var prods []model.Product
	d, err := StringToTime(date)
	prods, err = u.AvaliableProducts(d)
	for _, p := range prods {
		products = append(products, p.ToApi())
	}
	return
}

func (serv ServiceImpl) MyOrders(date string, email string) (orders []*api.Order, err error) {
	u := model.User{Email: email}

	d, err := StringToTime(date)
	orders, err = u.OrdersForApi(d)
	return
}

func (serv ServiceImpl) ProductListOfDate(date string) (products []*api.Product, err error) {
	dateT, err := StringToTime(date)
	if err != nil {
		return
	}
	products, err = model.ProductListOfDateForApi(dateT)

	return
}

func (serv ServiceImpl) OrderListOfDate(date string) (orders []*api.Order, err error) {
	dateT, err := StringToTime(date)
	if err != nil {
		return
	}
	orders, err = model.OrderListOfDateForApi(dateT)

	return
}
