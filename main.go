package main

import (
	"bab6_golang/database"
	"bab6_golang/routers"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "acsl2024"
	dbName   = "mcs_bab6_kerend"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	var PORT = ":8080" //Menentukan port
	psqlInfo := fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`, host, port, user, password, dbName,
	)

	// State penanganan error jika tidak berhasil membuat Postgre
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error Open DB: %v\n", err)
	}

	database.DBMigrate(DB)
	defer DB.Close()
	routers.StartServer().Run(PORT)  // Melakukan start server dengan port yang sudah ditentukan pada PORT 8080
	fmt.Printf("Berhasil Terhubung") // Pesan Jika sudah berhasil terhubung
}
