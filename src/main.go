package main

import (
	"fmt"
	"github.com/MikAoJk/go-crud-rest-api/src/db"
	"github.com/MikAoJk/go-crud-rest-api/src/router"
)

func main() {
	db.InitPostgresDB()
	err := router.InitRouter().Run()
	if err != nil {
		fmt.Println(err)
		return
	}

}
