package main

import (
	"net/http"

	"github.com/agn-7/web-service-gin/controllers"
	docs "github.com/agn-7/web-service-gin/docs"
	"github.com/agn-7/web-service-gin/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
// @Summary Get all albums
// @Description get all albums
// @Produce  json
// @Success 200 {object} album
// @Router /albums [get]
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
// @Summary Create new album
// @Description Create new album with input payload
// @ID create-album
// @Accept  json
// @Produce  json
// @Param input body album true "Album's info"
// @Success 201 {object} album
// @Router /albums [post]
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// getAlbums responds with the list of all albums as JSON.
// @Summary Get an album by id
// @Description get album by ID
// @ID get-album-by-id
// @Produce  json
// @Param id path int true "Album ID"
// @Success 200 {object} album
// @Router /albums/{id} [get]
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
    router := gin.Default()

    docs.SwaggerInfo.BasePath = "/api/v1"
    v1 := router.Group("/api/v1")
    {
        albums := v1.Group("/albums")
        {
            albums.GET("", getAlbums)
            albums.GET(":id", getAlbumByID)
            albums.POST("", postAlbums)
        }

        db := v1.Group("/db")
        {
            models.ConnectSQLite()
            db_albums := db.Group("/albums")
            {
                db_albums.GET("", controllers.GetAlbums)
                db_albums.GET(":id", controllers.GetAlbum)
                db_albums.POST("", controllers.InsertAlbums)
                db_albums.PUT(":id", controllers.UpdateAlbum)
                db_albums.DELETE(":id", controllers.DeleteAlbum)
            }

            models.ConnectPostgres()
            db_interactions := db.Group("/interactions")
            {
                db_interactions.GET("", controllers.GetInteractions)
            }

        }
    }

    router.GET("/docs", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
    })
    router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    err := router.Run("localhost:8080")
    if err != nil{
        return
    }
}
