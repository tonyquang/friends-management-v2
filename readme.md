# GO Friends Management REST API
A Restful API for simple Friends Management application with GO, using **gin-gonic/gin** (A most popular HTTP framework) and **gorm** (The fantastic ORM library for Golang)

## Structure

* [controller/](.\friends_management_v2\controller)
  * [common_respone/](.\friends_management_v2\controller\common_respone)
    * [http_respone.go](.\friends_management_v2\controller\common_respone\http_respone.go)
  * [friendship/](.\friends_management_v2\controller\friendship)
    * [friendship.go](.\friends_management_v2\controller\friendship\friendship.go)
    * [friendship_controller.go](.\friends_management_v2\controller\friendship\friendship_controller.go)
    * [friendship_controller_test.go](.\friends_management_v2\controller\friendship\friendship_controller_test.go)
  * [user/](.\friends_management_v2\controller\user)
    * [user.go](.\friends_management_v2\controller\user\user.go)
    * [user_controller.go](.\friends_management_v2\controller\user\user_controller.go)
    * [user_controller_test.go](.\friends_management_v2\controller\user\user_controller_test.go)
  * [routes.go](.\friends_management_v2\controller\routes.go)
* [docs/](.\friends_management_v2\docs)
  * [docs.go](.\friends_management_v2\docs\docs.go)
  * [swagger.json](.\friends_management_v2\docs\swagger.json)
  * [Swagger.PNG](.\friends_management_v2\docs\Swagger.PNG)
  * [swagger.yaml](.\friends_management_v2\docs\swagger.yaml)
  * [UnitTest.PNG](.\friends_management_v2\docs\UnitTest.PNG)
* [migrations/](.\friends_management_v2\migrations)
  * [1_create_table.up.sql](.\friends_management_v2\migrations\1_create_table.up.sql)
  * [1_delete_table.down.sql](.\friends_management_v2\migrations\1_delete_table.down.sql)
  * [init_migration.go](.\friends_management_v2\migrations\init_migration.go)
* [services/](.\friends_management_v2\services)
  * [friendship/](.\friends_management_v2\services\friendship)
    * [friendship.go](.\friends_management_v2\services\friendship\friendship.go)
    * [friendship_mock_service.go](.\friends_management_v2\services\friendship\friendship_mock_service.go)
    * [friendship_service.go](.\friends_management_v2\services\friendship\friendship_service.go)
    * [friendship_service_test.go](.\friends_management_v2\services\friendship\friendship_service_test.go)
  * [user/](.\friends_management_v2\services\user)
    * [user.go](.\friends_management_v2\services\user\user.go)
    * [user_mock_service.go](.\friends_management_v2\services\user\user_mock_service.go)
    * [user_service.go](.\friends_management_v2\services\user\user_service.go)
    * [user_service_test.go](.\friends_management_v2\services\user\user_service_test.go)
* [utils/](.\friends_management_v2\utils)
  * [connection_db.go](.\friends_management_v2\utils\connection_db.go)
  * [helpers.go](.\friends_management_v2\utils\helpers.go)
* [docker-compose.yml](.\friends_management_v2\docker-compose.yml)
* [Dockerfile](.\friends_management_v2\Dockerfile)
* [go.mod](.\friends_management_v2\go.mod)
* [go.sum](.\friends_management_v2\go.sum)
* [main.go](.\friends_management_v2\main.go)
* [README.md](.\friends_management_v2\README.md)
* [swagger.sh](.\friends_management_v2\swagger.sh)

## Installation & Run
```go
host     := "db"
port     := 5432
user     := "postgres"
password := "postgres"
dbname   := "FriendsManagement"
```

For run docker-compose, run these following commands in project's root folder:

```bash
docker-compose build
docker-compose up
```

#### API Endpoint
```bash
http://localhost:3000/swagger/index.html
```

## API Documentation
This is API self documentation by using Swagger. You can test all of them by expand specific api then click on try it out button.

![Swagger API Documentation](docs/Swagger.PNG)

## Test Coverage
All APIs have been tested carefully by mocking strategy. 
![Swagger API Documentation](docs/UnitTest.PNG)

## Achievement

- [x] Write the tests for all APIs.
- [x] Organize the code with packages
- [x] Make docs with Swagger
- [x] Building a deployment process 