

create-project:
	go mod init github.com/datastx/IceBridge

tidy:
	go mod tidy

run:
	go run src/main.go