package report

import (
	rpt "epiketv2/http/usecase/laporan"
	resp "epiketv2/pkg/helper/response"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler interface {
	GetReportHarian(c *gin.Context)
	GetReportKegiatanDc(c *gin.Context)
	GetReportKunjungan(c *gin.Context)
}

type handler struct {
	// dashboardUsecase dh.Usecase
	reportPage rpt.Usecase
}

// NewHandler ...
func NewHandler() Handler {
	return &handler{
		rpt.NewUsecase(),
	}
}

func (m *handler) GetReportKunjungan(c *gin.Context) {
	type RequestFilter struct {
		Tahun        int64  `json:"tahun,omitempty"`
		Tanggal      string `json:"tanggal"`
		IdDataCenter int64  `json:"id_dc,omitempty"`
		Bulan        int64  `json:"bulan,omitempty"`
	}

	var (
		rf       RequestFilter
		tahun, _ = strconv.ParseInt(c.Query("tahun"), 10, 64)
		tanggal  = c.Query("tanggal")
		id_dc, _ = strconv.ParseInt(c.Query("id_dc"), 10, 64)
		bulan, _ = strconv.ParseInt(c.Query("bulan"), 10, 64)
	)

	rf.Bulan = bulan
	rf.IdDataCenter = id_dc
	rf.Tahun = tahun
	rf.Tanggal = tanggal

	list, err := m.reportPage.GetLaporanKunjungan(rf.Tahun, rf.Bulan, rf.IdDataCenter, rf.Tanggal)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}
	c.FileAttachment(fmt.Sprintf("%s/"+list, os.Getenv("EXP_PDF_PATH")), list)
}

func (m *handler) GetReportKegiatanDc(c *gin.Context) {
	type RequestFilter struct {
		Tahun        int64  `json:"tahun,omitempty"`
		Tanggal      string `json:"tanggal"`
		IdDataCenter int64  `json:"id_dc,omitempty"`
		Bulan        int64  `json:"bulan,omitempty"`
	}

	var (
		rf       RequestFilter
		tahun, _ = strconv.ParseInt(c.Query("tahun"), 10, 64)
		tanggal  = c.Query("tanggal")
		id_dc, _ = strconv.ParseInt(c.Query("id_dc"), 10, 64)
		bulan, _ = strconv.ParseInt(c.Query("bulan"), 10, 64)
	)

	rf.Bulan = bulan
	rf.IdDataCenter = id_dc
	rf.Tahun = tahun
	rf.Tanggal = tanggal

	fmt.Println(tahun, tanggal, id_dc, bulan)

	list, err := m.reportPage.GetLaporanKegiatan(rf.Tahun, rf.Bulan, rf.IdDataCenter, rf.Tanggal)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}
	c.FileAttachment(fmt.Sprintf("%s/"+list, os.Getenv("EXP_PDF_PATH")), list)
}

func (m *handler) GetReportHarian(c *gin.Context) {
	type RequestFilter struct {
		Tahun        int64  `json:"tahun,omitempty"`
		Tanggal      string `json:"tanggal"`
		IdDataCenter int64  `json:"id_dc,omitempty"`
		Bulan        int64  `json:"bulan,omitempty"`
	}

	var (
		rf       RequestFilter
		tahun, _ = strconv.ParseInt(c.Query("tahun"), 10, 64)
		tanggal  = c.Query("tanggal")
		id_dc, _ = strconv.ParseInt(c.Query("id_dc"), 10, 64)
		bulan, _ = strconv.ParseInt(c.Query("bulan"), 10, 64)
	)

	rf.Bulan = bulan
	rf.IdDataCenter = id_dc
	rf.Tanggal = tanggal
	rf.Tahun = tahun

	fmt.Println(tahun, tanggal, id_dc, bulan)

	list, err := m.reportPage.GetLaporanHarian(rf.Tahun, rf.Bulan, rf.IdDataCenter, rf.Tanggal)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}
	c.FileAttachment(fmt.Sprintf("%s/"+list, os.Getenv("EXP_PDF_PATH")), list)
}
