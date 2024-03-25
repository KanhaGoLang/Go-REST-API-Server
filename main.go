package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	id        int
	name      string
	inventory int
	price     int
}

func main() {

	connectionString := "root:password@tcp(127.0.0.1:3306)/sanjeev"

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p Product

		rows.Scan(&p.id, &p.name, &p.inventory, &p.price)

		fmt.Println(p)
	}

}
