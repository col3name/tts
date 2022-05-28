fmt:twitch tts
	go fmt
build: fmt
	go build -o main.exe main.go
test:
	go test ./...