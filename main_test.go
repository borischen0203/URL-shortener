package main

import (
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"os"

	"strings"
	"testing"
	"url-shortener/config"
	"url-shortener/database"
	"url-shortener/dto"
	"url-shortener/errors"
	"url-shortener/logger"
	"url-shortener/router"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

var GenerateUrl = "/api/url-shortener/v1/url"

func TestMain(m *testing.M) {
	logger.Setup()
	r := m.Run()

	if r == 0 && testing.CoverMode() != "" {
		c := testing.Coverage() * 100
		l := 70.00
		fmt.Println("=================================================")
		fmt.Println("||               Coverage Report               ||")
		fmt.Println("=================================================")
		fmt.Printf("Cover mode: %s\n", testing.CoverMode())
		fmt.Printf("Coverage  : %.2f %% (Threshold: %.2f %%)\n\n", c, l)
		if c < l {
			fmt.Println("[Tests passed but coverage failed]")
			r = -1
		}
	}

	os.Exit(r)
}

func TestHealth(t *testing.T) {
	router := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestVersion(t *testing.T) {
	router := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/version", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

//Should return 200 with long URL and short URL(generate new one) when long URL is unused

//Should return 200 with long URL and short URL(existing) when long URL is used
func TestGenerateUrl(t *testing.T) {
	logger.Setup()
	config.Setup()
	database.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.UrlShortenerRequest{
		LongUrl: "https://www.youtube.com/",
	}
	// responseBody := dto.UrlResponse{
	// 	LongUrl:  "https://www.youtube.com/",
	// 	ShortUrl: "http://localhost:8080/gJXz8NqV7N40l",
	// }

	request, _ := json.Marshal(requestBody)
	// expected, _ := json.Marshal(responseBody)

	req, _ := http.NewRequest(http.MethodPost, GenerateUrl, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// assert.Equal(t, string(expected), w.Body.String())
}

//Should return 200 with long URL and short URL(generate new one wt alias) when alias is unused

//Should return 200 with long URL and short URL(existing) when alias is used by this input long URL
func TestValidAlias(t *testing.T) {
	logger.Setup()
	config.Setup()
	database.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.UrlShortenerRequest{
		LongUrl: "https://www.youtube.com/",
		Alias:   "myYoutube",
	}
	// responseBody := dto.UrlResponse{
	// 	LongUrl:  "https://www.youtube.com/",
	// 	ShortUrl: "http://localhost:8080/myYoutube",
	// }

	request, _ := json.Marshal(requestBody)
	// expected, _ := json.Marshal(responseBody)

	req, _ := http.NewRequest(http.MethodPost, GenerateUrl, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// assert.Equal(t, string(expected), w.Body.String())
}

//Should return 302 when short URL is valid
func TestShortUrlIsValid(t *testing.T) {
	logger.Setup()
	config.Setup()
	database.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/myYoutube", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
}

//Should return 400 when longUrl field is empty
func TestEmptyLongUrlFields(t *testing.T) {

	logger.Setup()
	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.UrlShortenerRequest{
		LongUrl: "",
	}
	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidLongUrlError)

	req, _ := http.NewRequest(http.MethodPost, GenerateUrl, strings.NewReader(string(request)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return 400 when longUrl field is missing
func TestMissingLongUrlFields(t *testing.T) {

	logger.Setup()
	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.UrlShortenerRequest{}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.InvalidLongUrlError)

	req, _ := http.NewRequest(http.MethodPost, GenerateUrl, strings.NewReader(string(request)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return 400 when longUrl is invalid

//Should return 400 when alias is invalid
func TestAliasIsInvalid(t *testing.T) {
	logger.Setup()
	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.UrlShortenerRequest{
		LongUrl: "https://google.com",
		Alias:   "*InvalidAlias",
	}

	request, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest(http.MethodPost, GenerateUrl, strings.NewReader(string(request)))
	router.ServeHTTP(w, req)

	actual := w.Body.String()
	expected, _ := json.Marshal(errors.InvalidAliasError)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), actual)
}

//Should return 403 when alias is used by the other long URL
func TestForbiddenAlias(t *testing.T) {
	logger.Setup()
	config.Setup()
	database.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.UrlShortenerRequest{
		LongUrl: "https://www.gogole.com/",
		Alias:   "myYoutube",
	}

	request, _ := json.Marshal(requestBody)
	expected, _ := json.Marshal(errors.AliasForbidenError)

	req, _ := http.NewRequest(http.MethodPost, GenerateUrl, strings.NewReader(string(request)))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

//Should return 404 when short url is not found
func TestShortUrlNotFound(t *testing.T) {
	logger.Setup()
	config.Setup()
	database.Setup()

	router := router.SetupRouter()
	w := httptest.NewRecorder()

	expected, _ := json.Marshal(errors.UrlNotFoundError)

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/notFound", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
