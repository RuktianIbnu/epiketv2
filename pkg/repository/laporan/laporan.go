package laporan

import (
	"epiketv2/pkg/helper"
	"epiketv2/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	GetAll(tahun int64, kode string) ([]*model.TxReportPiketHarian, error)
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
