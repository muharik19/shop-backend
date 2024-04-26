> Proposed gin framework based service .This project is still in development stage, any critic and suggestion all very based including but not limited to project name, function naming, folder structure etc. please refer to Readme.md.

## Prerequisite

- Install [go](https://golang.org/doc/install) in local machine for convenience development experience (auto complete, code sugestion, etc)
- Install golang plugin to your editor choice (ie. VSCode, Atom, Goland, Intellij IDE)
- [Docker](https://docs.docker.com/install/) and [docker-compose](https://docs.docker.com/compose/)

## Install gin framework

- Create project with go.mod
- Create project folder structure

  ### Folder structure.

        ├── constant                       # Define for define const
        ├── controller                     # Define for define handler endpoint
        ├── docs                           # Define documentation rest
        ├── middleware                     # Define middleware jwt, auth, cors
        ├── migration                      # Define for define sql
        ├── models                         # Define struct
        ├── pkg                            # Define function using general
        ├── repositories                   # Define data layer
        ├── routes                         # Define router
        ├── services                       # Define function that using for call surrounding system

- Install gin (https://github.com/gin-gonic/gin)
  go get -u github.com/gin-gonic/gin

## How to run

Despite, it is possible to run this project in local machine Please follow this steps:

- swag init for generate document swagger
- Run apps gp to root project `go run main.go`.

## Documentation API Swagger

[Swagger](http://localhost:9090/swagger/index.html)

## Documentation API Postman

[API](https://documenter.getpostman.com/view/4324137/2sA3Bt3Vee)
