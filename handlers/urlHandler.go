package handlers

/***
 *
 *
 * Version: 2021/07/15
 *
 */
import (
	"log"
	"net/http"

	"url-shortener/database"
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

	var response dto.UrlResponse
	//Alias requirement
	if len(request.Alias) > 0 {
		res, err := services.CreateShortUrlByAlias(request.LongUrl, request.Alias)
		if err != nil {
			log.Printf("[GenerateShortUrl] create url with alias error: %v\n", err)
			c.JSON(http.StatusInternalServerError, errors.InternalServerError)
			return
		}
		//Alias is used
		if (res == database.Url{}) {
			c.JSON(http.StatusForbidden, errors.AliasForbidenError)
			return
		}

		response.LongUrl = res.LongUrl
		response.ShortUrl = res.ShortUrl
		c.JSON(http.StatusOK, response)
		return
	}

	res, err := services.CreateShortUrl(request.LongUrl)
	if err != nil {
		log.Printf("[GenerateShortUrl] create url error: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
		return
	}
	response.LongUrl = res.LongUrl
	response.ShortUrl = res.ShortUrl

	c.JSON(http.StatusOK, response)
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
