<h1 align="center">Welcome to go-service-taxi 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
  <img alt="documentation: yes" src="https://img.shields.io/badge/Documentation-Yes-green.svg" />
  <img alt="maintained: yes" src="https://img.shields.io/badge/Maintained-Yes-green.svg" />
</p>

>This project using boilerplate designed for READY-TO-GO, which means every basic setup like CORS, Middleware, JWT, User Endpoint, and Redis already setup in this project, and this project support two communication layers gRPC and Restful, and you can just focus on your development, if you need example in this project have example endpoint too named foobar.



### Required Mod

Before run this project please run these commands first

```bash
go get -u gopkg.in/go-playground/validator.v9
go get -u github.com/labstack/echo/v4
go get -u github.com/rs/cors
go get -u github.com/spf13/cobra
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/sethvargo/go-password/password
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/joho/godotenv
go get -u github.com/jinzhu/gorm
go get -u github.com/biezhi/gorm-paginator/pagination
go get -u gopkg.in/go-playground/assert.v1
go get -u github.com/jinzhu/gorm/dialects/mysql
go get -u github.com/jinzhu/gorm/dialects/postgres
go get -u github.com/jinzhu/gorm/dialects/mssql
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/go-redis/redis/v7
go get -u github.com/satori/go.uuid
go get -u github.com/aliyun/aliyun-oss-go-sdk/oss
```



### Compile Proto File

Before your run these commands. You must setting up protoc binary in to you system environment variable

```bash
protoc --proto_path=api/protobuf --proto_path=third_party --go_out=plugins=grpc:api/protobuf foobar.proto
protoc --proto_path=api/protobuf --proto_path=third_party --go_out=plugins=grpc:api/protobuf car.proto
```



### Compile Executable

First compile main.go file into you desired filename for example taxi. Remember every time you change your database driver its means .env file changed too. That's why you need to recompile the main.go file again.

```bash
go build -o taxi main.go
```

If you are using windows you can add .exe extension at the last of filename

```bash
go build -o taxi.exe main.go
```



### Setting Up Database

You can chose which database driver you want to use, this project compatible with postgres, mysql, and mssql depend what you need. But you need to change some syntax in your model depends on selected database driver. Use this command to change database driver, here is the valid dbname.

`dbname [postgres, mysql, mssql, sqlite]`

```
taxi database dbname
go build -o taxi main.go
```



### Starting Project

To start this project you must prepare database and .env file, which the source of .env file is from .env.example you can copy its content to your own .env file, if you are done with configure your .env file you can run this command

Before run web service please choose mode first

`validmode [grpc, rest]`

```bash
taxi switch-mode rest
go build -o taxi main.go
taxi web-start
```



### Or Just Run This for instant

This instant run using mysql database you can change the command depends on your database you used



Run with gRPC

```bash
go build -o taxi main.go
taxi database mysql
taxi switch-mode grpc
go build -o taxi main.go
taxi web-start
```



Run with REST

```
go build -o taxi main.go
taxi database mysql
taxi switch-mode rest
go build -o taxi main.go
taxi web-start
```
