package data_center

import (
	"epiketv2/pkg/helper"
	"epiketv2/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	checkNamaExist(nama string) (exist bool)
	Create(data *model.MsDataCenter) (int64, error)
	UpdateOneByID(data *model.MsDataCenter) (int64, error)
	GetUserMetadataById(id int64) (*model.MsDataCenter, error)
	GetOneByID(id int64) (*model.MsDataCenter, error)
	GetAllByID(id int64) ([]*model.MsDataCenter, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.MsDataCenter, int, error)
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
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM ms_data_center").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.MsDataCenter) (int64, error) {
	query := `INSERT INTO ms_data_center(
		nama_dc, lokasi) VALUES(?,?)`

	res, err := m.DB.Exec(query,
		&data.Nama_dc,
		&data.Lokasi,
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
	nama_dc
	FROM ms_data_center 
	WHERE nama_dc = ?`

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

func (m *repository) UpdateOneByID(data *model.MsDataCenter) (int64, error) {
	query := `UPDATE ms_data_center set nama_dc=?, lokasi=?
	WHERE id = ?`

	res, err := m.DB.Exec(query,
		&data.Nama_dc,
		&data.Lokasi,
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
	query := `DELETE FROM ms_data_center WHERE id = ?`

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

func (m *repository) GetOneByID(id int64) (*model.MsDataCenter, error) {
	query := `SELECT
	id, nama_dc, lokasi
	FROM ms_data_center
	WHERE id = ?`

	data := &model.MsDataCenter{}
	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Nama_dc,
		&data.Lokasi,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *repository) GetAllByID(id int64) ([]*model.MsDataCenter, error) {
	var (
		list_data = make([]*model.MsDataCenter, 0)
	)

	query := `SELECT
	id, nama_dc, lokasi
	FROM ms_data_center
	WHERE id = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.MsDataCenter
		)

		if err := rows.Scan(
			&data.ID,
			&data.Nama_dc,
			&data.Lokasi,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.MsDataCenter, int, error) {
	var (
		list = make([]*model.MsDataCenter, 0)
	)

	query := `SELECT id, nama_dc, lokasi FROM ms_data_center`

	if dqp.Search != "" {
		query += ` WHERE MATCH(nama_dc) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.MsDataCenter
		)

		if err := rows.Scan(
			&data.ID,
			&data.Nama_dc,
			&data.Lokasi,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) GetUserMetadataById(id int64) (*model.MsDataCenter, error) {
	query := `SELECT
	id,
	nama_dc,
	lokasi
	FROM ms_data_center 
	WHERE id = ?`

	data := &model.MsDataCenter{}

	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Nama_dc,
		&data.Lokasi,
	); err != nil {
		return nil, err
	}

	return data, nil
}
