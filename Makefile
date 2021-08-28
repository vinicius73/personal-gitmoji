setup:
	go get ./...
	go mod vendor

generate:
	go run main.go > dist/gitmoji.json
	cat dist/gitmoji.json