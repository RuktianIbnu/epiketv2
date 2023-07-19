package harian

import (
	rr "epiketv2/pkg/repository/laporan"
	"fmt"
	"strconv"

	// "strconv"

	"github.com/jung-kurt/gofpdf"
)

// Usecase ...
type Usecase interface {
	GetReportMonitoringHarian(tahun int64, kode string) (string, error)
	GetReportKegiatanDc(tahun int64, kode string) (string, error)
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

func (m *usecase) GetReportKegiatanDc(tahun int64, kode string) (string, error) {
	dataKegiatan, err := m.reportRepo.GetReportKegiatanDc(tahun, kode)
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

	pdf.SetFont(baseFont, "B", 9.5)
	pdf.SetTextColor(255, 255, 255)

	// pdf.ImageOptions("output.png", 5, 0, 200, 0, true, imageOpt, 0, "")
	// filenname := strconv.FormatInt(time.Now().Unix(), 10) + "_" + dqp.Params["kode"].(string) + ".pdf"
	urlName := "filenname_kegiatan.pdf"
	if err := pdf.OutputFileAndClose(urlName); err != nil {
		return "", err
	}

	return "filenname_kegiatan.pdf", nil
}

func (m *usecase) GetReportMonitoringHarian(tahun int64, kode string) (string, error) {
	data, err := m.reportRepo.GetAll(tahun, kode)
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

	// pdf.ImageOptions("output.png", 5, 0, 200, 0, true, imageOpt, 0, "")
	// filenname := strconv.FormatInt(time.Now().Unix(), 10) + "_" + dqp.Params["kode"].(string) + ".pdf"
	urlName := "filenname.pdf"
	if err := pdf.OutputFileAndClose(urlName); err != nil {
		return "", err
	}

	return "filenname.pdf", nil
}
