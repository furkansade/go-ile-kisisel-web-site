package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repo struct {
	gorm.Model
	Title, Description, Repo_Url, Icon string
}

func (repo Repo) Migrate() {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&repo)
}

func (repo Repo) Get(where ...interface{}) Repo {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return repo
	}
	db.First(&repo, where...)
	return repo
}

func (repo Repo) GetAll(where ...interface{}) []Repo {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var repos []Repo
	db.Find(&repos, where...)
	return repos
}

func (repo Repo) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&repo).Update(column, value)
}

func (repo Repo) Updates(data Repo) {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&repo).Updates(data)
}

func (repo Repo) Delete() {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Delete(&repo, repo.ID)
}

func (repo Repo) Add() {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&repo)
}
