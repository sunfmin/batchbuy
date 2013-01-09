package api

type Service interface {
	PutProduct(id string, input *Product) (product *Product, err error)
	RemoveProduct(id string) (err error)
	PutUser(email string, input *User) (user *User, err error)
	RemoveUser(email string) (err error)

	PutOrder(date string, email string, productIds []string) (order *Order, err error)
	RemoveOrder(date string, email string)

	OrderListOfDate(date string) (orders []*Order, err error)
}
