package report

import (
	resp "epiketv2/pkg/helper/response"
	rpt "epiketv2/pkg/templatereport"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler interface {
	GetReportHarian(c *gin.Context) //changed
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

// func (m *handler) GetReportHarian(c *gin.Context) {
// 	var (
// 		dq = qry.Q{
// 			Ctx: c,
// 		}
// 		TanggalMulai   = c.Query("tanggal_mulai")
// 		TanggalSelesai = c.Query("tanggal_selesai")
// 		Tahun          = c.Query("tahun")
// 		Bulan          = c.Query("bulan")
// 		Kode           = c.Query("kode")
// 	)

// 	dqp, err := dq.DefaultQueryParam()
// 	if err != nil {
// 		c.JSON(resp.Format(http.StatusInternalServerError, err))
// 		return
// 	}

// 	dqp.Params["tanggal_selesai"] = TanggalSelesai
// 	dqp.Params["tanggal_mulai"] = TanggalMulai
// 	dqp.Params["tahun"] = Tahun
// 	dqp.Params["bulan"] = Bulan
// 	dqp.Params["kode"] = Kode

// 	list, err := m.reportPage.GetReportMonitoringHarian(dqp)
// 	if err != nil {
// 		c.JSON(resp.Format(http.StatusInternalServerError, err))
// 		return
// 	}
// 	c.FileAttachment(fmt.Sprintf("%s/"+list), list)
// }

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
