package handlers

import (
	"go-projects/dto"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/speps/go-hashids"
)

//test commit4
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
		ShortUrl: "This is short url",
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
		ShortUrl: "http://localhost:8080/" + generateUniqueID(),
	}

	c.JSON(http.StatusOK, urlResponse)
	// c.String(http.StatusOK, "Create URL")
}

func generateUniqueID() string {
	h, _ := hashids.NewWithData(hashids.NewData())
	now := time.Now()
	ID, _ := h.Encode([]int{int(now.UnixNano())})
	return ID
}
