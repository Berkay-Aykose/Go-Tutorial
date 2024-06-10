package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("sql", "user:password@/dbname")
	res, err := db.Exec("INSERT INTO Ogrenci (isim, soyisim, sinif, no) VALUES (?, ?)", "Berkay", "Aykose", "Üniversite 3", 21253804)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Etkileşimli satır sayısını alın
	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d satır eklendi\n", rowCount)
}
