package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type APIServer struct {
	listenAddress  string
	databaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func NewAPIServer(listenAddress string, dbConfig DatabaseConfig) *APIServer {
	return &APIServer{
		listenAddress:  listenAddress,
		databaseConfig: dbConfig,
	}
}

func (a *APIServer) Start() {
	// Configurar a string de conexão com o PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		a.databaseConfig.host, a.databaseConfig.port, a.databaseConfig.user, a.databaseConfig.password, a.databaseConfig.dbname)

	// Conectar ao PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// Testar a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server listening on %s\n", a.listenAddress)
}
