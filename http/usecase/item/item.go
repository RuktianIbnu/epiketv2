package item

import (
	"epiketv2/pkg/model"
	ir "epiketv2/pkg/repository/item"

	"errors"
)

// Usecase ...
type Usecase interface {
	Create(data *model.MsItem) (int64, error)
	GetOneByID(id int64) (*model.MsItem, error)
	UpdateOneByID(data *model.MsItem) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.MsItem, int, error)
}

type usecase struct {
	itemRepo ir.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		ir.NewRepository(),
	}
}

func (m *usecase) Create(data *model.MsItem) (int64, error) {
	return m.itemRepo.Create(data)
}

func (m *usecase) UpdateOneByID(data *model.MsItem) (int64, error) {
	rowsAffected, err := m.itemRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) (*model.MsItem, error) {
	return m.itemRepo.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.MsItem, int, error) {
	return m.itemRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.itemRepo.DeleteOneByID(id)
}
