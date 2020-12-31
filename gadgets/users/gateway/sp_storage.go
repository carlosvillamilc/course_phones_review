package gateway

import (
	"course-phones-review/gadgets/users/models"
	"course-phones-review/internal/database"
	"course-phones-review/internal/logs"
)

type UserStorageGateway interface {
	create(cmd *models.CreateUserCMD) (*models.User, error)
	GetUserByID(userID int64) *models.User
	Authenticate(cmd *models.CreateUserCMD) (*models.User, error)
}

type UserStorage struct {
	*database.MySqlClient
}

func (s *UserStorage) Authenticate(cmd *models.CreateUserCMD) (*models.User, error) {

	tx, err := s.MySqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot get user data")
		return nil, err
	}

	var res models.User

	err = tx.QueryRow(`select id, username, password from user
	where username = ?`, cmd.Username).Scan(&res.Id, &res.Username, &res.Password)

	if err != nil {
		logs.Log().Error("cannot execute select user statement")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &models.User{
		Id:       res.Id,
		Username: res.Username,
		Password: res.Password,
	}, nil
}

func (s *UserStorage) create(cmd *models.CreateUserCMD) (*models.User, error) {

	tx, err := s.MySqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create user transaction")
		return nil, err
	}

	res, err := tx.Exec(`insert into user (username, password) 
	values (?, ?)`, cmd.Username, cmd.Password)

	if err != nil {
		logs.Log().Error("cannot execute user statement")
		_ = tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		logs.Log().Error("cannot fetch user last id")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &models.User{
		Id:       id,
		Username: cmd.Username,
		Password: cmd.Password,
	}, nil
}

func (s *UserStorage) GetUserByID(userID int64) *models.User {
	tx, err := s.Begin()

	logs.Log().Debug("userID: ", userID)

	if err != nil {
		logs.Log().Error(err.Error())
		return nil
	}

	var res models.User
	err = tx.QueryRow(`select id, username, password from user
	where id = ?`, userID).Scan(&res.Id, &res.Username, &res.Password)

	if err != nil {
		logs.Log().Error(err.Error())
		_ = tx.Rollback()
		return nil
	}
	_ = tx.Commit()

	return &res
}
