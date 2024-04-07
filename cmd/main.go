package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MiniKartV1/calc/internal/adapters/app/api"
	"github.com/MiniKartV1/calc/internal/adapters/core/calc"
	"github.com/MiniKartV1/calc/internal/adapters/framework/left/rest"
	"github.com/MiniKartV1/calc/internal/adapters/framework/right/db"
	"github.com/MiniKartV1/calc/internal/ports"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting main function.")

	defer os.Exit(0)
	var err error
	// Create a Gin router with default middleware: logger and recovery (crash-free) middleware
	errLoad := godotenv.Load("../.env")
	if errLoad != nil {
		log.Fatalf("Error loading .env file: %v", errLoad)
	}

	// portrs
	var coreAdapter ports.CalculationPort
	var dbaseAdapter ports.DBPort
	var appAdapter ports.APIPort
	var restAdapter ports.RESTPort

	DB_URI := os.Getenv("DB_URI")
	fmt.Println("db_uri", DB_URI)
	// dbaseDriver := "mongodb"
	dbaseAdapter = db.NewAdapter(DB_URI)
	fmt.Println("Connected to database")

	if err != nil {
		fmt.Println("Error connecting to database")
	}

	defer dbaseAdapter.CloseDBConnection()

	coreAdapter = calc.NewAdapter()
	appAdapter = api.NewAdapter(coreAdapter, dbaseAdapter)
	restAdapter = rest.NewAdapter(appAdapter)

	restAdapter.Run()

}
