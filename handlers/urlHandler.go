package handlers

import (
	"go-projects/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UrlHandler(c *gin.Context) {
	// c.JSON(http.StatusOK, []byte("Hello, It Home!"))
	// c.Data(200, "text/plain", []byte("Hello, It Home!"))
	urlRequest := dto.UrlRequest{
		LongUrl: "This is long url",
	}

	// c.BindHeader(&urlRequest)
	// c.JSON(http.StatusOK, []byte("Hello, It Home!"))
	// c.Data(200, "text/plain", []byte("Hello, It Home!"))
	urlResponse := dto.UrlResponse{
		LongUrl:  urlRequest.LongUrl,
		ShortUrl: "This is shor url",
	}

	c.JSON(http.StatusOK, urlResponse)
	// c.String(http.StatusOK, "Hello World")
}

func CreateShortUrl(c *gin.Context) {
	urlRequest := dto.UrlRequest{}
	c.BindHeader(&urlRequest)
	// c.JSON(http.StatusOK, []byte("Hello, It Home!"))
	// c.Data(200, "text/plain", []byte("Hello, It Home!"))
	urlResponse := dto.UrlResponse{
		LongUrl:  urlRequest.LongUrl,
		ShortUrl: "This is shor url",
	}

	c.JSON(http.StatusOK, urlResponse)
	// c.String(http.StatusOK, "Create URL")
}
