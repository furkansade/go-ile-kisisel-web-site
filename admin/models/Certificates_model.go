package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Certificate struct {
	gorm.Model
	Title, Description, Picture_url, Filter string
}

func (certificate Certificate) Migrate() {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&certificate)
}

func (certificate Certificate) Get(where ...interface{}) Certificate {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return certificate
	}
	db.First(&certificate, where...)
	return certificate
}

func (certificate Certificate) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Model(&certificate).Update(column, value)
}

func (certificate Certificate) Updates(data Certificate) {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Model(&certificate).Updates(data)
}

func (certificate Certificate) Delete() {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&certificate, certificate.ID)
}

func (certificate Certificate) GetAll(where ...interface{}) []Certificate {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var certificates []Certificate
	db.Find(&certificates, where...)
	return certificates
}

func (certificate Certificate) Add() {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&certificate)
}
