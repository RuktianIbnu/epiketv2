package ruangan

import (
	"epiketv2/pkg/model"
	rr "epiketv2/pkg/repository/ruangan"

	"errors"
)

// Usecase ...
type Usecase interface {
	Create(data *model.MsRuangan) (int64, error)
	GetOneByID(id int64) (*model.MsRuangan, error)
	UpdateOneByID(data *model.MsRuangan) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.MsRuangan, int, error)
}

type usecase struct {
	ruanganRepo rr.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rr.NewRepository(),
	}
}

func (m *usecase) Create(data *model.MsRuangan) (int64, error) {
	return m.ruanganRepo.Create(data)
}

func (m *usecase) UpdateOneByID(data *model.MsRuangan) (int64, error) {
	rowsAffected, err := m.ruanganRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) (*model.MsRuangan, error) {
	return m.ruanganRepo.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.MsRuangan, int, error) {
	return m.ruanganRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.ruanganRepo.DeleteOneByID(id)
}
