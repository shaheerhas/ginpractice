package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Midnight", Artist: "One Direction", Price: 46.99},
	{ID: "3", Title: "Red", Artist: "Taylor Swift", Price: 66.99},
	{ID: "4", Title: "21", Artist: "Gracie Abrams", Price: 26.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album with this id not found"})
}

//remove element from an album slice
func remove(s []album, i int) []album {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	for i := range albums {
		if albums[i].ID == id {
			albums = remove(albums, i)
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album with this id not found"})
}

func main() {

	r := gin.Default()
	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbumById)
	r.DELETE("/albums/:id", deleteAlbumById)
	r.POST("/albums", postAlbums)

	r.Run("localhost:8080")

}
