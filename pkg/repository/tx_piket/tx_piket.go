package tx_piket

import (
	"epiketv2/pkg/helper"
	"epiketv2/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	Create(data *model.TxKegiatanPiket) (int64, error)
	UpdateByID(data *model.TxKegiatanPiket) (int64, error)
	GetMetadataById(id int64) (*model.TxKegiatanPiket, error)
	GetOneByID(id int64) (*model.TxKegiatanPiket, error)
	GetAllByID(id int64) ([]*model.TxKegiatanPiket, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.TxKegiatanPiket, int, error)
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
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM tx_kegiatan_piket").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.TxKegiatanPiket) (int64, error) {
	query := `INSERT INTO tx_kegiatan_piket(
		id_kegiatan, id_data_center, id_ruangan, id_item, id_users, nama_pic_vendor, nama_perusahaan, tanggal_mulai, tanggal_selesai, deskripsi, resiko, hasil, status, id_user_2) 
		VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	res, err := m.DB.Exec(query,
		&data.IdKegiatan,
		&data.IdDataCenter,
		&data.IdRuangan,
		&data.IdItem,
		&data.IdUsers,
		&data.NamaPicVendor,
		&data.NamaPerusahaan,
		&data.TanggalMulai,
		&data.TanggalSelesai,
		&data.Deskripsi,
		&data.Resiko,
		&data.Hasil,
		&data.Status,
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

func (m *repository) UpdateByID(data *model.TxKegiatanPiket) (int64, error) {
	query := `UPDATE tx_kegiatan_piket set id_kegiatan=?, id_data_center=?, id_ruangan=?, id_item=?, id_users=?, nama_pic_vendor=?, nama_perusahaan=?, tanggal_mulai=?, tanggal_selesai=?, 
	deskripsi=?, resiko=?, hasil=?, status=?, id_user_2=?
	WHERE id = ?`

	res, err := m.DB.Exec(query,
		&data.IdKegiatan,
		&data.IdDataCenter,
		&data.IdRuangan,
		&data.IdItem,
		&data.IdUsers,
		&data.NamaPicVendor,
		&data.NamaPerusahaan,
		&data.TanggalMulai,
		&data.TanggalSelesai,
		&data.Deskripsi,
		&data.Resiko,
		&data.Hasil,
		&data.Status,
		&data.IdUser2,
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
	query := `DELETE FROM tx_kegiatan_piket WHERE id = ?`

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

func (m *repository) GetOneByID(id int64) (*model.TxKegiatanPiket, error) {
	query := `SELECT
	id, id_kegiatan, id_data_center, id_ruangan, id_item, id_users, nama_pic_vendor, nama_perusahaan, tanggal_mulai, tanggal_selesai, deskripsi, resiko, hasil, status, id_user_2
	FROM tx_kegiatan_piket
	WHERE id = ?`

	data := &model.TxKegiatanPiket{}
	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.IdKegiatan,
		&data.IdDataCenter,
		&data.IdRuangan,
		&data.IdItem,
		&data.IdUsers,
		&data.NamaPicVendor,
		&data.NamaPerusahaan,
		&data.TanggalMulai,
		&data.TanggalSelesai,
		&data.Deskripsi,
		&data.Resiko,
		&data.Hasil,
		&data.Status,
		&data.IdUser2,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *repository) GetAllByID(id int64) ([]*model.TxKegiatanPiket, error) {
	var (
		list_data = make([]*model.TxKegiatanPiket, 0)
	)

	query := `SELECT
	id, id_kegiatan, id_data_center, id_ruangan, id_item, id_users, nama_pic_vendor, nama_perusahaan, tanggal_mulai, tanggal_selesai, deskripsi, resiko, hasil, status, id_user_2
	FROM tx_kegiatan_piket
	WHERE id = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.TxKegiatanPiket
		)

		if err := rows.Scan(
			&data.ID,
			&data.IdKegiatan,
			&data.IdDataCenter,
			&data.IdRuangan,
			&data.IdItem,
			&data.IdUsers,
			&data.NamaPicVendor,
			&data.NamaPerusahaan,
			&data.TanggalMulai,
			&data.TanggalSelesai,
			&data.Deskripsi,
			&data.Resiko,
			&data.Hasil,
			&data.Status,
			&data.IdUser2,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.TxKegiatanPiket, int, error) {
	var (
		list = make([]*model.TxKegiatanPiket, 0)
	)

	query := `SELECT
	a.id, a.id_kegiatan, a.id_data_center, a.id_ruangan, a.id_item, a.id_users, a.nama_pic_vendor, a.nama_perusahaan, a.tanggal_mulai, a.tanggal_selesai, a.deskripsi, a.resiko, 
	a.hasil, a.status, a.id_user_2, b.nama_kegiatan, b.deskripsi, c.nama_dc, c.lokasi, d.nama_ruangan, e.nama_item, e.deskripsi, f.nip, f.nama, f.no_hp, g.nip, g.nama, g.no_hp
	FROM tx_kegiatan_piket as a
	join ms_kegiatan as b on a.id = a.id_kegiatan
	join ms_data_center as c on c.id = a.id_data_center
	join ms_ruangan as d on d.id = a.id_ruangan
	join ms_item as e on e.id = a.id_item
	join ms_users as f on f.id = a.id_users
	join ms_users as g on g.id = a.id_user_2`

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
			data       model.TxKegiatanPiket
			dataCenter model.MsDataCenter
			kegiatan   model.MsKegiatan
			ruangan    model.MsRuangan
			item       model.MsItem
			user       model.MsUser
			user2      model.MsUser
		)

		if err := rows.Scan(
			&data.ID,
			&data.IdKegiatan,
			&data.IdDataCenter,
			&data.IdRuangan,
			&data.IdItem,
			&data.IdUsers,
			&data.NamaPicVendor,
			&data.NamaPerusahaan,
			&data.TanggalMulai,
			&data.TanggalSelesai,
			&data.Deskripsi,
			&data.Resiko,
			&data.Hasil,
			&data.Status,
			&data.IdUser2,
		); err != nil {
			return nil, -1, err
		}

		data.DetailKegiatan = &model.MsKegiatan{
			ID:            kegiatan.ID,
			Nama_kegiatan: kegiatan.Nama_kegiatan,
			Deskripsi:     kegiatan.Deskripsi,
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
		data.DetailItem = &model.MsItem{
			ID:        item.ID,
			Nama_item: item.Nama_item,
			Deskripsi: item.Deskripsi,
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

func (m *repository) GetMetadataById(id int64) (*model.TxKegiatanPiket, error) {
	query := `SELECT
	id, id_kegiatan, id_data_center, id_ruangan, id_item, id_users, nama_pic_vendor, nama_perusahaan, tanggal_mulai, tanggal_selesai, deskripsi, resiko, hasil, status, id_user_2
	FROM tx_kegiatan_piket 
	WHERE id = ?`

	data := &model.TxKegiatanPiket{}

	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.IdKegiatan,
		&data.IdDataCenter,
		&data.IdRuangan,
		&data.IdItem,
		&data.IdUsers,
		&data.NamaPicVendor,
		&data.NamaPerusahaan,
		&data.TanggalMulai,
		&data.TanggalSelesai,
		&data.Deskripsi,
		&data.Resiko,
		&data.Hasil,
		&data.Status,
		&data.IdUser2,
	); err != nil {
		return nil, err
	}

	return data, nil
}
