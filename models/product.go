package models

type Product struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type CreateProduct struct {
	Name    string `json:"name"`
	Surname string `json:"urname"`
}

type UpdateProduct struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"urname"`
}

type GetListRequestProduct struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListResponseProduct struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}
