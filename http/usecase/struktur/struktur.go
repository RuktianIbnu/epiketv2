package struktur

import (
	"epiketv2/pkg/model"
	sr "epiketv2/pkg/repository/struktur"
	"errors"
	"fmt"
)

// Usecase ...
type Usecase interface {
	Create(data *model.MsStruktur) (int64, error)
	GetOneByID(id int64) (*model.MsStruktur, error)
	UpdateOneByID(data *model.MsStruktur) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.MsStruktur, int, error)
}

type usecase struct {
	struktur sr.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		sr.NewRepository(),
	}
}

func (m *usecase) Create(data *model.MsStruktur) (int64, error) {

	namaStrukturIsActive := m.struktur.CheckNamaIsExist(data.Nama_struktur)
	if namaStrukturIsActive != true {
		return 500, fmt.Errorf("Nama Sub Direktorat Sudah Ada")
	}
	println(namaStrukturIsActive)

	return m.struktur.Create(data)
}

func (m *usecase) UpdateOneByID(data *model.MsStruktur) (int64, error) {
	rowsAffected, err := m.struktur.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) (*model.MsStruktur, error) {
	return m.struktur.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.MsStruktur, int, error) {
	return m.struktur.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.struktur.DeleteOneByID(id)
}
