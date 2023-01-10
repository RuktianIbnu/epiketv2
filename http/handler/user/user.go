package user

import (
	uu "epiketv2/http/usecase/user"
	qry "epiketv2/pkg/helper/query"
	resp "epiketv2/pkg/helper/response"
	"epiketv2/pkg/model"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler interface {
	Create(c *gin.Context)
	GetOneByID(c *gin.Context)
	UpdateOneByID(c *gin.Context)
	DeleteOneByID(c *gin.Context)
	GetAll(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type handler struct {
	userUsecase uu.Usecase
}

// NewHandler ...
func NewHandler() Handler {
	return &handler{
		uu.NewUsecase(),
	}
}

func (m *handler) Login(c *gin.Context) {
	type login struct {
		Nip      string `json:"nip" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}

	var (
		loginData login
	)

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(resp.Format(http.StatusBadRequest, err))
		return
	}

	tempToken, err := m.userUsecase.Login(loginData.Nip, loginData.Password)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil, gin.H{"temp_token": tempToken}))
}

func (m *handler) Register(c *gin.Context) {
	var (
		data model.MsUser
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(resp.Format(http.StatusBadRequest, err))
		return
	}

	isExistNIP := m.userUsecase.CheckNIPExist(data.Nip)
	if isExistNIP {
		c.JSON(resp.Format(http.StatusOK, nil, gin.H{"message": "NIP sudah terdaftar!"}))
		return
	}

	lastID, err := m.userUsecase.Register(data.Nip, data.Nama, data.No_hp, data.Password, data.Id_struktur, data.Aktif, data.Id_role)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}
	data.ID = lastID

	c.JSON(resp.Format(http.StatusOK, nil, data))
}

func (m *handler) Create(c *gin.Context) {
	var (
		data model.MsUser
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(resp.Format(http.StatusBadRequest, err))
		return
	}

	isExistNIP := m.userUsecase.CheckNIPExist(data.Nip)
	if isExistNIP {
		c.JSON(resp.Format(http.StatusOK, nil, gin.H{"message": "NIP sudah terdaftar!"}))
		return
	}

	lastID, err := m.userUsecase.Create(&data)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}
	data.ID = lastID

	c.JSON(resp.Format(http.StatusOK, nil, data))
}

func (m *handler) UpdateOneByID(c *gin.Context) {
	var (
		data   model.MsUser
		ids, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(resp.Format(http.StatusBadRequest, err))
		return
	}

	if ids <= 0 {
		c.JSON(resp.Format(http.StatusBadRequest, errors.New("Provide a valid ID")))
		return
	}

	data.ID = ids

	_, err := m.userUsecase.UpdateOneByID(&data)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil, data))
}

func (m *handler) GetOneByID(c *gin.Context) {
	var (
		ids, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	)

	if ids <= 0 {
		c.JSON(resp.Format(http.StatusBadRequest, errors.New("Provide a valid ID")))
		return
	}

	data, err := m.userUsecase.GetOneByID(ids)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil, data))
}

func (m *handler) GetAll(c *gin.Context) {
	var (
		dq = qry.Q{
			Ctx:     c,
			Sorting: []string{"nama_struktur", "nip", "deskripsi"},
		}
	)

	dqp, err := dq.DefaultQueryParam()
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}
	list, totalEntries, err := m.userUsecase.GetAll(dqp)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil, list, totalEntries, dqp.Page, dqp.Limit))
}

func (m *handler) DeleteOneByID(c *gin.Context) {
	var (
		ids, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	)

	if ids <= 0 {
		c.JSON(resp.Format(http.StatusBadRequest, errors.New("Provide a valid ID")))
		return
	}

	_, err := m.userUsecase.DeleteOneByID(ids)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil))
}
