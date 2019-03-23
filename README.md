# Requirements
- go get github.com/go-sql-driver/mysql
- go get github.com/gin-gonic/gin
- go get github.com/jmoiron/sqlx
# Install
- go get github.com/juankis/registry
# Run
- In registry folder run in the console: go run api/src/main/main.go
# Example
- insert: POST http://registry.compute.amazonaws.com:8080/registry
- Body example Post request:
``` json
{
    "name": "Ricardo Tapia",
    "email": "estilista@gmail.com",
    "cel": "78687698",
    "type_customer":"1" // 1 estilista, 2 cliente
}
```
