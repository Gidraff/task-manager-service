#### TaskMan
The service allows user to divided work into small manageable tasks, providing CRUD operations over RESTful API.

##### Technologies used:

* db/sql
* Echo
* Postgres Database
* go mod
* logrus

##### Prerequisite
To setup, you wll first have to grab a copy of this project by cloning it into your local machine. After you've cloned the repo make sure you have the following installed: 

- Go version: `go1.13.7`
- Docker engine
- Make
- Postgresql

##### Create DB locally
To create a db connection for testing purposes, make sure you've installed Postgres locally or you have an instance running on a server on the cloud. Create a secure DB (provide credentials). Now create tables on the DB you've just created by running
your sql in the root directory script as follow: `psql postgres -h 127.0.0.1 -d <example_db_name> -f schema.sql
`

##### Build and Test

Before starting the application, make sure everything is work by running: 

**NOTE:** *Make sure you are in the root folder.* 

To compile the app Run ```make build```. This will generate `.go` and `bin` folders. Bin folder contains the application binary which can be executed on the console.

If the build is successful, run `make test` to check if the tests are passing. Otherwise reach out to me on twitter.

##### Start the server

To start the server from main.go, run `go run cmd/taskman/main` on your terminal from the root folder.

Alternatively, you can start the server by executing the apps binary generated from the build step by running `./bin/${OS_ARCH}/taskman`.
The API should be ready to accept request. Send a `GET` request to `http://localhost/api/v1/` using your preferred client, you should see `It works!`. This means the API is ready for requests.

