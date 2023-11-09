package main

import (
	"log"

	"github.com/asaskevich/govalidator"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/routers"
	"github.com/sasmeka/user_product_management_duazsolusi/pkg"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	database, err := pkg.Postgres_Database()
	if err != nil {
		log.Fatal(err)
	}
	router := routers.Routers(database)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
