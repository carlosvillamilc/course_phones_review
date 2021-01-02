package gateway

import (
	"course-phones-review/internal/database"
	"course-phones-review/internal/logs"
	"course-phones-review/restaurant/buyers/models"
	"strconv"
	"strings"
)

type BuyerStorageGateway interface {
	SaveBuyers(cmd []models.Buyer) (*models.Buyer, error)
	SaveProducts(cmd string) (*models.Buyer, error)
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

func (s *BuyerStorage) SaveBuyers(cmd []models.Buyer) (*models.Buyer, error) {

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

func (s *BuyerStorage) SaveProducts(cmd string) (*models.Buyer, error) {

	tx, err := s.MySqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create products transaction")
		return nil, err
	}
	logs.Log().Debug("products ", cmd)

	prodString := strings.Split(cmd, "\n")
	logs.Log().Debug("products Array ", prodString[0])
	logs.Log().Debug("products Array ", prodString[1])
	logs.Log().Debug("products Array ", prodString[2])

	//for i := range prodString {
	for i := 0; i < len(prodString)-1; i++ {

		var productModel models.Product
		product := strings.Split(string(prodString[i]), "'")

		productModel.Id = product[0]
		productModel.Name = product[1]
		productModel.Price, err = strconv.Atoi(product[2])

		logs.Log().Debug("Product ", i, productModel)

		res, err := tx.Exec(`insert into product (id, name, price)
		values (?, ?, ?)`, productModel.Id, productModel.Name, productModel.Price)

		logs.Log().Debug("err ", err)
		logs.Log().Debug("res ", res)
	}

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
	//}
	//logs.Log().Debug("length ", size)

	_ = tx.Commit()

	return &models.Buyer{
		/*Id:   cmd[0].Id,
		Name: cmd[0].Name,
		Age:  cmd[0].Age,*/
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
