package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Paulo-Lopes-Estevao/vitebank/domain"
	"github.com/Paulo-Lopes-Estevao/vitebank/infra/repository"
	"github.com/Paulo-Lopes-Estevao/vitebank/usecase"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func main() {
	db := setupDB()
	defer db.Close()
	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "lopes"
	cc.ExpirationYear = 2021
	cc.ExpirationMonth = 7
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}

}

func setupDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("error connection to database n/", err)
	}

	return db
}

func setupTransactionUseCase(db *gorm.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	return useCase
}
