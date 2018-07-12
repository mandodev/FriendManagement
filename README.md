# Friend-Management

Friend Management API written on Go using Gin-Gonic Framework (https://github.com/gin-gonic/gin) and 
[GORM](https://github.com/jinzhu/gorm) as ORM Framework

## Run With Docker
run following command on your terminal, make sure you have installed docker and docker-compose (https://docs.docker.com/compose/install/)
`docker-compose up`

## API End Point
Diagnostic
* http://localhost:8080/api/v1/ping Show app version

Friend
* POST http://localhost:8080/api/v1/friend/connect      Connect as friend between two email
* POST http://localhost:8080/api/v1/friend/list         Show list of friend of an email
* POST http://localhost:8080/api/v1/friend/common       Show list of common friend between two email

Notfication
* POST http://localhost:8080/api/v1/notification/subscribe  Subscribe update to an email
* POST http://localhost:8080/api/v1/notification/update     Get list of subscribe Email
* POST http://localhost:8080/api/v1/notification/block      Block an email address 

