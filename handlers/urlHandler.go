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
	// "fmt"

	"net/http"

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
// @Success 200 {object} dto.UrlResponse "ok"
// @Failure 400 {object} errors.ErrorInfo "bad request"
// @Failure 403 {object} errors.ErrorInfo "Forbiden"
// @Failure 500 {object} errors.ErrorInfo "Internal server error"
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

// @Summary Redirect original URL by short URL
// @Description Redirect original URL by short URL
// @Tags Shorten Url
// @Accept json
// @Produce json
// @Param body body dto.UrlShortenerRequest true "body"
// @Success 200 {object} dto.UrlResponse "Redirect"
// @Failure 404 {object} errors.ErrorInfo "Not found"
// @Failure 500 {object} errors.ErrorInfo "Internal server error"
// @Router /:id [get]
func GetLongUrl(c *gin.Context) {
	request := dto.RedirectRequest{Id: c.Param("id")}
	// response := dto.UrlResponse{}
	logger.Info.Printf("[GetLongUrlHandler] request=%+v\n", request)
	statusCode, result, err := services.GetOriginalUrlService(request, response)

	switch statusCode {
	case 200:
		c.Redirect(http.StatusFound, result.LongUrl)
	case 404:
		c.JSON(http.StatusNotFound, err)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}

}
