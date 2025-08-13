.PHONY: swagger test test-race cover cover-html

swagger:
	swag init -g ./internal/router/router.go -d . -o ./docs -ot json,yaml

test:
	go test ./...

test-race:
	go test ./... -race

cover:
	go test ./... -race -covermode=atomic -coverprofile=coverage.out
	go tool cover -func=coverage.out

cover-html: test
	go tool cover -html=coverage.out -o coverage.html
