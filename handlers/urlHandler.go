package handlers

/***
 *
 * This file mainly handle URL
 *
 * @author: Boris
 * @version: 2021-07-08
 *
 */

import (
	"log"
	"net/http"

	"url-shortener/database"
	"url-shortener/dto"
	"url-shortener/errors"
	"url-shortener/logger"
	"url-shortener/services"

	"github.com/gin-gonic/gin"
)

// Generate short URL by long URL
// @Summary create shortening url
// @Description create shortening url
// @Tags Shorten Url
// @Accept json
// @Produce json
// @Param body body dto.UrlShortenerRequest true "body"
// @Success 200 {string} string "ok"
// @Router /api/url-shortener/v1/url [post]
func GenerateShortUrl(c *gin.Context) {
	request := dto.UrlShortenerRequest{}
	response := dto.UrlResponse{}
	c.BindJSON(&request)
	logger.Info.Printf("[GenerateShortUrlHandler] request=%+v\n", request)
	statusCode, result, err := services.GenerateShortUrlService(request, response)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 400:
		c.JSON(http.StatusBadRequest, err)
	case 403:
		c.JSON(http.StatusForbidden, err)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}
}

func GetLongUrl(c *gin.Context) {
	id := c.Param("id")
	result, err := services.GetUrlById(id)
	if err != nil {
		log.Printf("[GetLongUrl] search db error: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
		return
	}
	if (result == database.Url{}) {
		log.Printf("[GetLongUrl] Url is not found: %v\n", err)
		c.JSON(http.StatusNotFound, errors.UrlNotFoundError)
		return
	}

	location := result.LongUrl
	c.Redirect(http.StatusFound, location)

}
