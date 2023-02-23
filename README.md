# hello-go

## Prepare

```
$ docker compose up
$ docker compose exec app bash
```

## Hello, World!

```
$ cd hello

$ go fmt *.go

$ go run main.go
Hello, World!

$ go build main.go
$ ./main
Hello, World!
```

## Test

```
$ cd test
$ go test -coverprofile=cover.out
$ go tool cover -html=cover.out -o coverage.html

# Mac
$ open test/coverage.html
```

## With modules

```
$ cd package

$ go fmt

$ go run package
```

## Cross compile

```
$ cd hello

$ GOOS=darwin GOARCH=arm64 go build main.go

# Mac
$ ./hello/main 
Hello, World
```

## Install packages

```
$ go install github.com/go-task/task/v3/cmd/task@latest

$ ls /go/bin
task

$ task
task: No Taskfile found in "/app" (or any of the parent directories). Use "task --init" to create a new one
```

## Acknowledgements

- https://note.com/umotion/n/n45f63f59bed5
- https://tyablog.net/2020/01/05/difference-between-interface-and-pointer-in-golang/
- https://qiita.com/k0kubun/items/1b641dfd186fe46feb65
