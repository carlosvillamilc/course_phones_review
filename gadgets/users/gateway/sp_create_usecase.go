package gateway

import "course-phones-review/gadgets/users/models"

func (s *UserCreateGtw) Create(cmd *models.CreateUserCMD) (*models.User, error) {
	return s.create(cmd)
}
