package dash_kegiatan

import (
	"epiketv2/pkg/model"
	dkr "epiketv2/pkg/repository/dashboard"
)

// Usecase ...
type Usecase interface {
	GetAll(dqp *model.DefaultQueryParam) ([]*model.DashKegiatan, int, error)
	GetAllKondisiAbnormal(dqp *model.DefaultQueryParam) ([]*model.DashKondisiAbnormal, int, error)
	GetAllStatusPending(dqp *model.DefaultQueryParam) ([]*model.DashStatusPending, int, error)
	GetAllKunjungan(dqp *model.DefaultQueryParam) ([]*model.DashKunjungan, int, error)
	GetAllTamu(dqp *model.DefaultQueryParam) ([]*model.DashTamu, int, error)
}

type usecase struct {
	dashKegiatan dkr.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		dkr.NewRepository(),
	}
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.DashKegiatan, int, error) {
	return m.dashKegiatan.GetAll(dqp)
}

func (m *usecase) GetAllKondisiAbnormal(dqp *model.DefaultQueryParam) ([]*model.DashKondisiAbnormal, int, error) {
	return m.dashKegiatan.GetAllKondisiAbnormal(dqp)
}

func (m *usecase) GetAllStatusPending(dqp *model.DefaultQueryParam) ([]*model.DashStatusPending, int, error) {
	return m.dashKegiatan.GetAllStatusPending(dqp)
}

func (m *usecase) GetAllKunjungan(dqp *model.DefaultQueryParam) ([]*model.DashKunjungan, int, error) {
	return m.dashKegiatan.GetAllKunjungan(dqp)
}

func (m *usecase) GetAllTamu(dqp *model.DefaultQueryParam) ([]*model.DashTamu, int, error) {
	return m.dashKegiatan.GetAllTamu(dqp)
}
