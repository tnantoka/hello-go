# hello-go

## Hello, World!

```
$ docker compose up

$ docker compose exec app go run hello/main.go
Hello, World!

$ docker compose exec app go build -o hello/main hello/main.go
$ docker compose exec app ./hello/main
Hello, World!
```

## With modules

```
$ docker compose exec app bash -c "cd package && go run package"
```

## Format

```
$ docker compose exec app bash -c "go fmt hello/*"
```

## Cross compile

```
$ docker compose exec -e GOOS=darwin -e GOARCH=arm64 app go build -o hello/main hello/main.go
$ ./hello/main 
Hello, World
```

## Install packages

```
$ docker compose exec app go install github.com/go-task/task/v3/cmd/task@latest

$ docker compose exec app ls /go/bin
task

$ docker compose exec app task
task: No Taskfile found in "/app" (or any of the parent directories). Use "task --init" to create a new one
```

## Acknowledgements

- https://note.com/umotion/n/n45f63f59bed5
