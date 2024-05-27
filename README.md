A small project to experiment with the Gin Framework

Tutorial followed: https://go.dev/doc/tutorial/web-service-gin

This API provides access to a store selling vintage recordings on vinyl.
 ### Endpoints
```bash
/albums

GET – Get a list of all albums, returned as JSON.
POST – Add a new album from request data sent as JSON.
/albums/:id

GET – Get an album by its ID, returning the album data as JSON.

```


### Running etc.
Get GO packages:

```bash
$ go get . 
```

Run pkg:

```bash
$ go run .
```

Test with:

```bash
# Get all albums
curl http://localhost:8080/albums

# Get album by ID
curl http://localhost:8080/albums/2

# POST a new album record
curl http://0.0.0.0:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

### Running with Docker

Use the Makefile for this:

```bash
make up

make down
```

#### Run tests:

```bash
make up

# then

docker-compose exec -it app /bin/bash

# then

go test ./tests/...
```

### CI/CD Pipeline
[WIP] A Github Actions pipeline has been added to this project. 
It will be responsible for:
    1. Building Docker images
    2. Pushing them to a registry
    3. Running tests within the container
    





