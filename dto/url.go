package dto

/**
 * Construct of URL for API
 *
 * @author: Boris
 * @version: 2021-07-08
 *
 */
type UrlShortenerRequest struct {
	// LongUrl string `json:"long_url" binding:"required, NotEmptyValidator"`
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
