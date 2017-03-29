GOCMD=go
GOGET=$(GOCMD) get
GOBUILD=$(GOCMD) build
GOBUILDPROD=$(GOBUILD) -ldflags "-linkmode external -extldflags -static" 
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
DOCKER=docker
MYSQLDB=userservice_db
CONSUL=consul

deps: 
	$(GOCMD) get ./...

build:
	$(GOBUILD)

buildprod:
	$(GOBUILDPROD) -v -o userservice

clean:
	$(GOCLEAN) -n -i -x
	rm -f $(GOPATH)/bin/userservice
	rm -rf userservice

consul:
	$(CONSUL) agent --dev


docker:  buildprod
	$(DOCKER) build -t bketelsen/userservice .

test:
	$(GOTEST) -v ./db -race
	$(GOTEST) -v ./handler -race

db-up:  db-down
	$(DOCKER) run -d --name=$(MYSQLDB) -p 3306:3306 -e MYSQL_ROOT_PASSWORD=s3cretStuff -e MYSQL_DATABASE=users mysql:5.7

db-down: 
	$(DOCKER) stop $(MYSQLDB)
	$(DOCKER) rm $(MYSQLDB)

setup-dev: deps db-up

teardown-dev: clean db-down

run: setup-dev
	sleep 10
	$(GOCMD) run main.go --database_url="root:s3cretStuff@tcp(127.0.0.1:3306)/users"
