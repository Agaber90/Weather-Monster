package main

import (
	_cityHTTP "Weather-Monster/city/delivery-http"
	_cityRepo "Weather-Monster/city/repository"
	_cityUsecase "Weather-Monster/city/usecase"

	_tempHTTP "Weather-Monster/temperature/delivery-http"
	_tempRepo "Weather-Monster/temperature/repository"
	_tempUsecase "Weather-Monster/temperature/usecase"

	_forcastHTTP "Weather-Monster/forecast/delivery-http"
	_forecastRepo "Weather-Monster/forecast/repository"
	_forecastUsecase "Weather-Monster/forecast/usecase"

	_webhooksHTTP "Weather-Monster/webhook/delivery-http"
	_webhookRepo "Weather-Monster/webhook/repository"
	_webhookUsecase "Weather-Monster/webhook/usecase"

	"Weather-Monster/middleware"
	"database/sql"

	"time"

	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUsr := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsr, dbPass, dbHost, dbPort, dbName)
	dbConn, err := sql.Open("mysql", connection)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := middleware.InitMiddleware()

	e.Use(middL.CORS)

	cityRepo := _cityRepo.NewCityRepository(dbConn)
	tempereatureRepo := _tempRepo.NewTempRepository(dbConn)
	foreRepo := _forecastRepo.NewForecaseRepository(dbConn)
	webhookRepo := _webhookRepo.NewWebhookRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	ctyUsecase := _cityUsecase.NewCityUseCase(cityRepo, timeoutContext)
	tempUsecase := _tempUsecase.NewTempUseCase(tempereatureRepo, timeoutContext)
	foreUsecase := _forecastUsecase.NewForecastUsecase(foreRepo, timeoutContext)
	webhookUsecase := _webhookUsecase.NewWebhookUsecase(webhookRepo, timeoutContext)
	_cityHTTP.NewCityHandler(e, ctyUsecase)
	_tempHTTP.NewTempHandler(e, tempUsecase)
	_forcastHTTP.NewForecastHandler(e, foreUsecase)
	_webhooksHTTP.NewWebhookHandler(e, webhookUsecase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
