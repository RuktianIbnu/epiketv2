package tx_piket

import (
	"epiketv2/pkg/model"
	dcr "epiketv2/pkg/repository/data_center"
	ir "epiketv2/pkg/repository/item"
	kr "epiketv2/pkg/repository/kegiatan"
	rr "epiketv2/pkg/repository/ruangan"
	tpr "epiketv2/pkg/repository/tx_piket"
	ur "epiketv2/pkg/repository/user"

	"errors"
)

// Usecase ...
type Usecase interface {
	Create(data *model.TxKegiatanPiket) (int64, error)
	GetOneByID(id int64) (*model.TxKegiatanPiket, error)
	UpdateOneByID(data *model.TxKegiatanPiket) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.TxKegiatanPiket, int, error)
}

type usecase struct {
	txpiketRepo    tpr.Repository
	kegitanRepo    kr.Repository
	datacenterRepo dcr.Repository
	ruanganRepo    rr.Repository
	itemRepo       ir.Repository
	userRepo       ur.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		tpr.NewRepository(),
		kr.NewRepository(),
		dcr.NewRepository(),
		rr.NewRepository(),
		ir.NewRepository(),
		ur.NewRepository(),
	}
}

func (m *usecase) Create(data *model.TxKegiatanPiket) (int64, error) {
	return m.txpiketRepo.Create(data)
}

func (m *usecase) UpdateOneByID(data *model.TxKegiatanPiket) (int64, error) {
	rowsAffected, err := m.txpiketRepo.UpdateByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) (*model.TxKegiatanPiket, error) {
	dataTxPiket, err := m.txpiketRepo.GetOneByID(id)
	if err != nil {
		return nil, err
	}

	dataKegiatan, err := m.kegitanRepo.GetOneByID(dataTxPiket.IdKegiatan)
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

	dataItem, err := m.itemRepo.GetOneByID(dataTxPiket.IdItem)
	if err != nil {
		return nil, err
	}

	dataUser, err := m.userRepo.GetOneByID(dataTxPiket.IdUsers)
	if err != nil {
		return nil, err
	}

	dataUser2, err := m.userRepo.GetOneByID(dataTxPiket.IdUser2)
	if err != nil {
		return nil, err
	}

	dataTxPiket.DetailKegiatan = dataKegiatan
	dataTxPiket.DetailDataCenter = dataDc
	dataTxPiket.DetailRuangan = dataRuangan
	dataTxPiket.DetailItem = dataItem
	dataTxPiket.DetailUser = dataUser
	dataTxPiket.DetailUserTwo = dataUser2

	return dataTxPiket, nil
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.TxKegiatanPiket, int, error) {
	return m.txpiketRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.txpiketRepo.DeleteOneByID(id)
}
