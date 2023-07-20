package report

import (
	rpt "epiketv2/http/usecase/laporan"
	resp "epiketv2/pkg/helper/response"
	"fmt"
	"net/http"
	"os"

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
		rf RequestFilter
	)

	if err := c.ShouldBindJSON(&rf); err != nil {
		c.JSON(resp.Format(http.StatusBadRequest, err))
		return
	}

	list, err := m.reportPage.GetLaporanKunjungan(rf.Tahun, rf.Bulan, rf.IdDataCenter, rf.Tanggal)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}
	fmt.Println(list)
	// c.FileAttachment(fmt.Sprintf("%s/"+list), list)
	// c.JSON(resp.Format(http.StatusOK, nil, gin.H{"message": "berhasil"}))
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
		rf RequestFilter
	)

	if err := c.ShouldBindJSON(&rf); err != nil {
		c.JSON(resp.Format(http.StatusBadRequest, err))
		return
	}

	list, err := m.reportPage.GetLaporanKegiatan(rf.Tahun, rf.Bulan, rf.IdDataCenter, rf.Tanggal)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}
	fmt.Println(list)
	// c.FileAttachment(fmt.Sprintf("%s/"+list), list)
	// c.JSON(resp.Format(http.StatusOK, nil, gin.H{"message": "berhasil"}))
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
		rf RequestFilter
	)

	if err := c.ShouldBindJSON(&rf); err != nil {
		c.JSON(resp.Format(http.StatusBadRequest, err))
		return
	}

	list, err := m.reportPage.GetLaporanHarian(rf.Tahun, rf.Bulan, rf.IdDataCenter, rf.Tanggal)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}
	fmt.Println(list)
	// c.FileAttachment(fmt.Sprintf("%s/"+list), list)
	// c.JSON(resp.Format(http.StatusOK, nil, gin.H{"message": "berhasil"}))
	c.FileAttachment(fmt.Sprintf("%s/"+list, os.Getenv("EXP_PDF_PATH")), list)
}
