package dto

import (
	"regexp"
	e "url-shortener/errors"
)

/**
 * Construct of URL for API
 *
 * @author: Boris
 * @version: 2021-07-08
 *
 */
type UrlShortenerRequest struct {
	LongUrl string `json:"longUrl" binding:"required"`
	Alias   string `json:"alias, omitempty"`
}

type UrlResponse struct {
	LongUrl  string `json:"longUrl"`
	ShortUrl string `json:"shortUrl"`
}

type RedirectRequest struct {
	RedirectURL string `json:"redirect_url"`
}

func (r UrlShortenerRequest) Validate() (bool, e.ErrorInfo) {
	if r.IsEmpty() {
		return false, e.InvalidLongUrlError
	}
	if !r.IsValidAlias() {
		return false, e.InvalidAliasError
	}
	return true, e.NoError
}

func (r UrlShortenerRequest) IsValidAlias() bool {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{0,30}$", r.Alias); !ok {
		return false
	}
	return true
}

func (r UrlShortenerRequest) IsEmpty() bool {
	return r.LongUrl == ""
}
