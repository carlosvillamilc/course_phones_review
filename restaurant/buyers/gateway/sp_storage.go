package gateway

import (
	"course-phones-review/internal/database"
	"course-phones-review/internal/logs"
	"course-phones-review/restaurant/buyers/models"
)

type BuyerStorageGateway interface {
	LoadBuyers(cmd []models.Buyer) (*models.Buyer, error)
	/*GetUserByID(userID int64) *models.Buyer
	Authenticate(cmd *models.CreateBuyerCMD) (*models.Buyer, error)*/
}

type BuyerStorage struct {
	*database.MySqlClient
}

func (s *BuyerStorage) Authenticate(cmd *models.CreateBuyerCMD) (*models.Buyer, error) {

	tx, err := s.MySqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot get user data")
		return nil, err
	}

	var res models.Buyer

	err = tx.QueryRow(`select id, username, password from user
	where username = ?`, cmd.Name).Scan(&res.Id, &res.Name, &res.Age)

	if err != nil {
		logs.Log().Error("cannot execute select user statement")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &models.Buyer{
		Id:   res.Id,
		Name: res.Name,
		Age:  res.Age,
	}, nil

}

func (s *BuyerStorage) LoadBuyers(cmd []models.Buyer) (*models.Buyer, error) {

	tx, err := s.MySqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create buyer transaction")
		return nil, err
	}

	size := len(cmd)
	for value := range cmd {
		res, err := tx.Exec(`insert into buyer (id, name, age)
		values (?, ?, ?)`, cmd[value].Id, cmd[value].Name, cmd[value].Age)

		logs.Log().Debug("err ", err)
		logs.Log().Debug("res ", res)

		/*if err != nil {
			logs.Log().Error("cannot execute buyer insert statement")
			_ = tx.Rollback()
			//return nil, err
		}*/

		/*id, err := res.LastInsertId()
		logs.Log().Debug(id)

		if err != nil {
			logs.Log().Error("cannot fetch user last id")
			_ = tx.Rollback()
			return nil, err
		}*/
	}
	logs.Log().Debug("length ", size)

	_ = tx.Commit()

	return &models.Buyer{
		Id:   cmd[0].Id,
		Name: cmd[0].Name,
		Age:  cmd[0].Age,
	}, nil

}

func (s *BuyerStorage) GetUserByID(userID int64) *models.Buyer {
	tx, err := s.Begin()

	logs.Log().Debug("userID: ", userID)

	if err != nil {
		logs.Log().Error(err.Error())
		return nil
	}

	var res models.Buyer
	err = tx.QueryRow(`select id, username, password from user
	where id = ?`, userID).Scan(&res.Id, &res.Name, &res.Age)

	if err != nil {
		logs.Log().Error(err.Error())
		_ = tx.Rollback()
		return nil
	}
	_ = tx.Commit()

	return &res

}
