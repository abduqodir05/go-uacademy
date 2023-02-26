package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"

	"github.com/google/uuid"
)

type shopCartRepo struct {
	fileName string
}

func NewShopCartRepo(fileName string) *shopCartRepo {
	return &shopCartRepo{
		fileName: fileName,
	}
}

func (s *shopCartRepo) Read() ([]models.ShopCart, error) {
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return []models.ShopCart{}, err
	}

	var shopCarts []models.ShopCart
	err = json.Unmarshal(data, &shopCarts)
	if err != nil {
		return []models.ShopCart{}, err
	}
	return shopCarts, nil
}

func (s *shopCartRepo) AddShopCart(req *models.Add) (string, error) {
	shopCarts, err := s.Read()
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()
	shopCarts = append(shopCarts, models.ShopCart{
		Id:        uuid,
		ProductId: req.ProductId,
		UserId:    req.UserId,
		Count:     req.Count,
		Status:    false,
	})

	body, err := json.MarshalIndent(shopCarts, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (s *shopCartRepo) RemoveShopCart(req *models.Remove) error {
	shopCarts, err := s.Read()
	if err != nil {
		return err
	}

	flag := true
	for i, v := range shopCarts {
		if req.ProductId == v.ProductId && req.UserId == v.UserId {
			shopCarts = append(shopCarts[:i], shopCarts[i+1:]...)
			flag = false
			break
		}
	}

	if flag {
		return errors.New("UserId or ProductId doesn't exist")
	}

	body, err := json.MarshalIndent(shopCarts, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (s *shopCartRepo) GetUnpaidShopCarts(req *models.UserPrimaryKey) ([]models.ShopCart, error) {
	shopcarts, err := s.Read()
	if err != nil {
		return []models.ShopCart{}, err
	}
	newS := []models.ShopCart{}
	for _, v := range shopcarts {
		if v.UserId == req.Id && v.Status == false && v.Count > 0 {
			newS = append(newS, v)
		}
	}
	return newS, nil
}

func sortShopcarts(items []models.ShopCart) {
	sort.Slice(items, func(i, j int) bool {
		LAYOUT := "2006-01-02 15:04:05"
		tm1, _ := time.Parse(LAYOUT, items[i].Time)
		tm2, _ := time.Parse(LAYOUT, items[j].Time)
		return tm1.After(tm2)
	})
}
func (s *shopCartRepo) GetAllShopCarts() ([]models.ShopCart, error) {
	shopCarts, err := s.Read()
	if err != nil {
		return []models.ShopCart{}, err
	}
	return shopCarts, nil
}

func (s *shopCartRepo) GetUserShopCart(req *models.UserPrimaryKey) ([]models.ShopCart, error) {
	shopCarts, err := s.Read()
	if err != nil {
		return []models.ShopCart{}, err
	}
	fmt.Println(req.Id)
	userShopCarts := []models.ShopCart{}
	for _, v := range shopCarts {
		if v.UserId == req.Id && v.Status == false {
			// fmt.Println(v.UserId)
			userShopCarts = append(userShopCarts, v)

		}
	}
	// fmt.Println(userShopCarts)
	// if len(userShopCarts) == 0 {
	// 	return []models.ShopCart{}, errors.New("There are no unpaid products")
	// }

	return userShopCarts, nil
}
func (s *shopCartRepo) GetByProductId(req *models.ProductPrimaryKey) ([]models.ShopCart, error) {
	shopcarts, err := s.Read()
	if err != nil {
		return []models.ShopCart{}, err
	}
	newS := []models.ShopCart{}
	for _, v := range shopcarts {
		if req.Id == v.ProductId  {
			newS = append(newS, v)
		}
	}
	return newS, nil
}

func (s *shopCartRepo) UpdateShopCart(userId string) error {
	shopCarts, err := s.Read()
	if err != nil {
		return err
	}

	for i, v := range shopCarts {
		if v.UserId == userId && v.Status == false {
			shopCarts[i].Status = true
		}
	}

	body, err := json.MarshalIndent(shopCarts, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
