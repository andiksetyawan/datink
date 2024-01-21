# datink
Dating App

- [Installation](#installation)
- [Usage](#usage)
- [API DOCUMENTATION](#api-documentation)

## Installation

### Requirements

1. [Go](https://golang.org/doc/install) 1.17+
2. [Docker](https://docs.docker.com/engine/install/) 

### Setting up environment

For a starter, copy and rename file `env.example` to `.env`

## Getting Started
## Usage

### Development

#### Run 
Install Dependency
```
go mod tidy
```
And then build
```
make build
```

For run web server:
```
./datink server
```

For run db migration:
```
./datink db init
```
```
./datink db migrate
```

#### Deploy on container
Run all service with docker-compose:
```
docker-compose -f docker-compose.yaml up
```

##### Run restapi server:
#### Make requests:
HTTP/1.1 GET API with curl
```
$ curl \
--header "Content-Type: application/json" \
http://localhost:9000/
```

reponse:
```
datink be service
```


### Make test:
```
make test
```
or
```
go test ./...
```

### API DOCUMENTATION:
SWAGGER / Open API Spec:
- http://localhost:9000/docs/index.html#/

POSTMAN COLLECTION
- [```scripts/datink.postman_collection.json```](./scripts/datink.postman_collection.json)