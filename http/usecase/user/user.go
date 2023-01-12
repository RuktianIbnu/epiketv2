package user

import (
	"epiketv2/pkg/helper/bcrypt"
	"epiketv2/pkg/helper/jwt"
	"epiketv2/pkg/model"
	sr "epiketv2/pkg/repository/struktur"
	ur "epiketv2/pkg/repository/user"
	"errors"
	"fmt"
)

// Usecase ...
type Usecase interface {
	Create(data *model.MsUser) (int64, error)
	GetOneByID(id int64) (*model.MsUser, error)
	GetOneByNip(nip string) (*model.MsUser, error)
	UpdateOneByID(data *model.MsUser) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll() ([]*model.MsUser, int, error)
	CheckNIPExist(nip string) bool
	Register(nip string, nama, no_hp, password string, id_struktur, aktif, id_role int64) (int64, error)
	Login(nip, password string) (string, error)
}

type usecase struct {
	user     ur.Repository
	struktur sr.Repository
}

const (
	// ActionLogin ...
	ActionLogin = "login"

	// RoleRespondent ...
	RoleRespondent = "respondent"
)

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		ur.NewRepository(),
		sr.NewRepository(),
	}
}

func (m *usecase) Create(data *model.MsUser) (int64, error) {
	return m.user.Create(data)
}

func (m *usecase) Login(nip, password string) (string, error) {
	userMetadata, err := m.user.GetUserMetadataByNip(nip)
	if err != nil {
		return "", errors.New("NIP not registered")
	}

	userIsActive := m.user.CheckUserIsActive(nip)
	if !userIsActive {
		return "", fmt.Errorf("please activate your account")
	}

	if !bcrypt.Compare(password, userMetadata.Password) {
		return "", fmt.Errorf("incorrect email or password")
	}

	tempToken, err := jwt.CreateToken(userMetadata.ID, userMetadata.Nip, userMetadata.Nama, userMetadata.Id_role)
	if err != nil {
		return "", fmt.Errorf("failed generate temporary token %s", err.Error())
	}

	return tempToken, nil
}

func (m *usecase) Register(nip string, nama, no_hp, password string, id_struktur, aktif, id_role int64) (int64, error) {
	hashedPwd, err := bcrypt.Hash(password)
	if err != nil {
		return 500, err
	}

	return m.user.Register(nip, nama, no_hp, hashedPwd, id_struktur, aktif, id_role)
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

func (m *usecase) GetOneByNip(nip string) (*model.MsUser, error) {
	dataUser, err := m.user.GetOneByNip(nip)
	if err != nil {
		return nil, err
	}

	dataStruktur, err := m.struktur.GetOneByID(dataUser.Id_struktur)
	if err != nil {
		return nil, err
	}

	dataUser.Struktur = dataStruktur

	return dataUser, nil
}

func (m *usecase) GetAll() ([]*model.MsUser, int, error) {
	dataUser, count, err := m.user.GetAll()
	if err != nil {
		return nil, -1, err
	}
	return dataUser, count, err
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.user.DeleteOneByID(id)
}
