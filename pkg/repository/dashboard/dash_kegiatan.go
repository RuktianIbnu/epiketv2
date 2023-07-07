package dash_kegiatan

import (
	"epiketv2/pkg/helper"
	"epiketv2/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	GetAll(dqp *model.DefaultQueryParam) ([]*model.DashKegiatan, int, error)
	GetAllKondisiAbnormal(dqp *model.DefaultQueryParam) ([]*model.DashKondisiAbnormal, int, error)
	GetAllStatusPending(dqp *model.DefaultQueryParam) ([]*model.DashStatusPending, int, error)
	GetAllKunjungan(dqp *model.DefaultQueryParam) ([]*model.DashKunjungan, int, error)
	GetAllTamu(dqp *model.DefaultQueryParam) ([]*model.DashTamu, int, error)
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

func (m *repository) GetAllKondisiAbnormal(dqp *model.DefaultQueryParam) ([]*model.DashKondisiAbnormal, int, error) {
	var (
		list = make([]*model.DashKondisiAbnormal, 0)
	)

	queryStart := `select coalesce(jumlah,0) from vw_kondisi_abnormal`

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
			data model.DashKondisiAbnormal
		)

		if err := rows.Scan(
			&data.Jumlah,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, 0, nil
}

func (m *repository) GetAllStatusPending(dqp *model.DefaultQueryParam) ([]*model.DashStatusPending, int, error) {
	var (
		list = make([]*model.DashStatusPending, 0)
	)

	queryStart := `select coalesce(jumlah,0) from vw_status_pending`

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
			data model.DashStatusPending
		)

		if err := rows.Scan(
			&data.Jumlah,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, 0, nil
}

func (m *repository) GetAllKunjungan(dqp *model.DefaultQueryParam) ([]*model.DashKunjungan, int, error) {
	var (
		list = make([]*model.DashKunjungan, 0)
	)

	queryStart := `select coalesce(jumlah,0) from vw_kunjungan`

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
			data model.DashKunjungan
		)

		if err := rows.Scan(
			&data.Jumlah,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, 0, nil
}

func (m *repository) GetAllTamu(dqp *model.DefaultQueryParam) ([]*model.DashTamu, int, error) {
	var (
		list = make([]*model.DashTamu, 0)
	)

	queryStart := `select count(*) as jumlah from vw_tamu`

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
			data model.DashTamu
		)

		if err := rows.Scan(
			&data.Jumlah,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, 0, nil
}
