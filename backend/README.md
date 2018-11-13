# I'm Telling the Story server

This uses Golang. You should install <https://gin-gonic.github.io/gin/> and golang <https://golang.org> to run this server.

To install just do the following:

1. `vagrant up`
1. `cd /backend`
1. `go get github.com/gin-gonic/gin`
1. run `go run test.go`
1. run the executable that spawns forth

## Current API Specs

| Method    | Route     | Result    |
| ----------|:----------|:---------:|
| GET       |/test      |Returns a test JSON|
|GET|/test/:test|Returns a json with echo value equal to value inputted in test

---

## Future API

| Method    | Route     | Result    |
| ----------|:----------|:---------:|
| GET       |/story      |Returns current story in JSON format.|