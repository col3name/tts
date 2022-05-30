fmt:
	go fmt
lint:
	golangci-lint run
build: fmt
	go build -o main.exe main.go
releaseWin: build
	rar a -r windows-amd64.rar main.exe
buildLinux: fmt
	set GOOS=linux
	set GOARCH=amd64
	go build -o main main.go
releaseLinux: buildLinux
	rar a -r linux-amd64.rar main
buildMacos: fmt
	set GOOS=darwin
	set GOARCH=arm64
	go build -o main main.go
releaseMacos: buildMacos
	rar a -r macos-arm64.rar main
release: test lint releaseWin releaseLinux releaseMacos
test:
	go test ./...