package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/baskoroi/balance-sync/internal/dto"
	"github.com/baskoroi/balance-sync/internal/model"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=127.0.0.1 user=gorm password=gorm dbname=balance_sync port=%s sslmode=disable", os.Getenv("POSTGRES_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{}, &model.Transaction{}, &model.BalanceSnapshot{})
	return db, nil
}

func InitializeRoutes(db *gorm.DB) {
	e := echo.New()
	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "OK",
			"dc":     os.Getenv("DC_NAME"),
		})
	})

	e.POST("/users/fake", func(c echo.Context) error {
		user := new(model.User)
		if err := faker.FakeData(user, options.WithFieldsToIgnore("ID", "CreatedAt", "UpdatedAt")); err != nil {
			log.Println("ERROR faking data:", err)
			return err
		}

		if err := db.Create(user).Error; err != nil {
			log.Println("ERROR creating new user:", err)
			return err
		}
		return c.JSON(http.StatusOK, dto.OKResponse{})
	})

	e.POST("/deposit", func(c echo.Context) error {
		r := new(dto.DepositRequest)
		if err := c.Bind(r); err != nil {
			log.Println("ERROR binding deposit request:", err)
			return c.JSON(
				http.StatusBadRequest,
				dto.ErrorResponse{Message: "Invalid request body!", Error: err},
			)
		}

		db.Create(&model.Transaction{
			UserID: r.UserID,
			Amount: r.Amount,
			Type:   model.TransactionTypeDeposit,
		})
		return c.JSON(http.StatusOK, dto.OKResponse{})
	})

	// TODO implement spending credits

	e.Logger.Fatal(e.Start(os.Getenv("DC_ADDRESS")))
}

func main() {
	db, err := InitializeDB()
	if err != nil {
		log.Fatalln("Failed to connect/initialize DB", err)
	}

	InitializeRoutes(db)
}
