package item

import (
	"epiketv2/pkg/helper"
	"epiketv2/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	checkNamaExist(nama string) (exist bool)
	Create(data *model.MsItem) (int64, error)
	UpdateOneByID(data *model.MsItem) (int64, error)
	GetUserMetadataById(id int64) (*model.MsItem, error)
	GetOneByID(id int64) (*model.MsItem, error)
	GetAllByID(id int64) ([]*model.MsItem, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.MsItem, int, error)
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
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM ms_item").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.MsItem) (int64, error) {
	query := `INSERT INTO ms_item(
		nama_item, id_ruangan, deskripsi, parent_id) VALUES(?,?,?,?)`

	res, err := m.DB.Exec(query,
		&data.Nama_item,
		&data.Id_ruangan,
		&data.Deskripsi,
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

func (m *repository) checkNamaExist(nama string) (exist bool) {
	query := `SELECT 
	nama_item
	FROM ms_item 
	WHERE nama_item = ?`

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

func (m *repository) UpdateOneByID(data *model.MsItem) (int64, error) {
	query := `UPDATE ms_item set nama_item=?, id_ruangan=?, deskripsi=?, parent_id=?
	WHERE id = ?`

	res, err := m.DB.Exec(query,
		&data.Nama_item,
		&data.Id_ruangan,
		&data.Deskripsi,
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
	query := `DELETE FROM ms_item WHERE id = ?`

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

func (m *repository) GetOneByID(id int64) (*model.MsItem, error) {
	query := `SELECT
	id, nama_item, id_ruangan, deskripsi, parent_id
	FROM ms_item
	WHERE id = ?`

	data := &model.MsItem{}
	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Nama_item,
		&data.Id_ruangan,
		&data.Deskripsi,
		&data.Parent_id,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *repository) GetAllByID(id int64) ([]*model.MsItem, error) {
	var (
		list_data = make([]*model.MsItem, 0)
	)

	query := `SELECT
	id, nama_item, id_ruangan, deskripsi, parent_id
	FROM ms_item
	WHERE id = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.MsItem
		)

		if err := rows.Scan(
			&data.ID,
			&data.Nama_item,
			&data.Id_ruangan,
			&data.Deskripsi,
			&data.Parent_id,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.MsItem, int, error) {
	var (
		list = make([]*model.MsItem, 0)
	)

	query := `SELECT a.id, a.nama_item, a.id_ruangan, b.nama_ruangan, a.deskripsi FROM ms_item as A 
	join ms_ruangan as b on a.id_ruangan = b.id`

	if dqp.Search != "" {
		query += ` WHERE MATCH(nama_item) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data        model.MsItem
			dataRuangan model.MsRuangan
		)

		if err := rows.Scan(
			&data.ID,
			&data.Nama_item,
			&data.Id_ruangan,
			&dataRuangan.Nama_ruangan,
			&data.Deskripsi,
		); err != nil {
			return nil, -1, err
		}

		data.Ruangan = &model.MsRuangan{
			ID:           data.ID,
			Nama_ruangan: dataRuangan.Nama_ruangan,
		}

		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) GetUserMetadataById(id int64) (*model.MsItem, error) {
	query := `SELECT
	id, nama_item, id_ruangan, deskripsi, parent_id
	FROM ms_item 
	WHERE id = ?`

	data := &model.MsItem{}

	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Nama_item,
		&data.Id_ruangan,
		&data.Deskripsi,
		&data.Parent_id,
	); err != nil {
		return nil, err
	}

	return data, nil
}
