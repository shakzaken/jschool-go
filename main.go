package main

import (
	"database/sql"
	"jschool/handlers"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"
)



func main() {

	connectionString:= "host=localhost port=5432 dbname=jschool user=yakir password="
	con, err := sql.Open("pgx",connectionString)
	if(err != nil){
		log.Fatal("Cannot connect to the database",err)
	}
	defer con.Close()

	err = con.Ping()
	if(err != nil){
		log.Fatal("Cannot ping database",err)
	}

	appRepo := handlers.NewAppRepo(con)



	

	r := Routers(appRepo);

	port:= ":3000"
	log.Printf("Server runs on port %s",port)
	http.ListenAndServe(port,r)

	

}
