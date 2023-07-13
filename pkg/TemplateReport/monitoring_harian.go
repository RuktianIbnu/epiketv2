package monitoring_harian

import (
	"epiketv2/pkg/model"
	rr "epiketv2/pkg/repository/report"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
)

// Usecase ...
type Usecase interface {
	GetReportMonitoringHarian(dqp *model.DefaultQueryParam) (string, error)
}

type usecase struct {
	reportRepo rr.Repository
	data       *ReportHarian
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
	Data  []*ReportDataHarian
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

func (m *usecase) GetReportMonitoringHarian(dqp *model.DefaultQueryParam) (string, error) {
	data, err := m.reportRepo.GetAll(dqp)
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
	pdf.CellFormat(200, cellHeight, "Laporan Monitoring Harian "+dqp.Params["tahun"].(string)+dqp.Params["bulan"].(string), "", 1, "CM", false, 0, "")

	pdf.Ln(7)

	pdf.SetFont(baseFont, "B", 9.5)
	pdf.SetTextColor(255, 255, 255)

	pdf.CellFormat(10, cellHeight*2, "No", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(125, cellHeight*2, "Tanggal", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(20, cellHeight*2, "Jam", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(20, cellHeight*2, "Data Center", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(25, cellHeight*2, "Lokasi", "LTRB", 1, "CM", true, 0, "")
	pdf.CellFormat(25, cellHeight*2, "Ruangan", "LTRB", 1, "CM", true, 0, "")
	pdf.CellFormat(25, cellHeight*2, "Kondisi", "LTRB", 1, "CM", true, 0, "")

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont(baseFont, "", 9)

	//table
	for i := 0; i < len(data); i++ {
		if i%2 == 0 {
			pdf.SetFillColor(255, 255, 255)
		} else {
			pdf.SetFillColor(251, 228, 213)
		}

		pdf.CellFormat(10, cellHeight, strconv.Itoa(i+1), "LTRB", 0, "CM", true, 0, "")
		pdf.CellFormat(10, cellHeight, data[i].Tanggal.String(), "LTRB", 0, "LM", true, 0, "")
	}

	pdf.Ln(7)

	pdf.SetFont(baseFont, "B", 9)
	pdf.SetDrawColor(71, 71, 71)

	pdf.CellFormat(20, cellHeight, "Grade", "LTRB", 0, "CM", false, 0, "")
	pdf.CellFormat(25, cellHeight, "Keterangan", "LTRB", 0, "CM", false, 0, "")
	pdf.CellFormat(30, cellHeight, "Rentang Nilai", "LTRB", 1, "CM", false, 0, "")

	pdf.SetFont(baseFont, "", 9)

	pdf.CellFormat(20, cellHeight, "A", "LTRB", 0, "CM", false, 0, "")
	pdf.CellFormat(25, cellHeight, "Sangat Baik", "LTRB", 0, "CM", false, 0, "")
	pdf.CellFormat(30, cellHeight, "88.31 - 100", "LTRB", 1, "CM", false, 0, "")

	pdf.CellFormat(20, cellHeight, "B", "LTRB", 0, "CM", false, 0, "")
	pdf.CellFormat(25, cellHeight, "Baik", "LTRB", 0, "CM", false, 0, "")
	pdf.CellFormat(30, cellHeight, "76.61 - 88.3", "LTRB", 1, "CM", false, 0, "")

	pdf.CellFormat(20, cellHeight, "C", "LTRB", 0, "CM", false, 0, "")
	pdf.CellFormat(25, cellHeight, "Kurang Baik", "LTRB", 0, "CM", false, 0, "")
	pdf.CellFormat(30, cellHeight, "65 - 76.6", "LTRB", 1, "CM", false, 0, "")

	pdf.CellFormat(20, cellHeight, "D", "LTRB", 0, "CM", false, 0, "")
	pdf.CellFormat(25, cellHeight, "Tidak Baik", "LTRB", 0, "CM", false, 0, "")
	pdf.CellFormat(30, cellHeight, "25 - 64.99", "LTRB", 1, "CM", false, 0, "")

	pdf.Ln(18)

	pdf.ImageOptions("output.png", 5, 0, 200, 0, true, imageOpt, 0, "")
	directory := os.Getenv("EXP_PDF_PATH") + "/"
	filenname := strconv.FormatInt(time.Now().Unix(), 10) + "_" + dqp.Params["kode"].(string) + ".pdf"
	urlName := directory + filenname
	if err := pdf.OutputFileAndClose(urlName); err != nil {
		return "", err
	}

	return filenname, nil
}
