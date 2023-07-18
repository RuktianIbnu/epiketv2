package laporan

import (
	"epiketv2/pkg/helper"
	"epiketv2/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	GetAll(tahun int64, kode string) ([]*model.TxReportPiketHarian, error)
	GetKegiatanPiket(tahun int64, kode string) ([]*model.TxKegiatanPiketDc, error)
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

func (m *repository) GetKegiatanPiket(tahun int64, kode string) ([]*model.TxKegiatanPiketDc, error) {
	var (
		list = make([]*model.TxKegiatanPiketDc, 0)
	)

	queryStart := `SELECT
	a.id, a.id_kegiatan, a.id_data_center, a.id_ruangan, a.id_item, a.id_users, a.nama_pic_vendor, a.nama_perusahaan, a.tanggal_mulai, a.tanggal_selesai, a.deskripsi, a.resiko, 
	a.hasil, a.status, a.id_user_2, COALESCE(b.nama_kegiatan,0) as nama_kegiatan, COALESCE(b.deskripsi,0) as deskripsi, c.nama_dc, c.lokasi, d.nama_ruangan, 
	COALESCE(e.nama_item,0) as nama_item, 
	COALESCE(e.deskripsi,0) as deskripsi, 
	f.nip, f.nama, f.no_hp, g.nip as nip_user2, 
	g.nama as nama_user2, g.no_hp as no_hp_user2
	FROM tx_kegiatan_piket as a
	left join ms_kegiatan as b on b.id = a.id_kegiatan
	left join ms_data_center as c on c.id = a.id_data_center
	left join ms_ruangan as d on d.id = a.id_ruangan
	left join ms_item as e on e.id = a.id_item
	left join ms_users as f on f.id = a.id_users
	left join ms_users as g on g.id = a.id_user_2`

	if tahun != 0 {
		queryStart += ` WHERE tahun = :tahun`
	}
	queryStart += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(queryStart, map[string]interface{}{
		"tahun": tahun,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data       model.TxKegiatanPiketDc
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
			&kegiatan.Nama_kegiatan,
			&kegiatan.Deskripsi,
			&dataCenter.Nama_dc,
			&dataCenter.Lokasi,
			&ruangan.Nama_ruangan,
			&item.Nama_item,
			&item.Deskripsi,
			&user.Nip,
			&user.Nama,
			&user.No_hp,
			&user2.Nip,
			&user2.Nama,
			&user2.No_hp,
		); err != nil {
			return nil, err
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

	return list, nil
}

func (m *repository) GetAll(tahun int64, kode string) ([]*model.TxReportPiketHarian, error) {
	var (
		list = make([]*model.TxReportPiketHarian, 0)
	)
	queryStart := `SELECT
	id, tahun, bulan, tanggal, jam, id_data_center, id_ruangan, kondisi, id_user_1, id_user_2, nama_dc, lokasi, nama_ruangan, 
	nip, nama, no_hp, nip_user2, nama_user2, no_hp_user2
	FROM vw_monitoring_harian`

	if tahun != 0 {
		queryStart += ` WHERE tahun = :tahun`
	}

	rows, err := m.DB.NamedQuery(queryStart, map[string]interface{}{
		"tahun": tahun,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data       model.TxReportPiketHarian
			dataCenter model.MsDataCenter
			ruangan    model.MsRuangan
			user       model.MsUser
			user2      model.MsUser
		)

		if err := rows.Scan(
			&data.ID,
			&data.Tahun,
			&data.Bulan,
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
			return nil, err
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

	return list, nil
}
