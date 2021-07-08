package handlers

import (
	"log"
	"net/http"
	"time"
	"url-shortener/dto"

	"github.com/gin-gonic/gin"
	"github.com/speps/go-hashids"
)

//test commit5
func UrlHandler(c *gin.Context) {
	// c.JSON(http.StatusOK, []byte("Hello, It Home!"))
	// c.Data(200, "text/plain", []byte("Hello, It Home!"))
	urlRequest := dto.UrlShortenerRequest{
		LongUrl: "This is long url",
	}

	// c.BindHeader(&urlRequest)
	// c.JSON(http.StatusOK, []byte("Hello, It Home!"))
	// c.Data(200, "text/plain", []byte("Hello, It Home!"))
	urlResponse := dto.UrlResponse{
		LongUrl:  urlRequest.LongUrl,
		ShortUrl: "This is short url",
	}

	c.JSON(http.StatusOK, urlResponse)
	// c.String(http.StatusOK, "Hello World")
}

//Generate short URL by long URL
func GenerateShortUrl(c *gin.Context) {
	request := dto.UrlShortenerRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, []byte("Hello, It Home!"))
	// c.Data(200, "text/plain", []byte("Hello, It Home!"))
	response := dto.UrlResponse{
		LongUrl:  request.LongUrl,
		ShortUrl: "http://localhost:8080/" + generateUniqueID(),
	}
	log.Println("long url" + request.LongUrl)

	c.JSON(http.StatusOK, response)
	// c.String(http.StatusOK, "Create URL")
}

func generateUniqueID() string {
	h, _ := hashids.NewWithData(hashids.NewData())
	now := time.Now()
	ID, _ := h.Encode([]int{int(now.UnixNano())})
	return ID
}
