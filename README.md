# Multi tenant application based
Using goose, gin and river
Tenan migrations are provided with background jobs based on riverqueue

## Deps

```
go get github.com/riverqueue/river
go get github.com/riverqueue/river/riverdriver/riverpgxv5
go get github.com/pressly/goose/v3
go get -u github.com/gin-gonic/gin

```

## Setup db

```
docker-compose up -d
```
