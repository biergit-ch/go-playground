# GO Demo REST API

### Prerequisite

__GO Setup on OSX__

* `golang` http://sourabhbajaj.com/mac-setup/Go/README.html
* `dependency-management` https://github.com/golang/dep

__GO Setup on Windows__

* `golang` http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/
* `dependency-management` https://golang.github.io/dep/docs/installation.html

### Install Dependencies

```
$ dep status
$ dep ensure
````

### Run
```
$ docker-compose up 
$ go build && ./go-basics
```
* http://localhost:8000/users
* http://localhost:8000/groups
* http://localhost:8000/transactions

### Layout

```
  +--------------------------------------------------+   +-----------------+
  |                      api                         |   |     vendor      |
  |--------------------------------------------------|   |-----------------|
  | +---------+  +---------+  +---------+  +-------+ |   |                 |
  | |         |  |         |  |         |  |       | |   |                 |
  | |         |  |         |  |         |  |       | |   |                 |
  | |         |  |         |  |         |  |       | |   |                 |
  | |         |  |         |  |         |  |       | |   |                 |
  | | handler |+>| service |+>|  repo   |+>|  db   | |   |  vendor pkgs    |
  | |         |  |         |  |         |  |       | |   |                 |
  | |         |  |         |  |         |  |       | |   |                 |
  | |         |  |         |  |         |  |       | |   |                 |
  | |         |  |         |  |         |  |       | |   |                 |
  | +---------+  +---------+  +---------+  +-------+ |   |                 |
  +--------------------------------------------------+   +-----------------+
```


### Logging

Use this logging mechanism to be indexed by a logging system `https://github.com/sirupsen/logrus`

```
import (
	log "github.com/sirupsen/logrus"
)

log.WithFields(log.Fields{
		"key": "value",
	}).Debug("Save user in repository")
	
Result: time="2019-04-08T23:31:03+02:00" level=debug msg="Save user in repository" user_id=0
```


### Sources

__Basics__

* https://hackernoon.com/basics-of-golang-for-beginners-6bd9b40d79ae

__Conventions__

* http://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/
* https://golang.org/doc/effective_go.html

__Layout & Architecture__

* https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
* https://github.com/golang-standards/project-layout
* https://hackernoon.com/golang-clean-archithecture-efd6d7c43047
* https://rakyll.org/style-packages/
* https://www.reddit.com/r/golang/comments/8g26il/what_is_the_recommended_go_project_folder/


__Interfaces__

* https://blog.chewxy.com/2018/03/18/golang-interfaces/

__Persistence__

* http://doc.gorm.io
* https://www.alexedwards.net/blog/organising-database-access

__Networking__

* http://www.gorillatoolkit.org/pkg/mux
* https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo

__Json__

* https://golang.org/pkg/encoding/json

__Configuration__

* http://goinbigdata.com/persisting-application-configuration-in-golang/

__Best Practices__

* https://peter.bourgon.org/go-in-production/





