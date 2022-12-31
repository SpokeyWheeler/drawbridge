SHELL = /usr/local/bin/bash
local:
	go build -o cmd/drawbridge/drawbridge cmd/drawbridge/main.go

run:
	go run cmd/drawbridge/main.go

clean:
	rm cmd/drawbridge/drawbridge

linux:
	GOOS=linux GOARCH=amd64 go build -o cmd/drawbridge/drawbridge cmd/drawbridge/main.go
