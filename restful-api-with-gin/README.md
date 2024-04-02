# Developing a RESTful API with Go and Gin

Link to tutorial: https://go.dev/doc/tutorial/web-service-gin

## Testing

### Run app

```sh
go run .
```

### Get all albums

```sh
curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
```

## Get a specific album

```sh
curl http://localhost:8080/albums/2
```

### Add a new album

```sh
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```
