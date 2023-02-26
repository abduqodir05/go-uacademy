package storage

import (
	"app/models"
)

type StorageI interface {
	CloseDb()
	User() UserRepoI
	Product() ProductRepoI
	ShopCart() ShopCartRepoI
	Commission() CommissionRepoI
	Category() CategoryRepoI
	Branch() BranchRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (string, error)
	Delete(*models.UserPrimaryKey) error
	Update(*models.UpdateUser, string) error
	GetByID(*models.UserPrimaryKey) (models.User, error)
	GetAll(*models.GetListRequest) (models.GetListResponse, error)
}

type ProductRepoI interface {
	Create(*models.CreateProduct) (string, error)
	GetByID(*models.ProductPrimaryKey) (models.ProductWithCategory, error)
	GetAll(*models.GetListProductRequest) (models.GetListProduct, error)
	Update(*models.UpdateProduct, string) error
	Delete(*models.ProductPrimaryKey) error
}

type ShopCartRepoI interface {
	Read() ([]models.ShopCart, error)
	GetByProductId(req *models.ProductPrimaryKey) ([]models.ShopCart, error)
	GetUnpaidShopCarts(req *models.UserPrimaryKey) ([]models.ShopCart, error)
	AddShopCart(*models.Add) (string, error)
	RemoveShopCart(*models.Remove) error
	GetUserShopCart(*models.UserPrimaryKey) ([]models.ShopCart, error)
	UpdateShopCart(string) error
	GetAllShopCarts() ([]models.ShopCart, error)
}

type CommissionRepoI interface {
	AddCommission(*models.Commission) error
}

type CategoryRepoI interface {
	Create(*models.CreateCategory) (string, error)
	GetByID(*models.CategoryPrimaryKey) (models.Category, error)
	GetAll(*models.GetListCategoryRequest) (models.GetListCategoryResponse, error)
	Update(*models.UpdateCategory, string) error
	Delete(*models.CategoryPrimaryKey) error
}
type BranchRepoI interface {
	CreateBranch(req *models.BranchReq) (id string, err error)
	GetList(req *models.GetBranchListRequest) (*models.GetBranchListResponse, error)
	GetBranchById(req *models.BranchPrimaryKey) (models.Branch, error)
	UpdateBranch(req *models.Branch) (models.Branch, error)
	DeleteBranch(req *models.BranchPrimaryKey) (models.Branch, error)
}
