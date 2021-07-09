package database

type Url struct {
	Id       string `json:"id" bson:"id"`
	LongUrl  string `json:"longUrl" bson:"longUrl"`
	ShortUrl string `json:"shortUrl" bson:"shortUrl"`
}
