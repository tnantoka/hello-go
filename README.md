# hello-go

```
$ docker compose up

$ docker compose exec app bash -c "go fmt **/*"

$ docker compose exec app go run hello/main.go
Hello, World!

$ docker compose exec app go build hello/main.go
$ docker compose exec app ./main
Hello, World!
```

## Acknowledgements

- 