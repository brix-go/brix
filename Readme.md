# Brix - A tool for generating golang project

Brix is a tool used to generate Golang applications. Its name is derived from the toy "Lego Bricks," which are little components of lego that can be disassembled and reassembled to construct something desired when assembled into a single unit. in addition to brix, which may arrange components (modules) to form a Go application. Brix lets us to quickly create good go applications.

## Brix Generator
![Brix](https://res.cloudinary.com/dkkisw3mc/image/upload/v1702405618/brix_yl2krk.png)

## Feature
- Clean Architecture
- Postgres Connection
- Error Handling
- Middleware JWT
- Redis Scaffold

## Project Structure
```
.
├── cmd
│   ├──  application
│   │   ├──  app.go
├── config
│   ├── config.go
├── external
├── infrastructure
│   ├──  database
│   │   ├──  database.go
│   ├──  redis
│   │   ├──  redis.go
├── internal
│   ├──  domain
│   │   ├──  user
│   │   │   ├──    controller
│   │   │   ├──    service
│   │   │   ├──    repository
│   │   │   ├──    user.go
├── middleware
├── router
├── script
├── shared
├── config.yml
├── main.go
├── go.mod
└── go.sum
```
The following is a description of the project structure:
* `cmd` : this module is the initial stage of the application will run, in this folder there is an application folder which contains the `app.go` file. This file is the place to define repo, service, controller, router, database, router, etc.

* `config` : This module contains the `config.go` file which serves to provide config variables for various environments, production or development.

* `external` : this module is the place to use third-party libraries or APIs.

* `infrastructure` : this module is to define the infrastructure that will be used in the application (database, redis, etc).

* `internal` : This module is the core module of the application and contains the implementation of various business logic.

* `domain` : sub-modules of internal modules that contain more domain-specific business logic.

* `middleware` : This module contains middleware used in the application (auth, error handling, etc.).

* `router` : this module contains a router file where endpoints are defined as well as connecting the middleware and controller..

* `script` : this module contains a json file that is the errorContract of the application.

* `shared` : this module serves to store variables or functions that can be reused in various places.

* `config.yaml` : This file is a configuration file where config values for the development or production environment are set..

* `main.go` : this file will be run the first time.


## Requirements
To use Brix, you need to have the following software installed on your system:
* Go 1.19 or higher
* Git
* PostgreSQL 15 or higher (optional)
* Redis (optional)

## How to install Brix
You can install Brix with the following command:
```bash
go install github.com/brix-go/brix/v2
```
> Tips: If `go install` succeeds but the `brix` command is not recognized, it is because the environment variable is not configured. You can add the GOBIN directory to the environment variable.

## How to use Brix
### 1. Create Project
for creating a project you can use this following command :
```bash
brix new projectName
```
This command will create a folder named `projectName` and generate a clean architecture golang application project structure.

### 2. Create Component
You can create controller, services, repositories for your project using the following commands:
```bash
brix gen all domainName
```
ex : `brix gen all order`, this command will create `domain order`, `controller`, `service`, `repository` automatically.

or you can use this command to create one by one
```bash
brix gen handler user
brix gen service user
brix gen repository user
brix gen model user
```


### 3. Running Project
For starting the project, besides using the `go run main.go` command, you can also use the command below:
```bash
brix run
```

## Contribution
If you find any issues or have any improvement suggestions, please feel free to raise an issue or submit a pull request. Your contributions are highly appreciated!

## License
Brix is released under the MIT License. For more information, see the [LICENSE](https://github.com/brix-go/brix/blob/main/LICENSE) file.
