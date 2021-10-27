package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/hanifbg/cud-category-product/config"
	"github.com/hanifbg/cud-category-product/handler"
	cartHandler "github.com/hanifbg/cud-category-product/handler/cart"
	categoryHandler "github.com/hanifbg/cud-category-product/handler/category"
	checkoutHandler "github.com/hanifbg/cud-category-product/handler/checkout"
	productHandler "github.com/hanifbg/cud-category-product/handler/product"
	cartRepo "github.com/hanifbg/cud-category-product/repository/cart"
	categoryRepo "github.com/hanifbg/cud-category-product/repository/category"
	checkoutRepo "github.com/hanifbg/cud-category-product/repository/checkout"
	"github.com/hanifbg/cud-category-product/repository/migration"
	productRepo "github.com/hanifbg/cud-category-product/repository/product"
	cartService "github.com/hanifbg/cud-category-product/service/cart"
	categoryService "github.com/hanifbg/cud-category-product/service/category"
	checkoutService "github.com/hanifbg/cud-category-product/service/checkout"
	productService "github.com/hanifbg/cud-category-product/service/product"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newDatabaseConnection(config *config.AppConfig) *gorm.DB {

	configDB := map[string]string{
		"DB_Username": config.DbUsername,
		"DB_Password": config.DbPassword,
		"DB_Port":     strconv.Itoa(config.DbPort),
		"DB_Host":     config.DbAddress,
		"DB_Name":     config.DbName,
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

	cartRepo := cartRepo.NewGormDBRepository(dbConnection)
	cartService := cartService.NewService(cartRepo, productRepo)
	cartHandler := cartHandler.NewHandler(cartService)

	checkoutRepo := checkoutRepo.NewGormDBRepository(dbConnection)
	checkoutService := checkoutService.NewService(checkoutRepo, cartRepo)
	checkoutHandler := checkoutHandler.NewHandler(checkoutService)

	e := echo.New()

	handler.RegisterPath(e, categoryHandler, productHandler, cartHandler, checkoutHandler)
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
