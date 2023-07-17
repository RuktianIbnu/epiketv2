package monitoring_harian

import (
	"epiketv2/pkg/model"
	rr "epiketv2/pkg/repository/report"
	"fmt"
"strconv"
"strconv"
"strconv"

	/ "strconv"

"github.com/jug-kurt/gofpdf"
)

/

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
	pdf.CellFormat(200, cellHeight, "Laporan Monitoring Harian ", "", 1, "CM", false, 0, "")

	pdf.Ln(7)

	pdf.SetFont(baseFont, "B", 9.5)
	pdf.SetTextColor(255, 255, 255)

	pdf.CellFormat(10, cellHeight*2, "No", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(20, cellHeight*2, "Tanggal", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(20, cellHeight*2, "Jam", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(40, cellHeight*2, "Data Center", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(40, cellHeight*2, "Lokasi", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(40, cellHeight*2, "Ruangan", "LTRB", 0, "CM", true, 0, "")
	pdf.CellFormat(30, cellHeight*2, "Kondisi", "LTRB", 0, "CM", true, 0, "")

	
	urlName := directory + filenname
	if err := pdf.OutputFileAndClose(urlName); err != nil {
		return "", err
	}

	return filenname, nil
}