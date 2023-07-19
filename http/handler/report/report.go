package report

import (
	rpt "epiketv2/http/usecase/laporan"
	resp "epiketv2/pkg/helper/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler interface {
	GetReportHarian(c *gin.Context)
	GetReportKegiatanDc(c *gin.Context) //changed
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

func (m *handler) GetReportKegiatanDc(c *gin.Context) {
	type RequestFilter struct {
		Tahun int64  `json:"tahun,omitempty"`
		Kode  string `json:"kode"`
	}

	var (
		ReqFilter RequestFilter
	)

	if err := c.ShouldBindJSON(&ReqFilter); err != nil {
		c.JSON(resp.Format(http.StatusBadRequest, err))
		return
	}

	list, err := m.reportPage.GetReportKegiatanDc(ReqFilter.Tahun, ReqFilter.Kode)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}
	fmt.Println(list)
	// c.FileAttachment(fmt.Sprintf("%s/"+list), list)
	c.JSON(resp.Format(http.StatusOK, nil, gin.H{"message": "berhasil"}))
}

func (m *handler) GetReportHarian(c *gin.Context) {
	type RequestFilter struct {
		Tahun int64  `json:"tahun,omitempty"`
		Kode  string `json:"kode"`
	}

	var (
		ReqFilter RequestFilter
	)

	if err := c.ShouldBindJSON(&ReqFilter); err != nil {
		c.JSON(resp.Format(http.StatusBadRequest, err))
		return
	}

	list, err := m.reportPage.GetReportMonitoringHarian(ReqFilter.Tahun, ReqFilter.Kode)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}
	fmt.Println(list)
	// c.FileAttachment(fmt.Sprintf("%s/"+list), list)
	c.JSON(resp.Format(http.StatusOK, nil, gin.H{"message": "berhasil"}))
}
