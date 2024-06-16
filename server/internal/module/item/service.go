package item

import "github.com/mrspec7er/livebid/server/internal/database"

type Service struct {
	Store database.DBConn
}

func (s *Service) Create(item *database.Item) (int, error) {
	err := s.Store.Create(&item).Error
	if err != nil {
		return 400, err
	}

	return 200, nil
}

func (s *Service) FindOne(item *database.Item, id string) (int, error) {
	err := s.Store.First(&item, "number = ?", id).Error
	if err != nil {
		return 400, err
	}

	return 200, nil
}

func (s *Service) Delete(item *database.Item, id string) (int, error) {
	err := s.Store.Delete(&item, "number = ?", id).Error
	if err != nil {
		return 400, err
	}

	return 200, nil
}
