package main

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/admin"
	"github.com/qor/qor/media_library"
)

type CreditCard struct {
	Id     int64
	Number string
	Issuer string
}

type Address struct {
	Id       int
	UserId   int64
	Address1 string
	Address2 string
}

type Role struct {
	Id   int
	Name string
}

type Language struct {
	Id   int
	Name string
}

type User struct {
	Id           int64
	Name         string
	Gender       string
	Description  string
	File         media_library.FileSystem
	RoleId       int64
	Languages    []Language `gorm:"many2many:user_languages;"`
	CreditCard   CreditCard
	CreditCardId int64
	Addresses    []Address
}

var db gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "/tmp/qor.db")
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.AutoMigrate(&User{}, &CreditCard{}, &Address{}, &Role{}, &Language{}, &admin.AssetManager{})

	db.FirstOrCreate(&Role{}, Role{Name: "admin"})
	db.FirstOrCreate(&Role{}, Role{Name: "dev"})
	db.FirstOrCreate(&Role{}, Role{Name: "customer_support"})

	db.FirstOrCreate(&Language{}, Language{Name: "CN"})
	db.FirstOrCreate(&Language{}, Language{Name: "JP"})
	db.FirstOrCreate(&Language{}, Language{Name: "EN"})
	db.FirstOrCreate(&Language{}, Language{Name: "DE"})
}
