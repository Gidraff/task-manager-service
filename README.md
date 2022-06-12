[![CircleCI](https://circleci.com/gh/Gidraff/task-manager-service/tree/develop.svg?style=svg)](https://circleci.com/gh/Gidraff/task-manager-service/?branch=develop)
#### TaskMan RESTful API built in Go
The service allows user to divided work into small manageable tasks, providing CRUD operations over RESTful API.

#### Technologies used:

* GORM
* Gorilla/Mux
* PostgreSQL
* go mod
* logrus

#### Prerequisite
To setup, you will first have to grab a copy of this project by cloning it into your local machine. After you've cloned the repo make sure you have the following installed: 

- Go version: `go1.13.7`
- Docker engine
- Make
- Postgresql

#### Database connection
To create a db connection for testing purposes, make sure you've installed Postgres or you have an instance running on a server on the cloud. Create a secure DB (provide credentials). Provide start configurations either through the yaml file or env variables 
`
#### Build and Test

Before starting the application, make sure everything is work by running:

```make build```. This will generate the application binary which can be executed on the console. `.go` and `bin` folders. Bin folder contains the application binary which can be executed on the console.

`make test` to check if the tests are passing.

#### Start the server

To start the server from main.go, run `go run cmd/api/main.go` on your terminal from the root folder.

Alternatively, you can start the server by executing the apps binary generated from the build step by running `./<pathname>/api`.
The API should be ready to accept request. 

