package main

import (
	"crypto/rand"
	"crypto/rsa"
	"flag"
	"log"
	"os"

	"App-server/controller"
	"App-server/routes"

	"github.com/Gurpreet-Bacancy/bcl/dbconn"
	"github.com/Gurpreet-Bacancy/bcl/postgres"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	privateKey *rsa.PrivateKey
	conn       *dbconn.Postgres
	models     *postgres.Models
	app        *controller.Application
)

func init() {
	var err error

	rng := rand.Reader

	privateKey, err = rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}

	// Load Env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to DB
	conn, err = dbconn.NewPostgres(os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	// load Models
	models = postgres.NewModels(conn)
}

func main() {
	fibApp := fiber.New()

	runmigration := flag.Bool("runmigration", false, "run migration")
	runseeders := flag.Bool("runseeders", false, "run seeders")
	startserver := flag.Bool("startserver", false, "start server")
	flag.Parse()

	if *runmigration {
		dbconn.Initialmigration(os.Getenv("DB_URL"))
	}

	if *runseeders {
		dbconn.Seeder(os.Getenv("DB_URL"))
	}

	if *startserver {
		// initialize routes
		routes.InitializeRoutes(app, fibApp, privateKey)

		// starting server
		fibApp.Listen(":3000")
	}
}
