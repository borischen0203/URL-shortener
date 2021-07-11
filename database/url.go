package database

/**
 * Construct of URL for DB
 *
 * @author: Boris
 * @version: 2021-07-08
 *
 */
type Url struct {
	Id       string `json:"id,omitempty" bson:"id,omitempty" `
	Alias    string `json:"alias,omitempty" bson:"alias,omitempty"`
	LongUrl  string `json:"longUrl" bson:"longUrl"`
	ShortUrl string `json:"shortUrl" bson:"shortUrl"`
}
