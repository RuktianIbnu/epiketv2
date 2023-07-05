package dash_kegiatan

import (
	"epiketv2/pkg/helper"
	"epiketv2/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	GetAll(dqp *model.DefaultQueryParam) ([]*model.DashKegiatan, int, error)
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

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.DashKegiatan, int, error) {
	var (
		list = make([]*model.DashKegiatan, 0)
	)

	queryStart := `select nama_kegiatan, coalesce(jumlah,0) from vw_dash_kegiatan`

	if dqp.Params["tahun"] != "" {
		queryStart += `	 where tahun =  :tahun`
	}

	rows, err := m.DB.NamedQuery(queryStart, dqp.Params)
	if err != nil {
		return nil, -1, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			data model.DashKegiatan
		)

		if err := rows.Scan(
			&data.NamaKegiatan,
			&data.Jumlah,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, 0, nil
}
