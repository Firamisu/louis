.PHONY: run
run:
	go run ./cmd

.PHONY: gen-templ
gen-templ:
	templ generate ./internal/views
