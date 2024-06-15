package user

import "github.com/mrspec7er/livebid/server/internal/database"

type Service struct {
	Store database.DBConn
}

func (s *Service) Create(user *database.User) (int, error) {
	err := s.Store.Create(&user).Error
	if err != nil {
		return 400, err
	}

	return 200, nil
}

func (s *Service) FindOne(user *database.User, id string) (int, error) {
	err := s.Store.First(&user, id).Error
	if err != nil {
		return 400, err
	}

	return 200, nil
}

func (s *Service) Delete(user *database.User, id string) (int, error) {
	err := s.Store.Delete(&user, id).Error
	if err != nil {
		return 400, err
	}

	return 200, nil
}
