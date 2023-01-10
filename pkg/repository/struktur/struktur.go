package struktur

import (
	"epiketv2/pkg/helper"
	"epiketv2/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	CheckNamaIsExist(namastruktur string) (exist bool)
	Create(data *model.MsStruktur) (int64, error)
	UpdateOneByID(data *model.MsStruktur) (int64, error)
	GetUserMetadataById(id int64) (*model.MsStruktur, error)
	GetOneByID(id int64) (*model.MsStruktur, error)
	GetAllByID(id int64) ([]*model.MsStruktur, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.MsStruktur, int, error)
	DeleteOneByID(id int64) (int64, error)
	getTotalCount() (totalEntries int)
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

func (m *repository) getTotalCount() (totalEntries int) {
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM ms_struktur").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.MsStruktur) (int64, error) {
	query := `INSERT INTO ms_struktur(
		nama_struktur, nip, parent_id) VALUES(?,?,?)`

	res, err := m.DB.Exec(query,
		&data.Nip,
		&data.Nama_struktur,
		&data.Parent_id,
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

func (m *repository) CheckNamaIsExist(namastruktur string) (exist bool) {
	query := `SELECT 
	count(*)
	FROM ms_struktur 
	WHERE nama_struktur = ?`

	var e int64

	if err := m.DB.QueryRow(query, namastruktur).Scan(
		&e,
	); err != nil {
		return false
	}

	if e == 0 {
		exist = true
	}

	return
}

func (m *repository) UpdateOneByID(data *model.MsStruktur) (int64, error) {
	query := `UPDATE ms_struktur set nama_struktur=?, nip=?, parent_id=?
	WHERE id = ?`

	res, err := m.DB.Exec(query,
		&data.Nama_struktur,
		&data.Nip,
		&data.Parent_id,
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
	query := `DELETE FROM ms_struktur WHERE id = ?`

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

func (m *repository) GetOneByID(id int64) (*model.MsStruktur, error) {
	query := `SELECT
	id, nama_struktur, nip, parent_id
	FROM ms_struktur
	WHERE id = ?`

	data := &model.MsStruktur{}
	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Nama_struktur,
		&data.Nip,
		&data.Parent_id,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *repository) GetAllByID(id int64) ([]*model.MsStruktur, error) {
	var (
		list_data = make([]*model.MsStruktur, 0)
	)

	query := `SELECT
	id, nama_struktur, nip, parent_id
	FROM ms_struktur
	WHERE id = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.MsStruktur
		)

		if err := rows.Scan(
			&data.ID,
			&data.Nama_struktur,
			&data.Nip,
			&data.Parent_id,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.MsStruktur, int, error) {
	var (
		list = make([]*model.MsStruktur, 0)
	)

	query := `SELECT id, nama_struktur, nip, parent_id FROM ms_struktur`

	if dqp.Search != "" {
		query += ` WHERE MATCH(nip, nama_struktur) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.MsStruktur
		)

		if err := rows.Scan(
			&data.ID,
			&data.Nama_struktur,
			&data.Nip,
			&data.Parent_id,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) GetUserMetadataById(id int64) (*model.MsStruktur, error) {
	query := `SELECT
	id,
	nama_struktur,
	nip,
	parent_id
	FROM ms_struktur 
	WHERE id = ?`

	data := &model.MsStruktur{}

	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Nama_struktur,
		&data.Nip,
		&data.Parent_id,
	); err != nil {
		return nil, err
	}

	return data, nil
}
