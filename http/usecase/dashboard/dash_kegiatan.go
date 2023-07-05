package dash_kegiatan

import (
	"epiketv2/pkg/model"
	dkr "epiketv2/pkg/repository/dashboard"
)

// Usecase ...
type Usecase interface {
	GetAll(dqp *model.DefaultQueryParam) ([]*model.DashKegiatan, int, error)
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
