package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaheerhas/ginpractice/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Albums = []entity.Album{
	// {ID: 1, Title: "Blue Train", ArtistName: "John Coltrane", Price: 56.99},
	// {ID: 2, Title: "Midnight", ArtistName: "One Direction", Price: 46.99},
	// {ID: 3, Title: "Red", ArtistName: "Taylor Swift", Price: 66.99},
	// {ID: 4, Title: "21", ArtistName: "Gracie Abrams", Price: 26.99},
}

func SetupModels() *gorm.DB {

	dsn := "host=localhost user=postgres password=tiger123 dbname=dvdrental port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("couldn't connect to the database")
	}
	db.AutoMigrate(&entity.Album{})

	return db
}

type Service struct {
	Db *gorm.DB
}

func (svc Service) GetAlbums(c *gin.Context) {

	var albums []entity.Album
	svc.Db.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func (svc Service) PostAlbums(c *gin.Context) {

	var newAlbum entity.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := svc.Db.Create(&newAlbum)
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, result.Error.Error()+" couldn't create record in DB")
		return
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func (svc Service) GetAlbumById(c *gin.Context) {

	id := c.Param("id")
	var album entity.Album
	if err := svc.Db.Where("id = ?", id).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, "record with this id wasn't found")
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

//remove element from an album slice
func remove(s []entity.Album, i int) []entity.Album {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (svc Service) DeleteAlbumById(c *gin.Context) {
	id := c.Param("id")
	var album entity.Album
	if err := svc.Db.Where("id = ?", id).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, "record with this id wasn't found")
		return
	}
	result := svc.Db.Delete(&album)
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, result.Error.Error()+" couldn't delete record in DB")
		return
	}
	c.IndentedJSON(http.StatusCreated, album)
}
