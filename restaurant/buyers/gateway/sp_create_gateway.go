package gateway

import (
	"course-phones-review/internal/database"
	"course-phones-review/restaurant/buyers/models"
)

type BuyerCreateGateway interface {
	SaveBuyers(cmd []models.Buyer) (*models.Buyer, error)
	SaveProducts(cmd string) (*models.Buyer, error)
	/*GetUserByID(userID int64) *models.Buyer
	Authenticate(cmd *models.CreateBuyerCMD) (*models.Buyer, error)*/
}

type BuyerCreateGtw struct {
	BuyerStorageGateway
}

func NewBuyerCreateGateway(client *database.MySqlClient) BuyerCreateGateway {
	return &BuyerCreateGtw{&BuyerStorage{client}}
}
