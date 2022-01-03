package main

import "github.com/guswitch/spiderman-api.git/server"
import "github.com/guswitch/spiderman-api.git/database"

func main() {
    database.StartDB()
	server  := server.NewServer()

    server.Run()
}