package repository

import (
	"github.com/sachinadate/learnginframeworkp1/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video entity.Video)
	FindAll() []entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	sqlDB, err := db.connection.DB()
	if err != nil {
		panic("Failed to close database")
	}
	sqlDB.Close()
}

func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
