package router

import (
	"github.com/gin-gonic/gin"

	dashboardHandler "epiketv2/http/handler/dashboard"
	dcHandler "epiketv2/http/handler/data_center"
	itemHandler "epiketv2/http/handler/item"
	kegiatanHandler "epiketv2/http/handler/kegiatan"
	ruanganHandler "epiketv2/http/handler/ruangan"
	strukturHandler "epiketv2/http/handler/struktur"
	txharianHandler "epiketv2/http/handler/tx_harian"
	txpiketHandler "epiketv2/http/handler/tx_piket"
	userHandler "epiketv2/http/handler/user"
	"epiketv2/http/middleware/auth"
	"epiketv2/http/middleware/cors"
)

// Routes ...
func Routes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Middleware())

	// for health check purpose only
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userHandler := userHandler.NewHandler()
	strukturHendler := strukturHandler.NewHandler()
	dcHendler := dcHandler.NewHandler()
	kegiatanHandler := kegiatanHandler.NewHandler()
	itemHandler := itemHandler.NewHandler()
	ruanganHandler := ruanganHandler.NewHandler()
	txpiketHandler := txpiketHandler.NewHandler()
	txharianHandler := txharianHandler.NewHandler()
	dashboardHandler := dashboardHandler.NewHandler()

	v1 := r.Group("/v1")
	{
		v1.POST("/login", userHandler.Login)
		v1.POST("/register", userHandler.Register)
		// v1.GET("/test-file", func(c *gin.Context) {
		// 	log.Println("oke")
		// 	c.FileAttachment(fmt.Sprintf("%s/report.pdf", os.Getenv("EXP_PDF_PATH")), "report.pdf")
		// })

		resources := v1.Group("/resources").Use(auth.Middleware())
		{

			resources.POST("/struktur", strukturHendler.Create)
			resources.GET("/struktur/:id", strukturHendler.GetOneByID)
			resources.PUT("/struktur/:id", strukturHendler.UpdateOneByID)
			resources.DELETE("/struktur/:id", strukturHendler.DeleteOneByID)
			resources.GET("/struktur", strukturHendler.GetAll)

			resources.POST("/user", userHandler.Create)
			resources.GET("/user/:id", userHandler.GetOneByID)
			resources.PUT("/user/:id", userHandler.UpdateOneByID)
			resources.DELETE("/user/:id", userHandler.DeleteOneByID)
			resources.GET("/user", userHandler.GetAll)

			resources.POST("/item", itemHandler.Create)
			resources.GET("/item/:id", itemHandler.GetOneByID)
			resources.PUT("/item/:id", itemHandler.UpdateOneByID)
			resources.DELETE("/item/:id", itemHandler.DeleteOneByID)
			resources.GET("/item", itemHandler.GetAll)
			resources.GET("/item-by-id-ruangan/:id", itemHandler.GetAllByIdRuangan)

			resources.POST("/kegiatan", kegiatanHandler.Create)
			resources.GET("/kegiatan/:id", kegiatanHandler.GetOneByID)
			resources.PUT("/kegiatan/:id", kegiatanHandler.UpdateOneByID)
			resources.DELETE("/kegiatan/:id", kegiatanHandler.DeleteOneByID)
			resources.GET("/kegiatan", kegiatanHandler.GetAll)

			resources.POST("/data_center", dcHendler.Create)
			resources.GET("/data_center/:id", dcHendler.GetOneByID)
			resources.PUT("/data_center/:id", dcHendler.UpdateOneByID)
			resources.DELETE("/data_center/:id", dcHendler.DeleteOneByID)
			resources.GET("/data_center", dcHendler.GetAll)

			resources.POST("/ruangan", ruanganHandler.Create)
			resources.GET("/ruangan/:id", ruanganHandler.GetOneByID)
			resources.GET("/ruangan-by-id-dc/:id", ruanganHandler.GetAllById_dc)
			resources.PUT("/ruangan/:id", ruanganHandler.UpdateOneByID)
			resources.DELETE("/ruangan/:id", ruanganHandler.DeleteOneByID)
			resources.GET("/ruangan", ruanganHandler.GetAll)

			resources.POST("/tx-piket", txpiketHandler.Create)
			resources.GET("/tx-piket/:id", txpiketHandler.GetOneByID)
			resources.PUT("/tx-piket/:id", txpiketHandler.UpdateOneByID)
			resources.DELETE("/tx-piket/:id", txpiketHandler.DeleteOneByID)
			resources.GET("/tx-piket", txpiketHandler.GetAll)

			resources.POST("/tx-harian", txharianHandler.Create)
			resources.GET("/tx-harian/:id", txharianHandler.GetOneByID)
			resources.PUT("/tx-harian/:id", txharianHandler.UpdateOneByID)
			resources.DELETE("/tx-harian/:id", txharianHandler.DeleteOneByID)
			resources.GET("/tx-harian", txharianHandler.GetAll)

			resources.GET("/dash-kegiatan", dashboardHandler.GetAll)
			resources.GET("/dash-kondisi-abnormal", dashboardHandler.GetAllKondisiAbnormal)
			resources.GET("/dash-status-pending", dashboardHandler.GetAllStatusPending)
			resources.GET("/dash-kunjungan", dashboardHandler.GetAllKunjungan)
			resources.GET("/dash-tamu", dashboardHandler.GetAllTamu)
		}
	}

	return r
}
