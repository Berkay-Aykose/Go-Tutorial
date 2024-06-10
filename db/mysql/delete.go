package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(hostname:port)/dbname")
	res, err := db.Exec("DELETE FROM Ogrenci WHERE no=21253804")
	if err != nil {
		fmt.Println(err)
		return
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d satır Silindi\n", rowCount)
}
