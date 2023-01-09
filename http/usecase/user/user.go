package user

import (
	"epiketv2/pkg/model"
	ur "epiketv2/pkg/repository/user"

	"errors"
)

// Usecase ...
type Usecase interface {
	Create(data *model.MsUser) (int64, error)
	GetOneByID(id int64) (*model.MsUser, error)
	UpdateOneByID(data *model.MsUser) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.MsUser, int, error)
	CheckNIPExist(nip string) bool
	Register(nip string, password string, nama, no_hp string, id_struktur, aktif, id_role int64) (int64, error)
}

type usecase struct {
	user ur.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		ur.NewRepository(),
	}
}

func (m *usecase) Create(data *model.MsUser) (int64, error) {
	return m.user.Create(data)
}

func (m *usecase) Register(nip string, password string, nama, no_hp string, id_struktur, aktif, id_role int64) (int64, error) {
	return m.user.Register(nip, password, nama, no_hp, id_struktur, aktif, id_role)
}

func (m *usecase) CheckNIPExist(nip string) bool {
	return m.user.CheckNipExist(nip)
}

func (m *usecase) UpdateOneByID(data *model.MsUser) (int64, error) {
	rowsAffected, err := m.user.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) (*model.MsUser, error) {
	return m.user.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.MsUser, int, error) {
	return m.user.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.user.DeleteOneByID(id)
}
