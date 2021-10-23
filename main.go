package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/hanifbg/cud-category-product/config"
	"github.com/hanifbg/cud-category-product/handler"
	categoryHandler "github.com/hanifbg/cud-category-product/handler/category"
	productHandler "github.com/hanifbg/cud-category-product/handler/product"
	categoryRepo "github.com/hanifbg/cud-category-product/repository/category"
	"github.com/hanifbg/cud-category-product/repository/migration"
	productRepo "github.com/hanifbg/cud-category-product/repository/product"
	categoryService "github.com/hanifbg/cud-category-product/service/category"
	productService "github.com/hanifbg/cud-category-product/service/product"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newDatabaseConnection(config *config.AppConfig) *gorm.DB {

	configDB := map[string]string{
		"DB_Username": os.Getenv("YOUR_DB_USERNAME"),
		"DB_Password": os.Getenv("YOUR_DB_PASSWORD"),
		"DB_Port":     os.Getenv("YOUR_DB_PORT"),
		"DB_Host":     os.Getenv("YOUR_DB_ADDRESS"),
		"DB_Name":     os.Getenv("YOUR_DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])

	fmt.Println(connectionString)

	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	migration.InitMigrate(db)

	return db
}

func main() {
	config := config.GetConfig()

	dbConnection := newDatabaseConnection(config)

	categoryRepo := categoryRepo.NewGormDBRepository(dbConnection)
	categoryService := categoryService.NewService(categoryRepo)
	categoryHandler := categoryHandler.NewHandler(categoryService)

	productRepo := productRepo.NewGormDBRepository(dbConnection)
	productService := productService.NewService(productRepo)
	productHandler := productHandler.NewHandler(productService)

	e := echo.New()

	handler.RegisterPath(e, categoryHandler, productHandler)
	go func() {
		address := fmt.Sprintf("localhost:%d", config.AppPort)

		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
