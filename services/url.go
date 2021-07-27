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
 * @version: 2021-07-12
 *
 */

import (
	"context"
	"time"
	"url-shortener/config"
	"url-shortener/database"
	"url-shortener/dto"
	e "url-shortener/errors"
	"url-shortener/logger"

	"github.com/speps/go-hashids"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateShortUrlService(request dto.UrlShortenerRequest, response dto.UrlResponse) (int64, dto.UrlResponse, e.ErrorInfo) {
	if !request.Validate() {
		return 400, response, e.BadRequestError
	}

	if isAlias(request) {
		result, err := CreateShortUrlByAlias(request)
		if err != nil {
			logger.Error.Printf("[GenerateShortUrlService] creat short url with alias error: %v\n", err)
			return 500, response, e.InternalServerError // server error
		}
		//Alias is used -> invalid alias
		if (result == database.Url{}) {
			logger.Info.Printf("[GenerateShortUrlService] Alias was used: %v\n", request.Alias)
			return 403, response, e.AliasForbidenError
		}
		return 200, jsonMapping(result, response), e.NoError
	}

	result, err := CreatShortUrl(request)
	if err != nil {
		logger.Error.Printf("[GenerateShortUrlService] creat short url error: %v\n", err)
		return 500, response, e.InternalServerError // server error
	}
	return 200, jsonMapping(result, response), e.NoError
}

//Create short url with alias
func CreateShortUrlByAlias(request dto.UrlShortenerRequest) (database.Url, error) {
	//check the alias exist or not
	result, err := GetUrlById(request.Alias)
	if err != nil {
		logger.Error.Printf("[CreateShortUrlByAlias] search db error: %v\n", err)
		return database.Url{}, err
	}

	// alias is never used -> create successful
	if (result == database.Url{}) {
		result, err := InsertUrlDocument(request.LongUrl, request.Alias)
		if err != nil {
			logger.Error.Printf("[CreateShortUrlByAlias] insert db error: %v\n", err)
			return database.Url{}, err
		}
		return result, nil
	}

	//alias is used -> forbidden alias
	if (result != database.Url{} && result.LongUrl != request.LongUrl) {
		return database.Url{}, nil
	}

	return result, nil
}

//Create short url without alias
func CreatShortUrl(request dto.UrlShortenerRequest) (database.Url, error) {
	//check if long url already exist
	QueryResult, err := GetShortUrlByLongUrl(request.LongUrl)
	if err != nil {
		logger.Error.Printf("[CreatShortUrl] search url error: %v\n", err)
		return database.Url{}, err
	}

	//found long url and return existing info
	if (QueryResult != database.Url{}) {
		logger.Info.Printf("[CreateShortUrl] Found existing long url: %v\n", QueryResult)
		return QueryResult, nil
	}

	//Generate and insert URL info to Mongo
	InsertResult, err := InsertUrlDocument(request.LongUrl)
	if err != nil {
		logger.Error.Printf("[CreateShortUrl] insert db error: %v\n", err)
		return database.Url{}, err
	}
	return InsertResult, nil
}

//check long url exist or not
func GetShortUrlByLongUrl(longUrl string) (database.Url, error) {
	collection := database.UrlCollection
	var result database.Url

	filter := bson.M{"longUrl": longUrl, "id": bson.M{"$exists": true}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Error.Printf("[GetShortUrlByLongUrl] long url is not found: %v\n", longUrl)
			return database.Url{}, nil
		}
		logger.Error.Printf("[GetShortUrlByLongUrl] search db error: %v\n", err)
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
			logger.Info.Printf("[GetUrlById] id or alias is not found: %v\n", id)
			return database.Url{}, nil
		}
		logger.Error.Printf("[GetUrlById] search db error: %v\n", err)
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
	urlPrefix := config.Env.URL_HOST
	url.ShortUrl = urlPrefix + shortId

	if _, err := collection.InsertOne(context.TODO(), url); err != nil {
		logger.Error.Printf("[InsertUrlDocument] insert db error: %v\n", err)
		return database.Url{}, err
	}
	return url, nil
}

func isAlias(request dto.UrlShortenerRequest) bool {
	if len(request.Alias) > 0 {
		return true
	} else {
		return false
	}
}

func jsonMapping(db database.Url, response dto.UrlResponse) dto.UrlResponse {
	response.LongUrl = db.LongUrl
	response.ShortUrl = db.ShortUrl
	return response
}

func generateUniqueID() string {
	h, _ := hashids.NewWithData(hashids.NewData())
	now := time.Now()
	ID, _ := h.Encode([]int{int(now.UnixNano())})
	return ID
}
