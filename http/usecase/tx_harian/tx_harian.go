package tx_harian

import (
	"epiketv2/pkg/model"
	dcr "epiketv2/pkg/repository/data_center"
	rr "epiketv2/pkg/repository/ruangan"
	thr "epiketv2/pkg/repository/tx_harian"
	ur "epiketv2/pkg/repository/user"

	"errors"
)

// Usecase ...
type Usecase interface {
	Create(data *model.TxPiketHarian) (int64, error)
	GetOneByID(id int64) (*model.TxPiketHarian, error)
	UpdateOneByID(data *model.TxPiketHarian) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.TxPiketHarian, int, error)
}

type usecase struct {
	tx_harian      thr.Repository
	datacenterRepo dcr.Repository
	ruanganRepo    rr.Repository
	userRepo       ur.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		thr.NewRepository(),
		dcr.NewRepository(),
		rr.NewRepository(),
		ur.NewRepository(),
	}
}

func (m *usecase) Create(data *model.TxPiketHarian) (int64, error) {
	return m.tx_harian.Create(data)
}

func (m *usecase) UpdateOneByID(data *model.TxPiketHarian) (int64, error) {
	rowsAffected, err := m.tx_harian.UpdateByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) (*model.TxPiketHarian, error) {
	dataTxPiket, err := m.tx_harian.GetOneByID(id)
	if err != nil {
		return nil, err
	}

	dataDc, err := m.datacenterRepo.GetOneByID(dataTxPiket.IdDataCenter)
	if err != nil {
		return nil, err
	}

	dataRuangan, err := m.ruanganRepo.GetOneByID(dataTxPiket.IdRuangan)
	if err != nil {
		return nil, err
	}

	dataUser, err := m.userRepo.GetOneByID(dataTxPiket.IdUser1)
	if err != nil {
		return nil, err
	}

	dataUser2, err := m.userRepo.GetOneByID(dataTxPiket.IdUser2)
	if err != nil {
		return nil, err
	}

	dataTxPiket.DetailDataCenter = dataDc
	dataTxPiket.DetailRuangan = dataRuangan
	dataTxPiket.DetailUser = dataUser
	dataTxPiket.DetailUserTwo = dataUser2

	return dataTxPiket, nil
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.TxPiketHarian, int, error) {
	return m.tx_harian.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.tx_harian.DeleteOneByID(id)
}
