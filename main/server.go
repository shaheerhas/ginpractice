package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shaheerhas/ginpractice/service"
)

func main() {

	//dsn := "host=localhost user=postgres password=tiger123 dbname=dvdrental port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	var svc service.Service
	svc.Db = service.SetupModels()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", svc.Db)
		c.Next()
	})
	var c gin.Context
	fmt.Println(c.Get("db"))
	//service.Db.AutoMigrate(&entity.Album{})
	//db.Create(&service.Albums)

	r.GET("/albums", svc.GetAlbums)
	r.GET("/albums/:id", svc.GetAlbumById)
	r.DELETE("/albums/:id", svc.DeleteAlbumById)
	r.POST("/albums", svc.PostAlbums)

	r.Run("localhost:8081")

}
