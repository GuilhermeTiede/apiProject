package main

func main() {
	dbConfig := DatabaseConfig{
		host:     "localhost",
		port:     5434,
		user:     "postgres",
		password: "postgres",
		dbname:   "loginapi",
	}

	apiServer := NewAPIServer("localhost:8080", dbConfig)
	apiServer.Start()
}
