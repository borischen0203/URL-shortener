package dto

type UrlRequest struct {
	LongUrl string `json:"long_url"`
}

type UrlResponse struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}
