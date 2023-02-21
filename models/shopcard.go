package models

type ShopcartStruct struct {
	Product_Id      string `json:"Product_Id"`
	User_id    string `json:"User_id"`
	Count string `json:"count"`
}
type AddShopcart struct {
	Product_Id      string `json:"Product_Id"`
	User_id    string `json:"User_id"`
	Count string `json:"count"`
}
type RemoveShopcart struct {
	Product_Id      string `json:"Product_Id"`
	User_id    string `json:"User_id"`
	Count string `json:"count"`
}
