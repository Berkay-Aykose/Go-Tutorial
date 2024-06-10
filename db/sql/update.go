package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("sql", "user:password@/dbname")
	res, err := db.Exec("UPDATE Ogrenci SET isim='Berkay' AND soyisim = 'Aykose' AND sinif='Üniversite 3' WHERE no=21253804")
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

	fmt.Printf("%d satır düzenlendi\n", rowCount)
}
