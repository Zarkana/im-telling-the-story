# I'm Telling the Story server

## Installation

1. `vagrant up`
1. `go get github.com/gin-gonic/gin`
1. `cd /backend`
1. run `go run test.go`

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
