package kegiatan

import (
	"epiketv2/pkg/helper"
	"epiketv2/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	checkNamaExist(nama string) (exist bool)
	Create(data *model.MsKegiatan) (int64, error)
	UpdateOneByID(data *model.MsKegiatan) (int64, error)
	GetUserMetadataById(id int64) (*model.MsKegiatan, error)
	GetOneByID(id int64) (*model.MsKegiatan, error)
	GetAllByID(id int64) ([]*model.MsKegiatan, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.MsKegiatan, int, error)
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
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM ms_kegiatan").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.MsKegiatan) (int64, error) {
	query := `INSERT INTO ms_kegiatan(
		nama_kegiatan, deskripsi) VALUES(?,?)`

	res, err := m.DB.Exec(query,
		&data.Nama_kegiatan,
		&data.Deskripsi,
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

func (m *repository) checkNamaExist(nama string) (exist bool) {
	query := `SELECT 
	nama_kegiatan
	FROM ms_kegiatan 
	WHERE nama_kegiatan = ?`

	var e string

	if err := m.DB.QueryRow(query, nama).Scan(
		&e,
	); err != nil {
		return false
	}

	if e != "" {
		exist = true
	}

	return
}

func (m *repository) UpdateOneByID(data *model.MsKegiatan) (int64, error) {
	query := `UPDATE ms_kegiatan set nama_kegiatan=?, deskripsi=?
	WHERE id = ?`

	res, err := m.DB.Exec(query,
		&data.Nama_kegiatan,
		&data.Deskripsi,
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
	query := `DELETE FROM ms_kegiatan WHERE id = ?`

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

func (m *repository) GetOneByID(id int64) (*model.MsKegiatan, error) {
	query := `SELECT
	id, nama_kegiatan, deskripsi
	FROM ms_kegiatan
	WHERE id = ?`

	data := &model.MsKegiatan{}
	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Nama_kegiatan,
		&data.Deskripsi,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *repository) GetAllByID(id int64) ([]*model.MsKegiatan, error) {
	var (
		list_data = make([]*model.MsKegiatan, 0)
	)

	query := `SELECT
	id, nama_kegiatan, deskripsi
	FROM ms_kegiatan
	WHERE id = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.MsKegiatan
		)

		if err := rows.Scan(
			&data.ID,
			&data.Nama_kegiatan,
			&data.Deskripsi,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.MsKegiatan, int, error) {
	var (
		list = make([]*model.MsKegiatan, 0)
	)

	query := `SELECT id, nama_kegiatan, deskripsi FROM ms_kegiatan`

	if dqp.Search != "" {
		query += ` WHERE MATCH(nama_kegiatan) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.MsKegiatan
		)

		if err := rows.Scan(
			&data.ID,
			&data.Nama_kegiatan,
			&data.Deskripsi,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) GetUserMetadataById(id int64) (*model.MsKegiatan, error) {
	query := `SELECT
	id, nama_kegiatan, deskripsi
	FROM ms_kegiatan 
	WHERE id = ?`

	data := &model.MsKegiatan{}

	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Nama_kegiatan,
		&data.Deskripsi,
	); err != nil {
		return nil, err
	}

	return data, nil
}
