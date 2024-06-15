package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username, Password string
}

func main() {
	dsn := "root:123456789@tcp(127.0.0.1:3306)/blog_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Şema migrate edilerek veritabanında users tablosu oluşturulur:
	//db.AutoMigrate(&User{})

	// Yeni bir User(kullanıcı) kaydı oluşturulur:
	//db.Create(&User{Username: "admin", Password: "sifre123"})

	// Bir ürün kaydı okunur:
	//var user User

	// İlk olarak, id si 1 olan ürün kaydı bulur:
	//db.First(&user, 1)
	//fmt.Println(user.Username)

	// Sonra, Username "admin" olan ürün kaydı bulur:
	//db.First(&user, "Username = ?", "admin")
	//fmt.Println(user.Password)

	//birden fazla User kaydı bulur:
	//var users []User
	//db.Find(&users, "Password = ?", "sifre123")
	//fmt.Println(users)

	//Id si 1 olan Userın Username sütununu "updateAdmin" olarak güncelle
	//db.First(&user, 1)
	//db.Model(&user).Update("Username", "updateAdmin")

	//Id si 2 olan Userın birden çok  sütununu güncelleme
	//db.First(&user, 2)
	//db.Model(&user).Updates(User{Username: "Berkay", Password: "123"})

	//Id si 3 olan Userı silme
	db.Delete(&User{}, 3)
}
