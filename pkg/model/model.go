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
	ID          int64       `json:"id"`
	Nip         string      `json:"nip"`
	Nama        string      `json:"nama"`
	No_hp       string      `json:"no_hp"`
	Password    string      `json:"password"`
	Id_struktur int64       `json:"id_struktur"`
	Aktif       int64       `json:"aktif"`
	Id_role     int64       `json:"id_role"`
	Struktur    *MsStruktur `json:"struktur"`
}

// Struktur ...
type MsStruktur struct {
	ID            int64       `json:"id"`
	Nama_struktur string      `json:"nama_struktur"`
	Nip           string      `json:"nip"`
	Parent_id     int64       `json:"parent_id"`
	Parent_Detail *MsStruktur `json:"parent_detail"`
}

// Ruangan ...
type MsRuangan struct {
	ID           int64         `json:"id"`
	Nama_ruangan string        `json:"nama_ruangan"`
	Id_dc        int64         `json:"id_dc"`
	DataCenter   *MsDataCenter `json:"data_center"`
}

// Data Center ...
type MsDataCenter struct {
	ID      int64  `json:"id"`
	Nama_dc string `json:"nama_dc"`
	Lokasi  string `json:"lokasi"`
}

// Kegiatan ...
type MsKegiatan struct {
	ID            int64  `json:"id"`
	Nama_kegiatan string `json:"nama_kegiatan"`
	Deskripsi     string `json:"deskripsi"`
}

// Item ...
type MsItem struct {
	ID         int64      `json:"id"`
	Nama_item  string     `json:"nama_item"`
	Id_ruangan int64      `json:"id_ruangan"`
	Deskripsi  string     `json:"deskripsi"`
	Parent_id  int64      `json:"parent_id"`
	Ruangan    *MsRuangan `json:"ruangan"`
}

// Kegiatan Piket ...
type TxKegiatanPiket struct {
	ID               int64         `json:"id"`
	IdKegiatan       int64         `json:"id_kegiatan"`
	IdDataCenter     int64         `json:"id_data_center"`
	IdRuangan        int64         `json:"id_ruangan"`
	IdItem           int64         `json:"id_item"`
	IdUsers          int64         `json:"id_users"`
	NamaPicVendor    string        `json:"nama_pic_vendor"`
	NamaPerusahaan   string        `json:"nama_perusahaan"`
	TanggalMulai     *time.Time    `json:"tanggal_mulai"`
	TanggalSelesai   *time.Time    `json:"tanggal_selesai"`
	Deskripsi        string        `json:"deskripsi"`
	Resiko           string        `json:"resiko"`
	Hasil            string        `json:"hasil"`
	Status           string        `json:"status"`
	IdUser2          int64         `json:"id_user_2"`
	DetailKegiatan   *MsKegiatan   `json:"detail_kegiatan"`
	DetailDataCenter *MsDataCenter `json:"detail_dataCenter"`
	DetailRuangan    *MsRuangan    `json:"detail_ruangan"`
	DetailItem       *MsItem       `json:"detail_item"`
	DetailUser       *MsUser       `json:"detail_user"`
	DetailUserTwo    *MsUser       `json:"detail_user2"`
}
