# URL-shortener

## How to run

### Required

- Docker
- MongoDB

### Ready

Download docker and mongodb3.6.23

### Conf


```
DB_URL = "mongodb://localhost:27016"
DB_Name = "url_database"
Url_Info_Collection_Name = "url_info"
URL_HOST = "http://localhost:8080/"
PORT = "8080"
```

### Run
```
docker run -p 27016:27017 mongo:3.6.23-xenial

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
[GIN-debug] Listening and serving HTTP on 127.0.0.1:8080
```

## Features

- RESTful API
- Swagger
- logging
- Gin