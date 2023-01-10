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
	ID            int64  `json:"id"`
	Nama_struktur string `json:"nama_struktur"`
	Nip           string `json:"nip"`
	Parent_id     int64  `json:"parent_id"`
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
	Lokasi  int64  `json:"lokasi"`
}

// Kegiatan ...
type MsKegiatan struct {
	ID            int64  `json:"id"`
	Nama_kegiatan string `json:"nama_kegiatan"`
	Deskripsi     int64  `json:"deskripsi"`
}

// Item ...
type MsItem struct {
	ID         int64      `json:"id"`
	Nama_item  string     `json:"nama_item"`
	Id_ruangan int64      `json:"id_ruangan"`
	Deskripsi  int64      `json:"deskripsi"`
	Parent_id  int64      `json:"parent_id"`
	Ruangan    *MsRuangan `json:"ruangan"`
}

// Piket ...
type TxPiket struct {
	ID         int64      `json:"id"`
	Tanggal    string     `json:"tanggal"`
	Id_user    *time.Time `json:"id_user"`
	UserDetail *MsUser    `json:"user_detail"`
}

// Kegiatan Vendor ...
type TxKegiatanVendor struct {
	ID                int64       `json:"id"`
	Nama_pic          string      `json:"nama_pic"`
	Nama_vendor       string      `json:"nama_vendor"`
	Id_piket          int64       `json:"id_piket"`
	Id_kegiatan       int64       `json:"id_kegiatan"`
	Tanggal_mulai     *time.Time  `json:"tanggaal_mulai"`
	Tanggal_selesai   *time.Time  `json:"tanggal_selesai"`
	Resiko            string      `json:"resiko"`
	Hasil             string      `json:"hasil"`
	IdDokumenKegiatan int64       `json:"id_dokumen_kegiatana"`
	Status            string      `json:"status"`
	Piket             *TxPiket    `json:"piket"`
	Kegiatan          *MsKegiatan `json:"kegiatan"`
}

// Kegiatan Petugas ...
type TxKegiatanPetugas struct {
	ID                int64       `json:"id"`
	Id_piket          int64       `json:"id_piket"`
	Id_kegiatan       int64       `json:"id_kegiatan"`
	Deskripsi         string      `json:"resiko"`
	Hasil             string      `json:"hasil"`
	IdDokumenKegiatan int64       `json:"id_dokumen_kegiatana"`
	Piket             *TxPiket    `json:"piket"`
	Kegiatan          *MsKegiatan `json:"kegiatan"`
}

// Dokumen Kegiatan ...
type TxDokumenKegiatan struct {
	ID               int64  `json:"id"`
	Dokumen_kegiatan string `json:"dokumen_kegiatan"`
}
