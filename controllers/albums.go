package controllers

import (
	"net/http"

	"github.com/agn-7/web-service-gin/models"
	"github.com/gin-gonic/gin"
)


type InsertAlbumInput struct {
	Title  string `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price"`
}


type UpdateAlbumInput struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}


func GetAlbums(c *gin.Context) {
	var albums []models.Album
	models.SQLiteDB.Find(&albums)

	c.JSON(http.StatusOK, gin.H{"data": albums})
}


func InsertAlbums(c *gin.Context) {
	var input InsertAlbumInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album := models.Album{Title: input.Title, Artist: input.Artist, Price: input.Price}
	models.SQLiteDB.Create(&album)

	c.JSON(http.StatusOK, gin.H{"data": album})
}


func GetAlbum(c *gin.Context) {
	var album models.Album

	if err := models.SQLiteDB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}

	c.JSON(http.StatusOK, gin.H{"data": album})
  }


func UpdateAlbum(c *gin.Context) {
	var album models.Album

	if err := models.SQLiteDB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}

	// Validate input
	var input UpdateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}

	models.SQLiteDB.Model(&album).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": album})
}


func DeleteAlbum(c *gin.Context) {
	// Get model if exist
	var album models.Album
	if err := models.SQLiteDB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}

	models.SQLiteDB.Delete(&album)

	c.JSON(http.StatusOK, gin.H{"data": true})
}


// GetInteractions gets all interactions from the database
// @Summary Get all interactions
// @Description Retrieves a list of all interactions
// @ID get-all-interactions
// @Produce json
// @Success 200 {array} models.Interaction
// @Failure 500 {object} map[string]interface{}
// @Router /db/interactions [get]
func GetInteractions(c *gin.Context) {
    var interactions []models.Interaction
    result := models.PostgresDB.Preload("Messages").Find(&interactions)

    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": interactions})
}
