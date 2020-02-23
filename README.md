# Go-RECO

This repo for example implmentation golang and framework [Revel](http://revel.github.io/), and this repo using database [cockroachDB](cockroachlabs.com), this is simple example how to using Revel and CockroachDB Rest API

## Collection Postman

you can import collection in
> `utils/postman/go-postgre.postman_collection.json`

## Documentation API 

You can read documentation for this repo using this [link](https://github.com/cakazies/go-reco/wiki)

## Run Migration with

> go run application/migration/migrate.go

## How to run

- you can configuration config (**Database**) in folder `conf/config.toml` and compare your local configs, database, host and depends

you can install dependencies using `go mod` or one by one,

- install dependencies go get
  - `go get github.com/jinzhu/gorm`
  - `go get github.com/lib/pq`
  - `go get github.com/revel/revel`
  - `go get github.com/spf13/viper`

## Run Local
you can running this repo like revel running:

`revel run -a go-reco`
