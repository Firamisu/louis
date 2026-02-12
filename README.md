# Louis

Louis is a small HTTP API that looks up English words using the public dictionary API (https://api.dictionaryapi.dev/). It exposes a health endpoint and a route that returns an HTML page showing a word's phonetic, origin and meanings.

## Prerequisites
- [Templ](https://github.com/a-h/templ) - html templates in go 

## How to run locally
- Start the server with `make run`:
```
make run
```

- By default the server listens on port `8080` Endpoints:
  - `GET /health` — currently just returns string: "healthy"
  - `GET /{word}` — return dictionary entry page for `word`


## Running tests
- Run all tests:
```
go test ./... -v
```


## Regenerate `templ` templates:
- After changing any `.templ` files, run `make gen-templ` to regenerate the view code.
