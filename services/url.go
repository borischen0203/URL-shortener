package services

import (
	"context"
	"log"
	"time"
	"url-shortener/database"

	"github.com/speps/go-hashids"
	"go.mongodb.org/mongo-driver/bson"
)

func GetShortUrlByLongUrl(longUrl string) (database.Url, error) {

	collection := database.UrlCollection
	var result database.Url
	//if error !=nil means no found
	err := collection.FindOne(context.Background(),
		bson.M{"longUrl": longUrl}).Decode(&result)

	return result, err
}

func GetLongUrlById(id string) (database.Url, error) {
	collection := database.UrlCollection
	var result database.Url
	//if error !=nil means no found
	err := collection.FindOne(context.Background(),
		bson.M{"id": id}).Decode(&result)

	return result, err
}

func CreateShortUrl(longUrl string) (database.Url, error) {
	collection := database.UrlCollection

	//find exist url
	if findOne, err := GetShortUrlByLongUrl(longUrl); err == nil {
		return findOne, nil
	}

	//create url
	var url database.Url
	url.Id = generateUniqueID()
	url.LongUrl = longUrl
	url.ShortUrl = "http://localhost:8080/" + url.Id

	if _, err := collection.InsertOne(context.TODO(), url); err != nil {
		log.Println(err)
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
