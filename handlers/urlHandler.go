package handlers

import (
	"net/http"

	"url-shortener/dto"
	"url-shortener/errors"
	"url-shortener/services"

	"github.com/gin-gonic/gin"
)

//Generate short URL by long URL
func GenerateShortUrl(c *gin.Context) {
	request := dto.UrlShortenerRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := services.CreateShortUrl(request.LongUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
		return
	}

	response := dto.UrlResponse{
		LongUrl:  res.LongUrl,
		ShortUrl: res.ShortUrl,
	}

	c.JSON(http.StatusOK, response)
}

func GetLongUrl(c *gin.Context) {
	id := c.Param("id")
	result, err := services.GetLongUrlById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, errors.UrlNotFoundError)
		return
	}
	location := result.LongUrl
	c.Redirect(http.StatusFound, location)

}
