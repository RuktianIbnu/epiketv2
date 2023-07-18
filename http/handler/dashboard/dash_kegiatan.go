package dash_kegiatan

import (
	du "epiketv2/http/usecase/dashboard"
	qry "epiketv2/pkg/helper/query"
	resp "epiketv2/pkg/helper/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler interface {
	GetAll(c *gin.Context)
	GetAllKondisiAbnormal(c *gin.Context)
	GetAllStatusPending(c *gin.Context)
	GetAllKunjungan(c *gin.Context)
	GetAllTamu(c *gin.Context)
}

type handler struct {
	dashKegiatanUsecase du.Usecase
}

// NewHandler ...
func NewHandler() Handler {
	return &handler{
		du.NewUsecase(),
	}
}

func (m *handler) GetAll(c *gin.Context) {
	var (
		dq = qry.Q{
			Ctx: c,
		}
		tahun = c.Query("tahun")
	)

	dqp, err := dq.DefaultQueryParam()
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	dqp.Params["tahun"] = tahun
	fmt.Println("dqp.Params = ", tahun)
	list, totalEntries, err := m.dashKegiatanUsecase.GetAll(dqp)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil, list, totalEntries, dqp.Page, dqp.Limit))
}

func (m *handler) GetAllKondisiAbnormal(c *gin.Context) {
	var (
		dq = qry.Q{
			Ctx: c,
		}
		tahun = c.Query("tahun")
	)

	dqp, err := dq.DefaultQueryParam()
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	dqp.Params["tahun"] = tahun
	fmt.Println("dqp.Params = ", tahun)
	list, totalEntries, err := m.dashKegiatanUsecase.GetAllKondisiAbnormal(dqp)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil, list, totalEntries, dqp.Page, dqp.Limit))
}

func (m *handler) GetAllStatusPending(c *gin.Context) {
	var (
		dq = qry.Q{
			Ctx: c,
		}
		tahun = c.Query("tahun")
	)

	dqp, err := dq.DefaultQueryParam()
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	dqp.Params["tahun"] = tahun
	fmt.Println("dqp.Params = ", tahun)
	list, totalEntries, err := m.dashKegiatanUsecase.GetAllStatusPending(dqp)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil, list, totalEntries, dqp.Page, dqp.Limit))
}

func (m *handler) GetAllKunjungan(c *gin.Context) {
	var (
		dq = qry.Q{
			Ctx: c,
		}
		tahun = c.Query("tahun")
	)

	dqp, err := dq.DefaultQueryParam()
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	dqp.Params["tahun"] = tahun
	fmt.Println("dqp.Params = ", tahun)
	list, totalEntries, err := m.dashKegiatanUsecase.GetAllKunjungan(dqp)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil, list, totalEntries, dqp.Page, dqp.Limit))
}

func (m *handler) GetAllTamu(c *gin.Context) {
	var (
		dq = qry.Q{
			Ctx: c,
		}
		tahun = c.Query("tahun")
	)

	dqp, err := dq.DefaultQueryParam()
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	dqp.Params["tahun"] = tahun
	fmt.Println("dqp.Params = ", tahun)
	list, totalEntries, err := m.dashKegiatanUsecase.GetAllTamu(dqp)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil, list, totalEntries, dqp.Page, dqp.Limit))
}
