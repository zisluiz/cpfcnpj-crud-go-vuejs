module zisluiz.com/cpfcnpj-crud-api

go 1.16

replace zisluiz.com/cpfcnpj-crud-api/docs => ../docs

replace zisluiz.com/cpfcnpj-crud-api/http/adapter => ../http/adapter

replace zisluiz.com/cpfcnpj-crud-api/command/application => ../command/application

replace zisluiz.com/cpfcnpj-crud-api/command/domain/exception => ../command/domain/exception

replace zisluiz.com/cpfcnpj-crud-api/command/domain/model => ../command/domain/model

replace zisluiz.com/cpfcnpj-crud-api/command/domain/repository => ../command/domain/repository

replace zisluiz.com/cpfcnpj-crud-api/command/domain/factory => ../command/domain/factory

replace zisluiz.com/cpfcnpj-crud-api/query/application => ../query/application

replace zisluiz.com/cpfcnpj-crud-api/query/model => ../query/model

replace zisluiz.com/cpfcnpj-crud-api/query/dao => ../query/dao

replace zisluiz.com/cpfcnpj-crud-api/infra/config => ../infra/config

replace zisluiz.com/cpfcnpj-crud-api/infra/data => ../infra/data

replace zisluiz.com/cpfcnpj-crud-api/infra/middleware => ../infra/middleware

replace zisluiz.com/cpfcnpj-crud-api/test/mock => ../test/mock

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/cosmtrek/air v1.21.2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/creack/pty v1.1.11 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/labstack/echo/v4 v4.2.1
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/pelletier/go-toml v1.9.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/swaggo/echo-swagger v1.1.0
	github.com/swaggo/swag v1.7.0
	go.mongodb.org/mongo-driver v1.5.1
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210403161142-5e06dd20ab57 // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/tools v0.1.0 // indirect
)
