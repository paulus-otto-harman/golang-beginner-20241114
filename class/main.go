package main

import (
	"20241114/class/router"
	"20241114/database"
	"log"
	"net/http"
)

type pointerString *string

func main() {
	db := database.DbOpen("20241114a")
	defer db.Close()
	r := router.NewRouter(db)

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
