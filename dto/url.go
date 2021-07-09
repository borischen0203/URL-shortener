package dto

type UrlShortenerRequest struct {
	// LongUrl string `json:"long_url" binding:"required, NotEmptyValidator"`
	LongUrl string `json:"longUrl" binding:"required"`
}

type UrlResponse struct {
	LongUrl  string `json:"longUrl"`
	ShortUrl string `json:"shortUrl"`
}

type RedirectRequest struct {
	RedirectURL string `json:"redirect_url"`
}
