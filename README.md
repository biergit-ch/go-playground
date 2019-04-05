# GO Rest Demo

### Dependencies

```
go get github.com/gorilla/mux                   // needed for http routing
go get github.com/jinzhu/gorm                   // needed for database as orm
go get github.com/go-sql-driver/mysql           // mysql driver
go get github.com/jinzhu/gorm/dialects/mysql    // mysql dialect
```

### Run
```
$ docker-compose up 
$ go build && ./go-basics

navigate to: http://localhost:8000/
navigate to: http://localhost:8000/users
```


### Sources

__Conventions__

* http://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/
* https://golang.org/doc/effective_go.html

__Layout & Architecture__

* https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
* https://github.com/golang-standards/project-layout
* https://hackernoon.com/golang-clean-archithecture-efd6d7c43047


__Interfaces__

* https://blog.chewxy.com/2018/03/18/golang-interfaces/

__Persistence__

* http://doc.gorm.io
* https://www.alexedwards.net/blog/organising-database-access

__Networking__

* http://www.gorillatoolkit.org/pkg/mux
* https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo

__Configuration__

* http://goinbigdata.com/persisting-application-configuration-in-golang/



