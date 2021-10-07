![Gin gopher](https://raw.githubusercontent.com/scraly/gophers/main/mojito-gopher.png)

[![CICD](https://github.com/borischen0203/URL-shortener/actions/workflows/CICD.yml/badge.svg)](https://github.com/borischen0203/URL-shortener/actions/workflows/CICD.yml)

# URL-shortener service
This is a URL shortener service. You can generate a short URL by a long URL.
 Moreover, you can make your alias of short URL.

## Features

- Generate short URL by long URL
- Redirect to long URL by generated URL
- URL generation supports custom alias

## Demo
- No need to run main.go file, you can type the below command in the terminal directly.
- The app may sleep without using. Just wait for a few seconds to wake it up.
```bash
> curl -X POST -H "Content-Type: application/json" -d '{"longUrl" : "https://www.youtube.com/", "alias":""}' "https://short-url-sample.herokuapp.com/api/url-shortener/v1/url"
```

## Project information and existing API
- `GET`    /swagger/*any
- `GET`    /health
- `GET`    /version
- `GET`    /:id
- `POST`   /api/url-shortener/v1/url


## How to run in local

### Required

- Golang
- Gin framework >= 1.7
- Docker
- MongoDB

### Run
```bash
> docker pull mongo:3.6

> docker run -p 27017:27017 mongo:3.6

> go run main.go
```



### Demo in local
```bash
> curl -X POST -H "Content-Type: application/json" -d '{"longUrl" : "https://www.youtube.com/", "alias":""}' "http://localhost:8080/api/url-shortener/v1/url"
```




### Tech Stack
    - RESTful API
    - Swagger
    - Gin
    - Golang
    - MongoDB
    - Docker
    - Github action(CI)
    - Heroku (CD)

### Todo
- [ ] mongoDB add expired time field
- [ ] Add redis for cache
- [ ] MockDB for testing
- [ ] Deploy to AWS
