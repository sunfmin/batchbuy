package api

type Product struct {
	Id        string
	Name      string
	PhotoLink string
	Price     int64 // 分
}

type User struct {
	Name  string
	Email string
}

type Order struct {
	Date     string
	Products []*Product
	User     *User
}
