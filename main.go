package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/murali0706/customer/config"
	customerHttp "github.com/murali0706/customer/customer/http"
	repository "github.com/murali0706/customer/repository/mysql"
	"github.com/murali0706/customer/usecase"
)

func main() {
	dbConn, err := config.GetDB()

	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	customerRepo := repository.NewMysqlCustomerReposity(dbConn)

	au := usecase.NewCustomerUsecase(customerRepo)
	customerHttp.NewCustomerHandler(e, au)

	log.Fatal(e.Start(":8000"))
}
