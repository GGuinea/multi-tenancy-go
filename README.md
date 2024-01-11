# Multi tenant application based
CRUD applications with multitenant database architecture. Taking care of migrations, schema creation without any ORMs.
Using goose, gin, PostgreSQL

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
