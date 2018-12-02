package main

import (
	db "wine/database"
)
func main() {
	defer db.Db.Close()
}