package user

import (
	"epiketv2/pkg/helper"
	"epiketv2/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	CheckNipExist(nip string) (exist bool)
	CheckUserIsActive(nip string) (exist bool)
	Create(data *model.MsUser) (int64, error)
	UpdateOneByID(data *model.MsUser) (int64, error)
	GetUserMetadataByNip(nip string) (*model.MsUser, error)
	GetOneByID(id int64) (*model.MsUser, error)
	GetOneByNip(id string) (*model.MsUser, error)
	GetAllByID(id int64) ([]*model.MsUser, error)
	GetAll() ([]*model.MsUser, int, error)
	DeleteOneByID(id int64) (int64, error)
	getTotalCount() (totalEntries int)

	Register(nip, nama, no_hp, password string, id_struktur, aktif, id_role int64) (int64, error)
}

type repository struct {
	DB *sqlx.DB
}

// NewRepository ...
func NewRepository() Repository {
	return &repository{
		DB: helper.GetConnection(),
	}
}

func (m *repository) Register(nip, nama, no_hp, password string, id_struktur, aktif, id_role int64) (int64, error) {
	tx, err := m.DB.Begin()
	if err != nil {
		return -1, err
	}

	res, err := tx.Exec(`INSERT INTO ms_users(nip, nama, no_hp, password, id_struktur, aktif, id_role) VALUES(?, ?, ?, ?, ?, ?, ?)`, nip, nama, no_hp, password, id_struktur, aktif, id_role)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	lastIDUser, _ := res.LastInsertId()

	return lastIDUser, tx.Commit()
}

func (m *repository) getTotalCount() (totalEntries int) {
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM ms_users").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.MsUser) (int64, error) {
	query := `INSERT INTO ms_users(
		nip, nama, no_hp, password, id_struktur, aktif, id_role) VALUES(?,?,?,?,?,?,?)`

	res, err := m.DB.Exec(query,
		&data.Nip,
		&data.Nama,
		&data.No_hp,
		&data.Password,
		&data.Id_struktur,
		&data.Aktif,
		&data.Id_role,
	)

	if err != nil {
		return -1, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return lastID, nil
}

func (m *repository) CheckNipExist(nip string) (exist bool) {
	query := `SELECT 
	nip
	FROM ms_users 
	WHERE nip = ?`

	var e string

	if err := m.DB.QueryRow(query, nip).Scan(
		&e,
	); err != nil {
		return false
	}

	if e != "" {
		exist = true
	}

	return
}

func (m *repository) CheckUserIsActive(nip string) (exist bool) {
	query := `SELECT 
	aktif
	FROM ms_users 
	WHERE nip = ?`

	var e int

	if err := m.DB.QueryRow(query, nip).Scan(
		&e,
	); err != nil {
		return false
	}

	if e == 1 {
		exist = true
	}

	return
}

func (m *repository) UpdateOneByID(data *model.MsUser) (int64, error) {
	query := `UPDATE ms_users set nip=?, nama=?, no_hp=?, password=?, id_struktur=?, aktif=?, id_role=?
	WHERE id = ?`

	res, err := m.DB.Exec(query,
		&data.Nip,
		&data.Nama,
		&data.No_hp,
		&data.Password,
		&data.Id_struktur,
		&data.Aktif,
		&data.Id_role,
		&data.ID,
	)

	if err != nil {
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	return rowsAffected, nil
}

func (m *repository) DeleteOneByID(id int64) (int64, error) {
	query := `DELETE FROM ms_users WHERE id = ?`

	res, err := m.DB.Exec(query, id)
	if err != nil {
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	return rowsAffected, nil
}

func (m *repository) GetOneByID(id int64) (*model.MsUser, error) {
	query := `SELECT
	id, nip, nama, no_hp, password, id_struktur, aktif, id_role
	FROM ms_users
	WHERE id = ?`

	data := &model.MsUser{}
	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Nip,
		&data.Nama,
		&data.No_hp,
		&data.Password,
		&data.Id_struktur,
		&data.Aktif,
		&data.Id_role,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *repository) GetOneByNip(nip string) (*model.MsUser, error) {
	query := `SELECT
	id, nip, nama, no_hp, password, id_struktur, aktif, id_role
	FROM ms_users
	WHERE nip = ?`

	data := &model.MsUser{}
	if err := m.DB.QueryRow(query, nip).Scan(
		&data.ID,
		&data.Nip,
		&data.Nama,
		&data.No_hp,
		&data.Password,
		&data.Id_struktur,
		&data.Aktif,
		&data.Id_role,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *repository) GetAllByID(id int64) ([]*model.MsUser, error) {
	var (
		list_data = make([]*model.MsUser, 0)
	)

	query := `SELECT a.id, a.nip, a.nama, a.no_hp, a.password, a.id_struktur, a.aktif, a.id_role, b.id as 'id_struktur', b.nama_struktur, b.nip FROM ms_users as a join ms_struktur as b on b.id = a.id_struktur WHERE a.nip = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data         model.MsUser
			dataStruktur model.MsStruktur
		)

		if err := rows.Scan(
			&data.ID,
			&data.Nip,
			&data.Nama,
			&data.No_hp,
			&data.Password,
			&data.Id_struktur,
			&data.Aktif,
			&data.Id_role,
			&dataStruktur.ID,
			&dataStruktur.Nama_struktur,
			&dataStruktur.Nip,
		); err != nil {
			return nil, err
		}

		data.Struktur = &model.MsStruktur{
			ID:            dataStruktur.ID,
			Nama_struktur: dataStruktur.Nama_struktur,
			Nip:           dataStruktur.Nip,
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

func (m *repository) GetAll() ([]*model.MsUser, int, error) {
	var (
		list = make([]*model.MsUser, 0)
	)

	query := `SELECT a.id, a.nip, a.nama, a.no_hp, a.password, a.id_struktur, a.aktif, a.id_role, b.id as 'id_struktur', b.nama_struktur, b.nip 
	FROM ms_users as a join ms_struktur as b on b.id = a.id_struktur`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data         model.MsUser
			dataStruktur model.MsStruktur
		)

		if err := rows.Scan(
			&data.ID,
			&data.Nip,
			&data.Nama,
			&data.No_hp,
			&data.Password,
			&data.Id_struktur,
			&data.Aktif,
			&data.Id_role,
			&dataStruktur.ID,
			&dataStruktur.Nama_struktur,
			&dataStruktur.Nip,
		); err != nil {
			return nil, -1, err
		}

		data.Struktur = &model.MsStruktur{
			ID:            dataStruktur.ID,
			Nama_struktur: dataStruktur.Nama_struktur,
			Nip:           dataStruktur.Nip,
		}
		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) GetUserMetadataByNip(nip string) (*model.MsUser, error) {
	dataUser := &model.MsUser{}

	query := `SELECT
	id, nip, nama, no_hp, password, id_struktur, aktif, id_role
	FROM ms_users 
	WHERE nip = ?`

	if err := m.DB.QueryRow(query, nip).Scan(
		&dataUser.ID,
		&dataUser.Nip,
		&dataUser.Nama,
		&dataUser.No_hp,
		&dataUser.Password,
		&dataUser.Id_struktur,
		&dataUser.Aktif,
		&dataUser.Id_role,
	); err != nil {
		return nil, err
	}

	return dataUser, nil
}
