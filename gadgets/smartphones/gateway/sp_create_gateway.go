package gateway

import (
	"course-phones-review/gadgets/smartphones/models"
	"course-phones-review/internal/database"
)

type SmartphoneCreateGateway interface {
	Create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error)
	GetSmartphoneByID(smartphoneID int64) *models.Smartphone
}

type SmartphoneCreateGtw struct {
	SmartphoneStorageGateway
}

func NewSmartphoneCreateGateway(client *database.MySqlClient) SmartphoneCreateGateway {
	return &SmartphoneCreateGtw{&SmartphoneStorage{client}}
}
