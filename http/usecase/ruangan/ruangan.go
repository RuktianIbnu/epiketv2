package ruangan

import (
	"epiketv2/pkg/model"
	dcr "epiketv2/pkg/repository/data_center"
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
	GetAllById_dc(id_dc int64) ([]*model.MsRuangan, error)
}

type usecase struct {
	ruanganRepo    rr.Repository
	datacenterRepo dcr.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rr.NewRepository(),
		dcr.NewRepository(),
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
	dataRuangan, err := m.ruanganRepo.GetOneByID(id)
	if err != nil {
		return nil, err
	}

	dataDc, err := m.datacenterRepo.GetOneByID(dataRuangan.Id_dc)
	if err != nil {
		return nil, err
	}

	dataRuangan.DataCenter = dataDc

	return dataRuangan, nil
}

func (m *usecase) GetAllById_dc(id_dc int64) ([]*model.MsRuangan, error) {
	dataRuangan, err := m.ruanganRepo.GetAllByIDdataCenter(id_dc)
	if err != nil {
		return nil, err
	}

	return dataRuangan, nil
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.MsRuangan, int, error) {
	return m.ruanganRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.ruanganRepo.DeleteOneByID(id)
}
