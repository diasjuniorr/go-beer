# Go Beer API

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Contributing](../CONTRIBUTING.md)

## About <a name = "about"></a>

Simple API showing how to create a CRUD API with Go and sqlite.

## Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Installing

Clone the repo

```
https://github.com/diasjuniorr/go-beer.git
```

And run the project

```
cd go-beer
go build api/main.go
./main
```

## Usage <a name = "usage"></a>

The API provides the following endpoints:

- `POST /v1/beers` create a new beer
- `GET /v1/beers` list all beers
- `GET /v1/beers/{beerId}` get beer by id
- `PUT /v1/beers/{beerid}` update a beer by id
- `DELETE /v1/beers/{beerid}` delete a beer by id
