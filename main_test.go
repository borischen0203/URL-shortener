package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"
	"testing"
	"url-shortener/config"
	"url-shortener/database"
	"url-shortener/dto"
	"url-shortener/errors"
	"url-shortener/logger"
	"url-shortener/router"

	"github.com/stretchr/testify/assert"
)

var GenerateUrl = "/api/url-shortener/v1/url"

func TestMain(m *testing.M) {
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
	fmt.Println(w.Body.String())
}

//Should return 200 with long URL and short URL(existing) when long URL is used

//Should return 200 with long URL and short URL(generate new one wt alias) when alias is unused

//Should return 200 with long URL and short URL(existing) when alias is used by this input long URL

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
	fmt.Println(w.Body.String())
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
	fmt.Println(w.Body.String())
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
	fmt.Println(w.Body.String())
}

//Should return 403 when alias is used by the other long URL

//Should return 404 when short url is not found

//Should return 500 when internal server error
func TestInternalServerError(t *testing.T) {
	out, err := exec.Command("/bin/sh", "../env.sh").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	if err != nil {
		log.Fatal(err)
	}
	logger.Setup()
	config.Setup()
	database.Setup()
	router := router.SetupRouter()
	w := httptest.NewRecorder()

	requestBody := dto.UrlShortenerRequest{
		LongUrl: "https://google.com",
		Alias:   "",
	}

	request, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest(http.MethodPost, GenerateUrl, strings.NewReader(string(request)))
	router.ServeHTTP(w, req)

	actual := w.Body.String()
	expected, _ := json.Marshal(errors.InternalServerError)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), actual)
	fmt.Println(w.Body.String())
}
