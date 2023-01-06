package kegiatan

import (
	"epiketv2/pkg/model"
	kr "epiketv2/pkg/repository/kegiatan"

	"errors"
)

// Usecase ...
type Usecase interface {
	Create(data *model.MsKegiatan) (int64, error)
	GetOneByID(id int64) (*model.MsKegiatan, error)
	UpdateOneByID(data *model.MsKegiatan) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.MsKegiatan, int, error)
}

type usecase struct {
	kegiatanRepo kr.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		kr.NewRepository(),
	}
}

func (m *usecase) Create(data *model.MsKegiatan) (int64, error) {
	return m.kegiatanRepo.Create(data)
}

func (m *usecase) UpdateOneByID(data *model.MsKegiatan) (int64, error) {
	rowsAffected, err := m.kegiatanRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) (*model.MsKegiatan, error) {
	return m.kegiatanRepo.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.MsKegiatan, int, error) {
	return m.kegiatanRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.kegiatanRepo.DeleteOneByID(id)
}
