# GO Friends Management REST API
A Restful API for simple Friends Management application with GO, using **gin-gonic/gin** (A most popular HTTP framework) and **gorm** (The fantastic ORM library for Golang)

### Structure
```
friends_management_v2/
┣ controller/
┃ ┣ common_respone/
┃ ┃ ┗ http_respone.go
┃ ┣ friendship/
┃ ┃ ┣ friendship.go
┃ ┃ ┣ friendship_controller.go
┃ ┃ ┗ friendship_controller_test.go
┃ ┣ user/
┃ ┃ ┣ user.go
┃ ┃ ┣ user_controller.go
┃ ┃ ┗ user_controller_test.go
┃ ┗ routes.go
┣ docs/
┃ ┣ docs.go
┃ ┣ swagger.json
┃ ┗ swagger.yaml
┣ migrations/
┃ ┣ 1_create_table.up.sql
┃ ┣ 1_delete_table.down.sql
┃ ┗ init_migration.go
┣ services/
┃ ┣ friendship/
┃ ┃ ┣ friendship.go
┃ ┃ ┣ friendship_mock_service.go
┃ ┃ ┣ friendship_service.go
┃ ┃ ┗ friendship_service_test.go
┃ ┗ user/
┃   ┣ user.go
┃   ┣ user_mock_service.go
┃   ┣ user_service.go
┃   ┗ user_service_test.go
┣ utils/
┃ ┣ connection_db.go
┃ ┗ helpers.go
┣ go.mod
┣ go.sum
┣ main.go
┣ README.md
┗ swagger.sh
```

## Installation & Run
```go
dbDriver := "postgres"
dbUser := "postgres"
dbPass := "postgres"
dbName := "FriendsManagement"
dbPort := "5432"
dbHost := "localhost"