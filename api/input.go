package api

type ProductInput struct {
	Name      string `schema:"name"`
	Price     int64  `schema:"price"`
	PhotoLink string `schema:"photolink"`
	ValidFrom string `schema:"validfrom"`
	ValidTo   string `schema:"validto"`
}

type UserInput struct {
	Name       string `schema:"name"`
	Email      string `schema:"email"`
	AvatarLink string `schema:"avatarlink"`
}
