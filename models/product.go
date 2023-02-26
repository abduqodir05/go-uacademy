package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type Product struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Category Category `json:"category"`
}

type ProductWithCategory struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID string  `json:"category_id"`
}

type CreateProduct struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID string  `json:"category_id"`
}

type UpdateProduct struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID string  `json:"category_id"`
}

type GetListProduct struct {
	Name string `json:"name"`
	Count    int       `json:"count"`
	Products []Product `json:"products"`
}

type GetListProductRequest struct {
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	CategoryID string `json:"category_id"`
}
type CountProduct struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
type GetProductCount struct {
	CountProduct []CountProduct `json:"countProduct"`
}
type GetAllProductsList struct {
	AllProducts []Product `json:"Products"`
}