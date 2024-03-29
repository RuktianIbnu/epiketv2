package harian

import (
	rr "epiketv2/pkg/repository/laporan"
	"fmt"
	"os"
	"strconv"

	// "strconv"

	"github.com/jung-kurt/gofpdf"
)

// Usecase ...
type Usecase interface {
	GetLaporanHarian(tahun, bulan, id_datacenter int64, tanggal string) (string, error)
	GetLaporanKegiatan(tahun, bulan, id_datacenter int64, tanggal string) (string, error)
	GetLaporanKunjungan(tahun, bulan, id_datacenter int64, tanggal string) (string, error)
}

type usecase struct {
	reportRepo rr.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		reportRepo: rr.NewRepository(),
	}
}

// ReportHarian ...
type ReportHarian struct {
	Title string
}

// ReportDataHarian ...
type ReportDataHarian struct {
	Tanggal      string
	NamaDc       string
	Lokasi       string
	NamaRuangan  string
	Kondisi      string
	NipPetugas1  string
	NamaPetugas1 string
	NipPetugas2  string
	NamaPetugas2 string
}

func (m *usecase) GetLaporanKunjungan(tahun, bulan, id_datacenter int64, tanggal string) (string, error) {
	dataKegiatan, err := m.reportRepo.GetReportKunjungan(tahun, bulan, id_datacenter, tanggal)
	if err != nil {
		return "", err
	}

	var (
		baseFont   string  = "Arial"
		cellHeight float64 = 5.5
	)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(5, 10, 5)
	pdf.SetFont(baseFont, "", 7)
	pdf.SetDrawColor(237, 125, 49)
	pdf.SetFillColor(237, 125, 49)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetAutoPageBreak(true, 10)

	pdf.SetFooterFunc(func() {
		pdf.SetY(-12)
		pdf.SetFont(baseFont, "I", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("Halaman %d dari {nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
	})

	pdf.AliasNbPages("")
	pdf.AddPage()

	pdf.SetFont(baseFont, "B", 11)
	pdf.CellFormat(200, cellHeight, "Laporan Kunjungan Data Center", "", 1, "CM", false, 0, "")
	if bulan != 0 {
		pdf.CellFormat(200, cellHeight, "Tahun : "+strconv.FormatInt(tahun, 10)+" bulan :"+strconv.FormatInt(bulan, 10), "", 1, "CM", false, 0, "")
	} else {
		pdf.CellFormat(200, cellHeight, "Tahun : "+strconv.FormatInt(tahun, 10), "", 1, "CM", false, 0, "")
	}

	pdf.CellFormat(10, cellHeight*2, "No", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(25, cellHeight*2, "Tanggal", "LTRB", 0, "LM", true, 0, "")
	pdf.CellFormat(40, cellHeight*2, "Data Center", "LTRB", 0, "LM", true, 0, "")
	pdf.CellFormat(40, cellHeight*2, "Instansi / Perusahaan", "LTRB", 0, "LM", true, 0, "")
	pdf.CellFormat(45, cellHeight*2, "Nama PIC Pengunjung", "LTRB", 0, "LM", true, 0, "")
	pdf.CellFormat(40, cellHeight*2, "Nama Kegiatan", "LTRB", 0, "LM", true, 0, "")
	pdf.Ln(-1)
	pdf.SetTextColor(0, 0, 0)
	//table
	for i := 0; i < len(dataKegiatan); i++ {
		if i%2 == 0 {
			pdf.SetFillColor(255, 255, 255)
		} else {
			pdf.SetFillColor(251, 228, 213)
		}
		pdf.CellFormat(10, cellHeight, strconv.Itoa(i+1), "LTRB", 0, "CM", true, 0, "")
		pdf.CellFormat(25, cellHeight, dataKegiatan[i].TanggalMulai, "LTRB", 0, "LM", true, 0, "")
		pdf.CellFormat(40, cellHeight, dataKegiatan[i].DetailDataCenter.Nama_dc, "LTRB", 0, "LM", true, 0, "")
		pdf.CellFormat(40, cellHeight, dataKegiatan[i].NamaPerusahaan, "LTRB", 0, "LM", true, 0, "")
		pdf.CellFormat(45, cellHeight, dataKegiatan[i].NamaPicVendor, "LTRB", 0, "LM", true, 0, "")
		pdf.CellFormat(40, cellHeight, dataKegiatan[i].DetailKegiatan.Nama_kegiatan, "LTRB", 0, "LM", true, 0, "")
		pdf.Ln(-1)
	}

	directory := os.Getenv("EXP_PDF_PATH") + "/"
	urlName := directory + "filenname_kunjungan.pdf"
	if err := pdf.OutputFileAndClose(urlName); err != nil {
		return "", err
	}

	return "filenname_kunjungan.pdf", nil
}

func (m *usecase) GetLaporanKegiatan(tahun, bulan, id_datacenter int64, tanggal string) (string, error) {
	dataKegiatan, err := m.reportRepo.GetReportKegiatanDc(tahun, bulan, id_datacenter, tanggal)
	if err != nil {
		return "", err
	}

	var (
		baseFont   string  = "Arial"
		cellHeight float64 = 5.5
	)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(5, 10, 5)
	pdf.SetFont(baseFont, "", 7)
	pdf.SetDrawColor(237, 125, 49)
	pdf.SetFillColor(237, 125, 49)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetAutoPageBreak(true, 10)

	pdf.SetFooterFunc(func() {
		pdf.SetY(-12)
		pdf.SetFont(baseFont, "I", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("Halaman %d dari {nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
	})

	pdf.AliasNbPages("")
	pdf.AddPage()

	pdf.SetFont(baseFont, "B", 11)
	pdf.CellFormat(200, cellHeight, "Laporan Kegiatan Data Center", "", 1, "CM", false, 0, "")
	pdf.CellFormat(200, cellHeight, "Tanggal : "+dataKegiatan[0].TanggalMulai+" - "+dataKegiatan[0].TanggalSelesai, "", 1, "CM", false, 0, "")

	pdf.Ln(7)

	// pdf.SetFont(baseFont, "B", 9.5)
	// pdf.SetTextColor(0, 0, 0)

	pdf.CellFormat(200, cellHeight, "Kegiatan : "+dataKegiatan[0].DetailKegiatan.Nama_kegiatan, "", 0, "LM", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(200, cellHeight, "Pelaksana kegiatan : "+dataKegiatan[0].NamaPicVendor+" dari perusahaan/instansi : "+dataKegiatan[0].NamaPerusahaan, "", 0, "LM", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(200, cellHeight, "Deskripsi kegiatan : "+dataKegiatan[0].Deskripsi, "", 0, "LM", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(200, cellHeight, "Resiko kegiatan apabila gagal : "+dataKegiatan[0].Resiko, "", 0, "LM", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(200, cellHeight, "Hasil kegiatan : "+dataKegiatan[0].Hasil, "", 0, "LM", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(200, cellHeight, "Status Kegiatan : "+dataKegiatan[0].Status, "", 0, "LM", false, 0, "")
	pdf.Ln(-1)
	pdf.Ln(-1)
	pdf.Ln(-1)
	pdf.Ln(-1)
	pdf.CellFormat(50, cellHeight, "Petugas 1", "", 0, "LM", false, 0, "")
	pdf.CellFormat(75, cellHeight, "", "", 0, "LM", false, 0, "")
	pdf.CellFormat(50, cellHeight, "Petugas 2", "", 0, "LM", false, 0, "")
	pdf.Ln(-7)
	pdf.Ln(-1)
	pdf.Ln(-1)
	pdf.Ln(-1)
	pdf.Ln(-1)
	pdf.Ln(-1)
	pdf.Ln(-1)
	pdf.CellFormat(50, cellHeight, dataKegiatan[0].DetailUser.Nama, "", 0, "LM", false, 0, "")
	pdf.CellFormat(75, cellHeight, "", "", 0, "LM", false, 0, "")
	pdf.CellFormat(50, cellHeight, dataKegiatan[0].DetailUserTwo.Nama, "", 0, "LM", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(50, cellHeight, "NIP : "+dataKegiatan[0].DetailUser.Nip, "", 0, "LM", false, 0, "")
	pdf.CellFormat(75, cellHeight, "", "", 0, "LM", false, 0, "")
	pdf.CellFormat(50, cellHeight, "NIP : "+dataKegiatan[0].DetailUserTwo.Nip, "", 0, "LM", false, 0, "")

	directory := os.Getenv("EXP_PDF_PATH") + "/"
	urlName := directory + "filenname_kegiatan.pdf"
	if err := pdf.OutputFileAndClose(urlName); err != nil {
		return "", err
	}

	return "filenname_kegiatan.pdf", nil
}

func (m *usecase) GetLaporanHarian(tahun, bulan, id_datacenter int64, tanggal string) (string, error) {
	data, err := m.reportRepo.GetAll(tahun, bulan, id_datacenter, tanggal)
	if err != nil {
		return "", err
	}

	var (
		baseFont   string = "Arial"
		imageOpt   gofpdf.ImageOptions
		cellHeight float64 = 5.5
	)

	imageOpt.ImageType = "png"

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(5, 10, 5)
	pdf.SetFont(baseFont, "", 7)
	pdf.SetDrawColor(237, 125, 49)
	pdf.SetFillColor(237, 125, 49)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetAutoPageBreak(true, 10)

	pdf.SetFooterFunc(func() {
		pdf.SetY(-12)
		pdf.SetFont(baseFont, "I", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("Halaman %d dari {nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
	})

	pdf.AliasNbPages("")
	pdf.AddPage()

	pdf.SetFont(baseFont, "B", 11)
	pdf.CellFormat(200, cellHeight, "Laporan Monitoring Harian ", "", 1, "CM", false, 0, "")

	pdf.Ln(7)

	pdf.SetFont(baseFont, "B", 9.5)
	pdf.SetTextColor(255, 255, 255)

	pdf.CellFormat(10, cellHeight*2, "No", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(30, cellHeight*2, "Tanggal", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(20, cellHeight*2, "Jam", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(40, cellHeight*2, "Data Center", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(30, cellHeight*2, "Lokasi", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(40, cellHeight*2, "Ruangan", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(30, cellHeight*2, "Kondisi", "LTRB", 0, "CM", true, 0, "")

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont(baseFont, "", 9)
	pdf.Ln(-1)
	pdf.SetTextColor(0, 0, 0)
	//table
	for i := 0; i < len(data); i++ {
		if i%2 == 0 {
			pdf.SetFillColor(255, 255, 255)
		} else {
			pdf.SetFillColor(251, 228, 213)
		}

		pdf.CellFormat(10, cellHeight, strconv.Itoa(i+1), "LTRB", 0, "CM", true, 0, "")
		pdf.CellFormat(30, cellHeight, data[i].Tanggal, "LTRB", 0, "LM", true, 0, "")
		pdf.CellFormat(20, cellHeight, data[i].Jam, "LTRB", 0, "LM", true, 0, "")
		pdf.CellFormat(40, cellHeight, data[i].DetailDataCenter.Nama_dc, "LTRB", 0, "LM", true, 0, "")
		pdf.CellFormat(30, cellHeight, data[i].DetailDataCenter.Lokasi, "LTRB", 0, "LM", true, 0, "")
		pdf.CellFormat(40, cellHeight, data[i].DetailRuangan.Nama_ruangan, "LTRB", 0, "LM", true, 0, "")
		pdf.CellFormat(30, cellHeight, data[i].Kondisi, "LTRB", 0, "LM", true, 0, "")
		pdf.Ln(-1)
	}

	directory := os.Getenv("EXP_PDF_PATH") + "/"
	urlName := directory + "filenname.pdf"
	if err := pdf.OutputFileAndClose(urlName); err != nil {
		return "", err
	}

	return "filenname.pdf", nil
}
