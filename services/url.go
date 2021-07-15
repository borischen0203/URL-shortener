package services

/**
 * This services mainly handle mongoDB operation
 *  1.CreatShortURL
 * 	2.CreateShortUrlByAlias
 *  3.GetShortUrlByLongUrl
 * 	4.etUrlById
 *  5.InsertUrlDocument
 *
 * @author: Boris
 * @version: 2021-07-12ˇ
 *
 */

import (
	"context"
	"log"
	"time"
	"url-shortener/database"

	"github.com/speps/go-hashids"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateShortUrl(longUrl string) (database.Url, error) {

	//check if long url already exist
	result, err := GetShortUrlByLongUrl(longUrl)
	if err != nil {
		log.Printf("[CreateShortUrl] search url error: %v\n", err)
		return database.Url{}, err
	}
	if (result != database.Url{}) {
		return result, nil
	}

	//Generate and insert URL info
	url, err := InsertUrlDocument(longUrl)
	if err != nil {
		log.Printf("[CreateShortUrl] insert db error: %v\n", err)
		return database.Url{}, err
	}

	return url, nil
}

func CreateShortUrlByAlias(longUrl string, alias string) (database.Url, error) {
	//check the alias exist or not
	result, err := GetUrlById(alias)
	if err != nil {
		log.Printf("[CreateShortUrlByAlias] search db error: %v\n", err)
		return database.Url{}, err
	}

	// alias is never used
	if (result == database.Url{}) {
		result, err := InsertUrlDocument(longUrl, alias)
		if err != nil {
			return database.Url{}, err
		}
		return result, nil
	}

	//alias is used
	if (result != database.Url{} && result.LongUrl != longUrl) {
		return database.Url{}, nil
	}

	return result, nil
}

//check long url exist or not
func GetShortUrlByLongUrl(longUrl string) (database.Url, error) {
	collection := database.UrlCollection
	var result database.Url

	filter := bson.M{"longUrl": longUrl, "id": bson.M{"$exists": true}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("[GetShortUrlByLongUrl] long url is not found: %v\n", longUrl)
			return database.Url{}, nil
		}
		log.Printf("[GetShortUrlByLongUrl] search db error: %v\n", err)
		return database.Url{}, err
	}
	return result, nil
}

//check id or alias exist or not
func GetUrlById(id string) (database.Url, error) {
	collection := database.UrlCollection
	var result database.Url

	filter := bson.D{{"$or", bson.A{bson.M{"id": id}, bson.M{"alias": id}}}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("[GetUrlById] id or alias is not found: %v\n", id)
			return database.Url{}, nil
		}
		log.Printf("[GetUrlById] search db error: %v\n", err)
		return database.Url{}, err
	}
	return result, nil
}

//Inert URL document to DB
func InsertUrlDocument(s ...string) (database.Url, error) {
	collection := database.UrlCollection
	var url database.Url
	var shortId string

	url.LongUrl = s[0]
	if len(s) > 1 {
		url.Alias = s[1]
		shortId = url.Alias
	} else {
		url.Id = generateUniqueID()
		shortId = url.Id
	}

	url.ShortUrl = "http://localhost:8080/" + shortId

	if _, err := collection.InsertOne(context.TODO(), url); err != nil {
		log.Printf("[InsertUrlDoc] insert db error: %v\n", err)
		return database.Url{}, err
	}
	return url, nil
}

func generateUniqueID() string {
	h, _ := hashids.NewWithData(hashids.NewData())
	now := time.Now()
	ID, _ := h.Encode([]int{int(now.UnixNano())})
	return ID
}
