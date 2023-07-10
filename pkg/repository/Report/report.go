package report

import (
	"epiketv2/pkg/helper"
	"epiketv2/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	GetAll(dqp *model.DefaultQueryParam) ([]*model.TxPiketHarian, int, error)
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
		query += ` WHERE `
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

	return list, 0, nil
}
