package tx_harian

import (
	"epiketv2/pkg/helper"
	"epiketv2/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	Create(data *model.TxPiketHarian) (int64, error)
	UpdateByID(data *model.TxPiketHarian) (int64, error)
	GetMetadataById(id int64) (*model.TxPiketHarian, error)
	GetOneByID(id int64) (*model.TxPiketHarian, error)
	GetAllByID(id int64) ([]*model.TxPiketHarian, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.TxPiketHarian, int, error)
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
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM tx_kegiatan_harian").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.TxPiketHarian) (int64, error) {
	query := `INSERT INTO tx_kegiatan_harian(
		tanggal, jam, id_data_center, id_ruangan, kondisi, id_user_1, id_user_2) 
		VALUES(?,?,?,?,?,?,?)`

	res, err := m.DB.Exec(query,
		&data.Tanggal,
		&data.Jam,
		&data.IdDataCenter,
		&data.IdRuangan,
		&data.Kondisi,
		&data.IdUser1,
		&data.IdUser2,
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

func (m *repository) UpdateByID(data *model.TxPiketHarian) (int64, error) {
	query := `UPDATE tx_kegiatan_harian set tanggal=?, jam=?, id_data_center=?, id_ruangan=?, kondisi=?, id_user_1=?, id_user_2=?
	WHERE id = ?`

	res, err := m.DB.Exec(query,
		&data.Tanggal,
		&data.Jam,
		&data.IdDataCenter,
		&data.IdRuangan,
		&data.Kondisi,
		&data.IdUser1,
		&data.IdUser2,
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
	query := `DELETE FROM tx_kegiatan_harian WHERE id = ?`

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

func (m *repository) GetOneByID(id int64) (*model.TxPiketHarian, error) {
	query := `SELECT
	id, tanggal, jam, id_data_center, id_ruangan, kondisi, id_user_1, id_user_2
	FROM tx_kegiatan_harian
	WHERE id = ?`

	data := &model.TxPiketHarian{}
	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Tanggal,
		&data.Jam,
		&data.IdDataCenter,
		&data.IdRuangan,
		&data.Kondisi,
		&data.IdUser1,
		&data.IdUser2,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *repository) GetAllByID(id int64) ([]*model.TxPiketHarian, error) {
	var (
		list_data = make([]*model.TxPiketHarian, 0)
	)

	query := `SELECT
	id, tanggal, jam, id_data_center, id_ruangan, kondisi, id_user_1, id_user_2
	FROM tx_kegiatan_harian
	WHERE id = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.TxPiketHarian
		)

		if err := rows.Scan(
			&data.ID,
			&data.Tanggal,
			&data.Jam,
			&data.IdDataCenter,
			&data.IdRuangan,
			&data.Kondisi,
			&data.IdUser1,
			&data.IdUser2,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.TxPiketHarian, int, error) {
	var (
		list = make([]*model.TxPiketHarian, 0)
	)

	query := `SELECT
	a.id, a.tanggal, a.jam, a.id_data_center, a.id_ruangan, a.kondisi, a.id_user_1, a.id_user_2, c.nama_dc, c.lokasi, d.nama_ruangan, 
	f.nip, f.nama, f.no_hp, g.nip as nip_user2, 
	g.nama as nama_user2, g.no_hp as no_hp_user2
	FROM tx_kegiatan_harian as a
	left join ms_data_center as c on c.id = a.id_data_center
	left join ms_ruangan as d on d.id = a.id_ruangan
	left join ms_users as f on f.id = a.id_user_1
	left join ms_users as g on g.id = a.id_user_2`

	if dqp.Search != "" {
		query += ` WHERE MATCH(a.nama_ruangan) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data       model.TxPiketHarian
			dataCenter model.MsDataCenter
			ruangan    model.MsRuangan
			user       model.MsUser
			user2      model.MsUser
		)

		if err := rows.Scan(
			&data.ID,
			&data.Tanggal,
			&data.Jam,
			&data.IdDataCenter,
			&data.IdRuangan,
			&data.Kondisi,
			&data.IdUser1,
			&data.IdUser2,
			&dataCenter.Nama_dc,
			&dataCenter.Lokasi,
			&ruangan.Nama_ruangan,
			&user.Nip,
			&user.Nama,
			&user.No_hp,
			&user2.Nip,
			&user2.Nama,
			&user2.No_hp,
		); err != nil {
			return nil, -1, err
		}
		data.DetailDataCenter = &model.MsDataCenter{
			ID:      dataCenter.ID,
			Nama_dc: dataCenter.Nama_dc,
			Lokasi:  dataCenter.Lokasi,
		}
		data.DetailRuangan = &model.MsRuangan{
			ID:           ruangan.ID,
			Nama_ruangan: ruangan.Nama_ruangan,
		}
		data.DetailUser = &model.MsUser{
			ID:    user.ID,
			Nip:   user.Nip,
			Nama:  user.Nama,
			No_hp: user.No_hp,
		}
		data.DetailUserTwo = &model.MsUser{
			ID:    user2.ID,
			Nip:   user2.Nip,
			Nama:  user2.Nama,
			No_hp: user2.No_hp,
		}

		list = append(list, &data)
	}

	return list, m.GetTotalCount(), nil
}

func (m *repository) GetMetadataById(id int64) (*model.TxPiketHarian, error) {
	query := `SELECT
	id, tanggal, jam, id_data_center, id_ruangan, kondisi, id_user_1, id_user_2
	FROM tx_kegiatan_harian 
	WHERE id = ?`

	data := &model.TxPiketHarian{}

	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.ID,
		&data.Tanggal,
		&data.Jam,
		&data.IdDataCenter,
		&data.IdRuangan,
		&data.Kondisi,
		&data.IdUser1,
		&data.IdUser2,
	); err != nil {
		return nil, err
	}

	return data, nil
}
