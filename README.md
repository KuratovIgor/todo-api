# Todo API

### Simple REST API for working with to-do lists.
### Project implementation is built on REST and based on clean architecture.
## Project structue:
 - *cmd* - contains main.go which is the entry point
 - *configs* - configs for starting server and database startup
 - *docs* - generated swagger doc
 - *pkg*
   - *handlers* - contains http requests handlers
   - *service* - contains all the business logic of the project
   - *repository* - contains the logic of working with the database
## Running project
### You can run the project both locally and in docker.
### Run for development:
 - Create *.env* in root directory of project and fill it in according to *.env.example*
 - Pull postgres docker image and run container. Make sure you have ***DOCKER*** installed before running the command
   ```dockerfile
   docker pull postgres
   docker run -d -p 5432:5432 --name todo-db -e POSTGRES_PASSWORD='<pass>' --rm postgres
   ```
 - Create database structure by schema. Make sure you have ***MIGRATE*** installed before running the command
   ```dockerfile
   migrate -path ./schema --database 'postgres://postgres:<pass>@localhost:5432/postgres?sslmode=disable' up
   ```
 - Run project
   ```go
   go run cmd/main.go
   ```
 ### Run by Docker
 - Create .env in root directory of project and fill it in according to .env.example
 - Run command for running app in docker from *Makefile*
   ```
   make run
   ```
 ## Swagger
### Run `make swag` command for init swagger. Swagger can be opened by URL ***swagger/index.html***.