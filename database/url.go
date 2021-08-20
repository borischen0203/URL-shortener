package database

/**
 * Construct of URL for DB
 *
 * @author: Boris
 * @version: 2021-07-08
 *
 */

/**
 * This struct creates a URL info from the four parts: Id, Alias,
 * LongUrl and ShortUrl, which are String, respectively.
 *
 * @param Id               the unique Id of short url as a string
 * @param Alias            the unique Alias of short url as a String
 * @param LongUrl          the original Url as a String
 * @param ShortUrl         the short url as a String
 */
type Url struct {
	Id       string `json:"id,omitempty" bson:"id,omitempty" `
	Alias    string `json:"alias,omitempty" bson:"alias,omitempty"`
	LongUrl  string `json:"longUrl" bson:"longUrl"`
	ShortUrl string `json:"shortUrl" bson:"shortUrl"`
}
