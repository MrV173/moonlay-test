>##  This App is using PostgreSQL Database to store and manage to do list / sublist (Create, Read, Update, Delete)
---
# Let's Start

--

- Run this code to start the app

```bash
go run main.go
```

- Run this code to test the to do list unit test

```bash
go test -v ./tests/to_do_list_unit_test.go 
```

- Run this code to test the sub to do list unit test

```bash
go test -v ./tests/sub_to_do_list_unit_test.go
```

- if you want to using MySQL , write this code

    > 	File: `pkg/database/database.go
```go
    package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/moonlay?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}
```
