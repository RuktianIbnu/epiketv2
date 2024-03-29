package laporan

import (
	"epiketv2/pkg/helper"
	"epiketv2/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	GetAll(tahun, bulan, id_datacenter int64, tanggal string) ([]*model.TxReportPiketHarian, error)
	GetReportKegiatanDc(tahun, bulan, id_datacenter int64, tanggal string) ([]*model.TxReportKegiatanPiketDc, error)
	GetReportKunjungan(tahun, bulan, id_datacenter int64, tanggal string) ([]*model.TxReportKegiatanPiketDc, error)
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

func (m *repository) GetReportKunjungan(tahun, bulan, id_datacenter int64, tanggal string) ([]*model.TxReportKegiatanPiketDc, error) {
	var (
		list = make([]*model.TxReportKegiatanPiketDc, 0)
	)

	queryStart := `SELECT id, id_kegiatan, id_data_center, id_ruangan, id_item, id_users, nama_pic_vendor, nama_perusahaan,
	tanggal_mulai, tanggal_selesai, deskripsi, resiko, hasil, status, id_user_2, nama_kegiatan, deskripsi_kegiatan, nama_dc, 
	lokasi, nama_ruangan, nama_item, deskripsi_item, nip, nama, no_hp, nip_user2, nama_user2, no_hp_user2 
	from vw_kegiatan_dc`

	if tahun != 0 {
		queryStart += ` WHERE tahun = :tahun`
		if bulan != 0 {
			queryStart += ` and bulan = :bulan`
			if id_datacenter != 0 {
				queryStart += ` and id_data_center = :id_datacenter`
			}
		}
	}

	// fmt.Println(queryStart)

	rows, err := m.DB.NamedQuery(queryStart, map[string]interface{}{
		"tahun":         tahun,
		"bulan":         bulan,
		"id_datacenter": id_datacenter,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data       model.TxReportKegiatanPiketDc
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

func (m *repository) GetReportKegiatanDc(tahun, bulan, id_datacenter int64, tanggal string) ([]*model.TxReportKegiatanPiketDc, error) {
	var (
		list = make([]*model.TxReportKegiatanPiketDc, 0)
	)

	queryStart := `SELECT id, id_kegiatan, id_data_center, id_ruangan, id_item, id_users, nama_pic_vendor, nama_perusahaan,
	tanggal_mulai, tanggal_selesai, deskripsi, resiko, hasil, status, id_user_2, nama_kegiatan, deskripsi_kegiatan, nama_dc, 
	lokasi, nama_ruangan, nama_item, deskripsi_item, nip, nama, no_hp, nip_user2, nama_user2, no_hp_user2 
	from vw_kegiatan_dc`

	if tanggal != "" {
		queryStart += ` WHERE tanggal_mulai = :tanggal`
	}

	// fmt.Println(tanggal, queryStart)

	rows, err := m.DB.NamedQuery(queryStart, map[string]interface{}{
		"tanggal": tanggal,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data       model.TxReportKegiatanPiketDc
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

func (m *repository) GetAll(tahun, bulan, id_datacenter int64, tanggal string) ([]*model.TxReportPiketHarian, error) {
	var (
		list = make([]*model.TxReportPiketHarian, 0)
	)
	queryStart := `SELECT
	id, tahun, bulan, tanggal, jam, id_data_center, id_ruangan, kondisi, id_user_1, id_user_2, nama_dc, lokasi, nama_ruangan, 
	nip, nama, no_hp, nip_user2, nama_user2, no_hp_user2
	FROM vw_monitoring_harian`

	if tahun != 0 {
		queryStart += ` WHERE tahun = :tahun`
		if bulan != 0 {
			queryStart += ` and bulan = :bulan`
			if id_datacenter != 0 {
				queryStart += ` and id_data_center = :id_datacenter`
			}
		}
	}

	rows, err := m.DB.NamedQuery(queryStart, map[string]interface{}{
		"tahun":         tahun,
		"bulan":         bulan,
		"id_datacenter": id_datacenter,
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
