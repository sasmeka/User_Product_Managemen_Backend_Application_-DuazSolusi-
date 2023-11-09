package routers

import (
	"github.com/sasmeka/user_product_management_duazsolusi/internal/handlers"
	middleware "github.com/sasmeka/user_product_management_duazsolusi/internal/middlewares"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func products(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	repo := repositories.New_Products(d)
	handler := handlers.New_Products(repo)

	route.GET("/", handler.Get_Data_Products)
	route.POST("/", middleware.AuthJwt("user"), handler.Post_Data_Product)
	route.PUT("/:id", middleware.AuthJwt("user"), handler.Put_Data_Product)
	route.DELETE("/:id", middleware.AuthJwt("user"), handler.Delete_Data_Product)

}
