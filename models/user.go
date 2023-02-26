package models

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type User struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Balance float64 `json:"balance"`
}

type UpdateUser struct {
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Balance float64 `json:"balance"`
}

type CreateUser struct {
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Balance float64 `json:"balance"`
}

type GetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListResponse struct {
	Count int    `json:"count"`
	Users []User `json:"users"`
}

type UserHistory struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Total float64 `json:"total"`
	Count int     `json:"count"`
	Time  string  `json:"time"`
}

type GetUserHistoryDto struct {
	UserHistory []UserHistory `json:"userHistory"`
}
type GetUserTotalPrice struct {
	UserTotalPrice []UserHistory `json:"userTotalPrice"`
}