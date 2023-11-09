package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sasmeka/user_product_management_duazsolusi/config"
)

func Routers(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(config.CorsConfig))

	auth(router, db)
	users(router, db)
	products(router, db)

	return router
}
