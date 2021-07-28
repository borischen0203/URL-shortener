package test

import (
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"

	"encoding/json"
	"url-shortener/config"
	"url-shortener/database"
	"url-shortener/dto"
	"url-shortener/logger"
	"url-shortener/router"

	"github.com/stretchr/testify/assert"
	// "github.com/tedmax100/gin-angular/router"
)

var GenerateUrl = "/api/url-shortener/v1/url"

// func TestMain(m *testing.M) {
// 	r := m.Run()

// 	if r == 0 && testing.CoverMode() != "" {
// 		c := testing.Coverage() * 100
// 		l := 70.00
// 		fmt.Println("=================================================")
// 		fmt.Println("||               Coverage Report               ||")
// 		fmt.Println("=================================================")
// 		fmt.Printf("Cover mode: %s\n", testing.CoverMode())
// 		fmt.Printf("Coverage  : %.2f %% (Threshold: %.2f %%)\n\n", c, l)
// 		if c < l {
// 			fmt.Println("[Tests passed but coverage failed]")
// 			r = -1
// 		}
// 	}

// 	os.Exit(r)
// }

// func TestHealth(t *testing.T) {
// 	router := SetupRouter()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/health", nil)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

// func TestVersion(t *testing.T) {
// 	router := setupRouter()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/version", nil)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

func TestGenerateUrl(t *testing.T) {
	logger.Setup()
	config.Setup()
	database.Setup()
	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.UrlShortenerRequest{
		LongUrl: "https://www.youtube.com/",
		Alias:   "",
	}
	responseBody := dto.UrlResponse{
		LongUrl:  "https://www.youtube.com/",
		ShortUrl: "http://localhost:8080/gJXz8NqV7N40l",
	}

	request, _ := json.Marshal(requestBody)
	expectedBody, _ := json.Marshal(responseBody)

	req, _ := http.NewRequest(http.MethodPost, GenerateUrl, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expectedBody), w.Body.String())
}
