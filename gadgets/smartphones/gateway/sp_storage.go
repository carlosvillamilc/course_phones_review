package gateway

import (
	"course-phones-review/gadgets/smartphones/models"
	"course-phones-review/internal/database"
	"course-phones-review/internal/logs"
)

type SmartphoneStorageGateway interface {
	//Add(cmd *models.CreateSmartphoneCMD)(*models.Smartphone, error)
	create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error)
	GetSmartphoneByID(smartphoneID int64) *models.Smartphone
}

type SmartphoneStorage struct {
	*database.MySqlClient
}

func (s *SmartphoneStorage) create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error) {
//func (s *SmartphoneStorage) Add(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error) {
	tx, err := s.MySqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create transaction")
		return nil, err
	}

	res, err := tx.Exec(`insert into smartphone (name, price, country_origin, operative_system) 
	values (?, ?, ?, ?)`, cmd.Name, cmd.Price, cmd.CountryOrigin, cmd.OperativeSystem)

	if err != nil {
		logs.Log().Error("cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		logs.Log().Error("cannot fetch last id")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &models.Smartphone{
		Id:            		id,
		Name:          		cmd.Name,
		Price:         		cmd.Price,
		CountryOrigin: 		cmd.CountryOrigin,
		OperativeSystem:	cmd.OperativeSystem,
	}, nil
}

func (s *SmartphoneStorage) GetSmartphoneByID(smartphoneID int64) *models.Smartphone {
	tx, err := s.Begin()	

	logs.Log().Debug("smartphoneID: ",smartphoneID)

	if err != nil {
		logs.Log().Error(err.Error())
		return nil
	}

	var res models.Smartphone	
	err = tx.QueryRow(`select id, name, price, country_origin, operative_system from smartphone 
	where id = ?`, smartphoneID).Scan(&res.Id, &res.Name, &res.Price, &res.CountryOrigin, &res.OperativeSystem)

	if err != nil {
		logs.Log().Error(err.Error())
		_ = tx.Rollback()
		return nil
	}
	_ = tx.Commit()

	return &res	
}