package monitoring_harian

// import (
// 	"epiketv2/pkg/model"
// 	dh "epiketv2/pkg/repository/dashboard"
// 	"fmt"
// 	"math"
// 	"os"
// 	"strconv"
// 	"time"

// 	"github.com/jung-kurt/gofpdf"
// 	"github.com/wcharczuk/go-chart"
// )

// // Usecase ...
// type Usecase interface {
// 	GetReportMonitoringHarian(dqp *model.DefaultQueryParam) (string, error)
// }

// type usecase struct {
// 	dashboardRepo dh.Repository
// 	data          *ReportIndeksKepuasanInternal
// }

// // NewUsecase ...
// func NewUsecase() Usecase {
// 	return &usecase{
// 		dashboardRepo: dh.NewRepository(),
// 	}
// }

// // ReportIndeksKepuasanInternal ...
// type ReportIndeksKepuasanInternal struct {
// 	Title string
// 	Data  []*ReportIndeksKepuasanInternalReportData
// }

// // ReportIndeksKepuasanInternalReportData ...
// type ReportIndeksKepuasanInternalReportData struct {
// 	SatuanKerja string
// 	Index       float64
// 	Grade       string
// 	Keterangan  string
// }

// func (m *usecase) GetReportMonitoringHarian(dqp *model.DefaultQueryParam) (string, error) {
// 	data, err := m.dashboardRepo.GetAllKunjungan(dqp)
// 	if err != nil {
// 		return "", err
// 	}
// 	fmt.Println(data)
// 	chart.DefaultBackgroundColor = chart.ColorTransparent
// 	chart.DefaultCanvasColor = chart.ColorTransparent

// 	var dataBar []chart.Value
// 	for i := 0; i < len(data); i++ {
// 		dataRecord := chart.Value{
// 			Value: data[i].AccumulateAvrage,
// 			Label: data[i].NamaInstansi,
// 		}
// 		dataBar = append(dataBar, dataRecord)
// 	}

// 	var title string
// 	if dqp.Params["metriks_penilaian"] == "iki" {
// 		title = "Visualisasi Indeks Kepuasan Internal "
// 	} else if dqp.Params["metriks_penilaian"] == "ikk" {
// 		title = "Visualisasi Indeks Kualitas Kebijakan "
// 	} else {
// 		title = "Visualisasi Indeks Pengamanan Keimigrasian "
// 	}
// 	graph := chart.BarChart{
// 		Title: title + dqp.Params["year"].(string),
// 		YAxis: chart.YAxis{
// 			Range: &chart.ContinuousRange{
// 				Min: 0.00,
// 				Max: 100.00,
// 			},
// 		},
// 		Background: chart.Style{
// 			Padding: chart.Box{
// 				Top: 40,
// 			},
// 		},
// 		Height:   512,
// 		BarWidth: chart.DefaultBarWidth,
// 		Bars:     dataBar,
// 	}

// 	f, _ := os.Create("output.png")
// 	defer f.Close()

// 	if err := graph.Render(chart.PNG, f); err != nil {
// 		panic(err)
// 	}

// 	var (
// 		baseFont   string = "Arial"
// 		imageOpt   gofpdf.ImageOptions
// 		cellHeight float64 = 5.5
// 	)

// 	imageOpt.ImageType = "png"

// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	pdf.SetMargins(5, 10, 5)
// 	pdf.SetFont(baseFont, "", 7)
// 	pdf.SetDrawColor(237, 125, 49)
// 	pdf.SetFillColor(237, 125, 49)
// 	pdf.SetTextColor(0, 0, 0)
// 	pdf.SetAutoPageBreak(true, 10)

// 	pdf.SetFooterFunc(func() {
// 		pdf.SetY(-12)
// 		pdf.SetFont(baseFont, "I", 8)
// 		pdf.CellFormat(0, 10, fmt.Sprintf("Halaman %d dari {nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
// 	})

// 	pdf.AliasNbPages("")
// 	pdf.AddPage()

// 	pdf.SetFont(baseFont, "B", 11)
// 	pdf.CellFormat(200, cellHeight, "Laporan Indeks Kepuasan Internal Tahun "+dqp.Params["year"].(string), "", 1, "CM", false, 0, "")

// 	pdf.Ln(7)

// 	pdf.SetFont(baseFont, "B", 9.5)
// 	pdf.SetTextColor(255, 255, 255)

// 	pdf.CellFormat(10, cellHeight*2, "No", "LTRB", 0, "CM", true, 0, "")
// 	pdf.CellFormat(125, cellHeight*2, "Satuan Kerja", "LTRB", 0, "CM", true, 0, "")
// 	pdf.CellFormat(20, cellHeight*2, "Indeks", "LTRB", 0, "CM", true, 0, "")
// 	pdf.CellFormat(20, cellHeight*2, "Grade", "LTRB", 0, "CM", true, 0, "")
// 	pdf.CellFormat(25, cellHeight*2, "Keterangan", "LTRB", 1, "CM", true, 0, "")

// 	pdf.SetTextColor(0, 0, 0)
// 	pdf.SetFont(baseFont, "", 9)
// 	// tabel
// 	for i := 0; i < len(data); i++ {
// 		if i%2 == 0 {
// 			pdf.SetFillColor(255, 255, 255)
// 		} else {
// 			pdf.SetFillColor(251, 228, 213)
// 		}
// 		average := strconv.FormatFloat(math.Round(data[i].AccumulateAvrage), 'f', 6, 64)
// 		pdf.CellFormat(10, cellHeight, strconv.Itoa(i+1), "LTRB", 0, "CM", true, 0, "")
// 		pdf.CellFormat(125, cellHeight, data[i].NamaInstansi, "LTRB", 0, "LM", true, 0, "")

// 		switch {
// 		case data[i].AccumulateAvrage >= 88.31:
// 			pdf.CellFormat(20, cellHeight, average, "LTRB", 0, "CM", true, 0, "")
// 			pdf.CellFormat(20, cellHeight, "A", "LTRB", 0, "CM", true, 0, "")
// 			pdf.CellFormat(25, cellHeight, "Sangat Baik", "LTRB", 1, "CM", true, 0, "")
// 		case data[i].AccumulateAvrage >= 76.61:
// 			pdf.CellFormat(20, cellHeight, average, "LTRB", 0, "CM", true, 0, "")
// 			pdf.CellFormat(20, cellHeight, "B", "LTRB", 0, "CM", true, 0, "")
// 			pdf.CellFormat(25, cellHeight, "Baik", "LTRB", 1, "CM", true, 0, "")
// 		case data[i].AccumulateAvrage >= 65:
// 			pdf.CellFormat(20, cellHeight, average, "LTRB", 0, "CM", true, 0, "")
// 			pdf.CellFormat(20, cellHeight, "C", "LTRB", 0, "CM", true, 0, "")
// 			pdf.CellFormat(25, cellHeight, "Kurang Baik", "LTRB", 1, "CM", true, 0, "")
// 		default:
// 			pdf.CellFormat(20, cellHeight, average, "LTRB", 0, "CM", true, 0, "")
// 			pdf.CellFormat(20, cellHeight, "D", "LTRB", 0, "CM", true, 0, "")
// 			pdf.CellFormat(25, cellHeight, "Tidak Baik", "LTRB", 1, "CM", true, 0, "")
// 		}

// 	}

// 	pdf.Ln(7)

// 	pdf.SetFont(baseFont, "B", 9)
// 	pdf.SetDrawColor(71, 71, 71)

// 	pdf.CellFormat(20, cellHeight, "Grade", "LTRB", 0, "CM", false, 0, "")
// 	pdf.CellFormat(25, cellHeight, "Keterangan", "LTRB", 0, "CM", false, 0, "")
// 	pdf.CellFormat(30, cellHeight, "Rentang Nilai", "LTRB", 1, "CM", false, 0, "")

// 	pdf.SetFont(baseFont, "", 9)

// 	pdf.CellFormat(20, cellHeight, "A", "LTRB", 0, "CM", false, 0, "")
// 	pdf.CellFormat(25, cellHeight, "Sangat Baik", "LTRB", 0, "CM", false, 0, "")
// 	pdf.CellFormat(30, cellHeight, "88.31 - 100", "LTRB", 1, "CM", false, 0, "")

// 	pdf.CellFormat(20, cellHeight, "B", "LTRB", 0, "CM", false, 0, "")
// 	pdf.CellFormat(25, cellHeight, "Baik", "LTRB", 0, "CM", false, 0, "")
// 	pdf.CellFormat(30, cellHeight, "76.61 - 88.3", "LTRB", 1, "CM", false, 0, "")

// 	pdf.CellFormat(20, cellHeight, "C", "LTRB", 0, "CM", false, 0, "")
// 	pdf.CellFormat(25, cellHeight, "Kurang Baik", "LTRB", 0, "CM", false, 0, "")
// 	pdf.CellFormat(30, cellHeight, "65 - 76.6", "LTRB", 1, "CM", false, 0, "")

// 	pdf.CellFormat(20, cellHeight, "D", "LTRB", 0, "CM", false, 0, "")
// 	pdf.CellFormat(25, cellHeight, "Tidak Baik", "LTRB", 0, "CM", false, 0, "")
// 	pdf.CellFormat(30, cellHeight, "25 - 64.99", "LTRB", 1, "CM", false, 0, "")

// 	pdf.Ln(18)

// 	pdf.ImageOptions("output.png", 5, 0, 200, 0, true, imageOpt, 0, "")
// 	directory := os.Getenv("EXP_PDF_PATH") + "/"
// 	filenname := strconv.FormatInt(time.Now().Unix(), 10) + "_" + dqp.Params["metriks_penilaian"].(string) + ".pdf"
// 	urlName := directory + filenname
// 	if err := pdf.OutputFileAndClose(urlName); err != nil {
// 		return "", err
// 	}

// 	return filenname, nil
// }
