package api

type ProductInput struct {
	Name      string
	Price     int64
	PhotoLink string
}

type UserInput struct {
	Email string
	Name  string
}
