package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(hostname:port)/dbname")
	res, err := db.Exec("INSERT INTO Ogrenci (isim, soyisim, sinif, no) VALUES (?, ?, ?, ?)", "Berkay", "Aykose", "Üniversite 3", 21253804)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d satır eklendi\n", rowCount)
}
