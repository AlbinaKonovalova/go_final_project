package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AlbinaKonovalova/go_final_project/db"
	"github.com/AlbinaKonovalova/go_final_project/handlers"
	"github.com/AlbinaKonovalova/go_final_project/server"
	"github.com/AlbinaKonovalova/go_final_project/storage"
)

func main() {
	dbFilePath, err := db.GetDBFilePath()
	if err != nil {
		log.Fatalf("Failed to get DB file path: %v", err)
	}

	db.InitDB(dbFilePath)

	s := storage.NewStorage(db.DB)
	h := handlers.NewHandler(s)

	port := os.Getenv("TODO_PORT")
	if port == "" {
		fmt.Println("TODO_PORT not set, using default port 7540")
		port = "7540"
	}

	server.InitHandlers(h)

	log.Printf("Server starting on port %s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
