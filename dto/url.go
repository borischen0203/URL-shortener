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

/**
 * This struct creates a URL shortener request from the two parts:
 * LongUrl and ShortUrl, which are String, respectively.
 *
 * @param LongUrl          the original Url as a String
 * @param ShortUrl         the short Url as a String
 */
type UrlShortenerRequest struct {
	LongUrl string `json:"longUrl" binding:"required"`
	Alias   string `json:"alias, omitempty"`
}

/**
 * This struct creates a URL response from the two parts:
 * LongUrl and ShortUrl, which are String, respectively.
 *
 * @param LongUrl          the original Url as a String
 * @param ShortUrl         the short Url as a String
 */
type UrlResponse struct {
	LongUrl  string `json:"longUrl"`
	ShortUrl string `json:"shortUrl"`
}

/**
 * This struct creates a URL Redirect Request from the one part:
 * Id, which is String, respectively.
 *
 * @param Id               the unique Id of short url as a string
 */
type RedirectRequest struct {
	Id string `param:"id"`
}

//Validate function validate the request
func (r UrlShortenerRequest) Validate() (bool, e.ErrorInfo) {
	if r.IsEmpty() {
		return false, e.InvalidLongUrlError
	}
	if !r.IsValidAlias() {
		return false, e.InvalidAliasError
	}
	return true, e.NoError
}

//IsValidAlias function validate Alias is valid alphabet or number, and less than 30 lengths
func (r UrlShortenerRequest) IsValidAlias() bool {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{0,30}$", r.Alias); !ok {
		return false
	}
	return true
}

//IsEmpty function validate long url is not empty input
func (r UrlShortenerRequest) IsEmpty() bool {
	return r.LongUrl == ""
}
