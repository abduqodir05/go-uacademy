package models

type ShopCartPrimaryKey struct {
	Id string `json:"id"`
}

type ShopCart struct {
	Id        string  `json:"id"`
	ProductId string  `json:"product_id"`
	UserId    string  `json:"user_id"`
	Count     int     `json:"count"`
	Status    bool    `json:"status"`
	Time      string  `json:"time"`
	Price     float64 `json:"price"`
}

type Add struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Time      string `json:"time"`
}

type Remove struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
}

type TimeFilter struct {
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`
}
type GetNameCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
type GetShopCartsList struct {
	ShopCarts []ShopCart `json:"shopCarts"`
}
type GetShopCartsNameCount struct {
	GetnameCount []GetNameCount `json:"shopCarts"`
}
