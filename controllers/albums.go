package controllers

import (
"net/http"

"github.com/gin-gonic/gin"
"github.com/agn-7/web-service-gin/models"
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
	models.DB.Find(&albums)

	c.JSON(http.StatusOK, gin.H{"data": albums})
}


func InsertAlbums(c *gin.Context) {
	var input InsertAlbumInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album := models.Album{Title: input.Title, Artist: input.Artist, Price: input.Price}
	models.DB.Create(&album)

	c.JSON(http.StatusOK, gin.H{"data": album})
}


func GetAlbum(c *gin.Context) {
	var album models.Album

	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}

	c.JSON(http.StatusOK, gin.H{"data": album})
  }


func UpdateAlbum(c *gin.Context) {
	var album models.Album

	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	// Validate input
	var input UpdateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	models.DB.Model(&album).Updates(input)
  
	c.JSON(http.StatusOK, gin.H{"data": album})
}


func DeleteAlbum(c *gin.Context) {
	// Get model if exist
	var album models.Album
	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	models.DB.Delete(&album)
  
	c.JSON(http.StatusOK, gin.H{"data": true})
}
