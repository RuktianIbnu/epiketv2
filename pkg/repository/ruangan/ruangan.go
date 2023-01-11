package ruangan

import (
	"epiketv2/pkg/helper"
	"epiketv2/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	checkNamaExist(nama string) (exist bool)
	Create(data *model.MsRuangan) (int64, error)
	UpdateOneByID(data *model.MsRuangan) (int64, error)
	GetUserMetadataById(id int64) (*model.MsRuangan, error)
	GetOneByID(id int64) (*model.MsRuangan, error)
	GetAllByID(id int64) ([]*model.MsRuangan, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.MsRuangan, int, error)
	DeleteOneByID(id int64) (int64, error)
	GetTotalCount() (totalEntries int)
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

func (m *repository) GetTotalCount() (totalEntries int) {
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM ms_ruangan").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.MsRuangan) (int64, error) {
	query := `INSERT INTO ms_ruangan(
		id_dc, nama_ruangan) VALUES(?,?)`

	res, err := m.DB.Exec(query,
		&data.Id_dc,
		&data.Nama_ruangan,
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
	nama_ruangan
	FROM ms_ruangan 
	WHERE nama_ruangan = ?`

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

func (m *repository) UpdateOneByID(data *model.MsRuangan) (int64, error) {
	query := `UPDATE ms_ruangan set id_dc=?, nama_ruangan=?
	WHERE id = ?`

	res, err := m.DB.Exec(query,
		&data.Id_dc,
		&data.Nama_ruangan,
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
	query := `DELETE FROM ms_ruangan WHERE id = ?`

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

func (m *repository) GetOneByID(id int64) (*model.MsRuangan, error) {
	query := `SELECT
	id, id_dc, nama_ruangan
	FROM ms_ruangan
	WHERE id = ?`

	data := &model.MsRuangan{}
	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Id_dc,
		&data.Nama_ruangan,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *repository) GetAllByID(id int64) ([]*model.MsRuangan, error) {
	var (
		list_data = make([]*model.MsRuangan, 0)
	)

	query := `SELECT
	id, id_dc, nama_ruangan
	FROM ms_ruangan
	WHERE id = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.MsRuangan
		)

		if err := rows.Scan(
			&data.ID,
			&data.Id_dc,
			&data.Nama_ruangan,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.MsRuangan, int, error) {
	var (
		list = make([]*model.MsRuangan, 0)
	)

	query := `SELECT id, id_dc, nama_ruangan FROM ms_ruangan`

	if dqp.Search != "" {
		query += ` WHERE MATCH(nama_ruangan) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.MsRuangan
		)

		if err := rows.Scan(
			&data.ID,
			&data.Id_dc,
			&data.Nama_ruangan,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, m.GetTotalCount(), nil
}

func (m *repository) GetUserMetadataById(id int64) (*model.MsRuangan, error) {
	query := `SELECT
	id,
	id_dc,
	nama_ruangan
	FROM ms_ruangan 
	WHERE id = ?`

	data := &model.MsRuangan{}

	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Id_dc,
		&data.Nama_ruangan,
	); err != nil {
		return nil, err
	}

	return data, nil
}
