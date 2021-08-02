# URL-shortener

## Features
- Generate short URL by long URL
- Redirect to long URL by generated URL
- URL generation supports custom alias

## How to run

### Required

- Docker
- MongoDB

### Ready

Download docker and mongodb3.6.23 image

### Conf


```
DB_URL = "mongodb://localhost:27017"
DB_Name = "url_database"
Url_Info_Collection_Name = "url_info"
URL_HOST = "http://localhost:8080/"
PORT = "8080"
```

### Run
```
$ docker run -p 27017:27017 mongo:3.6.23

$ go run main.go
```

Project information and existing API

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (4 handlers)
[GIN-debug] GET    /health                   --> url-shortener/handlers.HealthHandler (4 handlers)
[GIN-debug] GET    /version                  --> url-shortener/handlers.VersionHandler (4 handlers)
[GIN-debug] GET    /:id                      --> url-shortener/handlers.GetLongUrl (4 handlers)
[GIN-debug] POST   /api/url-shortener/v1/url --> url-shortener/handlers.GenerateShortUrl (4 handlers)
[GIN-debug] Listening and serving HTTP on localhost:8080
```

### API example
```
    $ curl -X POST -H "Content-Type: application/json" -d '{"longUrl" : "https://www.google.com/", "alias":"myGoogle"}' "http://localhost:8080/api/url-shortener/v1/url"
```

## Tech Stack
    - RESTful API
    - Swagger
    - logging
    - Gin
    - MongoDB