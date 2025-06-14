package handler

import (
	"net/http"

	"github.com/amirkh8006/url-shortener/shortener"
	"github.com/amirkh8006/url-shortener/store"
	"github.com/gin-gonic/gin"
)

type urlCreationRequestBody struct {
	Url    string `json:"url" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest urlCreationRequestBody
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortUrl := shortener.GenerateShortUrl(creationRequest.Url, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.Url, creationRequest.UserId)

	c.JSON(http.StatusOK, gin.H{
		"messaage": "ShortUrl Created Successfuly",
		"shortUrl": "http://localhost:8080/" + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	initialLink := store.RetrieveInitialUrl(shortUrl)

	c.Redirect(http.StatusPermanentRedirect, initialLink)
}
