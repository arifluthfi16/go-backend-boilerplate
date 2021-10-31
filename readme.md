# MVC Boilerplate for Golang Web App

This boilerplate are intended to be a starting point on starting a new go web app
that uses mvc structure

Technology Used:
1. Gin
2. Gorm - PostgreSQL as Default Driver

By default database are disabled, you will need to configure the database 
connection first.

Notes :
```go
import (
    "fmt"
    "github.com/arifluthfi16/gomvcboilerplate/config"
    "github.com/arifluthfi16/gomvcboilerplate/routers"
    "github.com/arifluthfi16/gomvcboilerplate/services"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)
```
If you wanted to push the code into another repo, it is best to change the module name
and also the import statement.

To run the project:
1. Clone the project

```bash
git clone https://github.com/arifluthfi16/gomvcboilerplate/
```
   
2. Install modules
```bash
go mod download
```

3. Run the project
```bash
go run main.go
```
