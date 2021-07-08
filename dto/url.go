package dto

type UrlShortenerRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	// LongUrl string `json:"long_url"`
}

type UrlResponse struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

type RedirectRequest struct {
	RedirectURL string `json:"redirect_url"`
}
