.PHONY: lint
lint:
	golangci-lint run --timeout 5m

.PHONY: test
test:
	go test ./...

.PHONY: dev/tailwind
dev/tailwind:
	tailwindcss -i ui/styles.pcss -o static/styles.css --watch

.PHONY: dev/templ
dev/templ:
	templ generate --watch --proxy="http://localhost:6969" --cmd="go run . serve" --open-browser=false

.PHONY: dev/serve
dev/serve:
	make -j2 dev/tailwind dev/templ
