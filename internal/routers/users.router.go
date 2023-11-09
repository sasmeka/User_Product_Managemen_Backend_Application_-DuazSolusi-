package routers

import (
	"github.com/sasmeka/user_product_management_duazsolusi/internal/handlers"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func users(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repositories.New_Users(d)
	handler := handlers.New_Users(repo)

	route.GET("/byId/:id", handler.Get_Data_Users_byId)
	route.GET("/", handler.Get_Data_Users)
	route.POST("/", handler.Post_Data_User)
	route.PUT("/:id", handler.Put_Data_User)
	route.DELETE("/:id", handler.Delete_Data_User)

}
