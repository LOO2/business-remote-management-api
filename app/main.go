package main

import (
	"github.com/LOO2/business-remote-management-api/database"
	"github.com/LOO2/business-remote-management-api/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
