package api

type Product struct {
	Id        string
	Name      string
	PhotoLink string
	Price     int64 // 分
	ValidFrom string
	ValidFrom string
}

type User struct {
	Name       string
	Email      string
	AvatarLink string
}

type Order struct {
	Date     string
	Products *Product
	User     []*User
	Count    int
}
