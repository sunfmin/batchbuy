package api

type Service interface {
	PutProduct(id string, input ProductInput) (product *Product, err error)
	RemoveProduct(id string) (err error)
	PutUser(email string, input UserInput) (user *User, err error)
	RemoveUser(email string) (err error)

	// currently, one product one order support only. multiple products in one order have to be solved by adding multiple orders.
	// PutOrder(date string, email string, productIds []string) (order *Order, err error)
	PutOrder(date string, email string, productId string, count int) (order *Order, err error)
	RemoveOrder(date string, email string)

	ProductListOfDate(date string) (products []*Product, err error)
	OrderListOfDate(date string) (orders []*Order, err error)
}