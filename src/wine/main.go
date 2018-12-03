package main

import (
	db "wine/database"
	"wine/router"
)
func main() {
	defer db.Db.Close()
	router := router.InitRouter()
	router.Run(":10086")
}