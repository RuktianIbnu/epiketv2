package data_center

import (
	"epiketv2/pkg/model"
	dcr "epiketv2/pkg/repository/data_center"

	"errors"
)

// Usecase ...
type Usecase interface {
	Create(data *model.MsDataCenter) error
	GetOneByID(id int64) (*model.MsDataCenter, error)
	UpdateOneByID(data *model.MsDataCenter) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.MsDataCenter, int, error)
}

type usecase struct {
	datacenterRepo dcr.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		dcr.NewRepository(),
	}
}

func (m *usecase) Create(data *model.MsDataCenter) error {
	lastID, err := m.datacenterRepo.Create(data)
	if err != nil {
		return err
	}
	data.ID = lastID

	return nil
}

func (m *usecase) UpdateOneByID(data *model.MsDataCenter) (int64, error) {
	rowsAffected, err := m.datacenterRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) (*model.MsDataCenter, error) {
	return m.datacenterRepo.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.MsDataCenter, int, error) {
	return m.datacenterRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.datacenterRepo.DeleteOneByID(id)
}
