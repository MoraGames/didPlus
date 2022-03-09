package main

import (
	"database/sql"
	"github.com/MoraGames/didPlus/go/internal/config"
	"github.com/MoraGames/didPlus/go/internal/database"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var cfg *config.Config
const configPath string = "configs/"

func init(){
	var err error

	log.Println("Reading configuration...")
	//READ CONFIG
	cfg, err = config.NewConfig(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	//VALIDATE CONFIG
	err = cfg.CheckConfig()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("CONFIG:", cfg)

	// Open database connection
	dbTmp, err := database.NewSqlDb(&cfg.Database)
	if err != nil {
		log.Fatalln(err)
	}
	db = dbTmp.Client

	// Set Environment Variables
	if v := os.Getenv("API_VERSION"); v == "" {
		os.Setenv("API_VERSION", "v1")
	}
}

func main() {
	// Make router
	r := routerInit()

	// Make Server
	server := &http.Server{
		Addr:              "0.0.0.0" + cfg.Server.Port,
		Handler:           r,
	}

	// Start server
	log.Println("Listening on 0.0.0.0" + cfg.Server.Port)
	log.Fatalln(server.ListenAndServe())
}
