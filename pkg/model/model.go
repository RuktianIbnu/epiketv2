package model

import "time"

// DefaultQueryParam ...
type DefaultQueryParam struct {
	Search  string
	Page    int
	Limit   int
	Offset  int
	Sorting map[string]string
	Params  map[string]interface{}
}

// User ...
type MsUser struct {
	ID          int64       `json:"id,omitempty"`
	Nip         string      `json:"nip,omitempty"`
	Nama        string      `json:"nama,omitempty"`
	No_hp       string      `json:"no_hp,omitempty"`
	Password    string      `json:"password,omitempty"`
	Id_struktur int64       `json:"id_struktur,omitempty"`
	Aktif       int64       `json:"aktif,omitempty"`
	Id_role     int64       `json:"id_role,omitempty"`
	Struktur    *MsStruktur `json:"struktur,omitempty"`
}

// Struktur ...
type MsStruktur struct {
	ID            int64       `json:"id,omitempty"`
	Nama_struktur string      `json:"nama_struktur,omitempty"`
	Nip           string      `json:"nip,omitempty"`
	Parent_id     int64       `json:"parent_id,omitempty"`
	Parent_Detail *MsStruktur `json:"parent_detail,omitempty"`
}

// Ruangan ...
type MsRuangan struct {
	ID           int64         `json:"id,omitempty"`
	Nama_ruangan string        `json:"nama_ruangan,omitempty"`
	Id_dc        int64         `json:"id_dc,omitempty"`
	DataCenter   *MsDataCenter `json:"data_center,omitempty"`
}

// Data Center ...
type MsDataCenter struct {
	ID      int64  `json:"id,omitempty"`
	Nama_dc string `json:"nama_dc,omitempty"`
	Lokasi  string `json:"lokasi,omitempty"`
}

// Kegiatan ...
type MsKegiatan struct {
	ID            int64  `json:"id,omitempty"`
	Nama_kegiatan string `json:"nama_kegiatan,omitempty"`
	Deskripsi     string `json:"deskripsi,omitempty"`
}

// Item ...
type MsItem struct {
	ID         int64      `json:"id,omitempty"`
	Nama_item  string     `json:"nama_item,omitempty"`
	Id_ruangan int64      `json:"id_ruangan,omitempty"`
	Deskripsi  string     `json:"deskripsi,omitempty"`
	Parent_id  int64      `json:"parent_id,omitempty"`
	Ruangan    *MsRuangan `json:"ruangan,omitempty"`
}

// Kegiatan Piket ...
type TxKegiatanPiket struct {
	ID               int64         `json:"id,omitempty"`
	IdKegiatan       int64         `json:"id_kegiatan,omitempty"`
	IdDataCenter     int64         `json:"id_data_center,omitempty"`
	IdRuangan        int64         `json:"id_ruangan,omitempty"`
	IdItem           int64         `json:"id_item,omitempty"`
	IdUsers          int64         `json:"id_users,omitempty"`
	NamaPicVendor    string        `json:"nama_pic_vendor,omitempty"`
	NamaPerusahaan   string        `json:"nama_perusahaan,omitempty"`
	TanggalMulai     *time.Time    `json:"tanggal_mulai,omitempty"`
	TanggalSelesai   *time.Time    `json:"tanggal_selesai,omitempty"`
	Deskripsi        string        `json:"deskripsi,omitempty"`
	Resiko           string        `json:"resiko,omitempty"`
	Hasil            string        `json:"hasil,omitempty"`
	Status           string        `json:"status,omitempty"`
	IdUser2          int64         `json:"id_user_2,omitempty"`
	DetailKegiatan   *MsKegiatan   `json:"detail_kegiatan,omitempty"`
	DetailDataCenter *MsDataCenter `json:"detail_dataCenter,omitempty"`
	DetailRuangan    *MsRuangan    `json:"detail_ruangan,omitempty"`
	DetailItem       *MsItem       `json:"detail_item,omitempty"`
	DetailUser       *MsUser       `json:"detail_user,omitempty"`
	DetailUserTwo    *MsUser       `json:"detail_user2,omitempty"`
}

// Kegiatan Harian ...
type TxPiketHarian struct {
	ID               int64         `json:"id,omitempty"`
	Tanggal          *time.Time    `json:"tanggal,omitempty"`
	Jam              *time.Time    `json:"jam,omitempty"`
	IdDataCenter     int64         `json:"id_data_center,omitempty"`
	IdRuangan        int64         `json:"id_ruangan,omitempty"`
	Kondisi          string        `json:"kondisi,omitempty"`
	IdUser1          int64         `json:"id_user_1,omitempty"`
	IdUser2          int64         `json:"id_user_2,omitempty"`
	DetailDataCenter *MsDataCenter `json:"detail_dataCenter,omitempty"`
	DetailRuangan    *MsRuangan    `json:"detail_ruangan,omitempty"`
	DetailUser       *MsUser       `json:"detail_user,omitempty"`
	DetailUserTwo    *MsUser       `json:"detail_user2,omitempty"`
}

type DashKegiatan struct {
	NamaKegiatan string `json:"nama_kegiatan,omitempty"`
	Jumlah       string `json:"jumlah,omitempty"`
}

type DashKondisiAbnormal struct {
	Jumlah int64 `json:"jumlah,omitempty"`
}

type DashStatusPending struct {
	Jumlah int64 `json:"jumlah,omitempty"`
}

type DashTamu struct {
	Jumlah int64 `json:"jumlah,omitempty"`
}

type DashKunjungan struct {
	Jumlah int64 `json:"jumlah,omitempty"`
}

// Report Kegiatan Harian ...
type TxReportPiketHarian struct {
	ID               int64         `json:"id,omitempty"`
	Tahun            string        `json:"tahun,omitempty"`
	Bulan            string        `json:"bulan,omitempty"`
	Tanggal          *time.Time    `json:"tanggal,omitempty"`
	Jam              *time.Time    `json:"jam,omitempty"`
	IdDataCenter     int64         `json:"id_data_center,omitempty"`
	IdRuangan        int64         `json:"id_ruangan,omitempty"`
	Kondisi          string        `json:"kondisi,omitempty"`
	IdUser1          int64         `json:"id_user_1,omitempty"`
	IdUser2          int64         `json:"id_user_2,omitempty"`
	DetailDataCenter *MsDataCenter `json:"detail_dataCenter,omitempty"`
	DetailRuangan    *MsRuangan    `json:"detail_ruangan,omitempty"`
	DetailUser       *MsUser       `json:"detail_user,omitempty"`
	DetailUserTwo    *MsUser       `json:"detail_user2,omitempty"`
}
