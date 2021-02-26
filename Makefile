all: clean build run

build:
	CGO_ENABLED=0 go build -trimpath -ldflags '-w -s -extldflags "-static"' -o dist/app ./main.go
run:
	go run ./main.go

clean:
	
