build:
	go build -o bin/ccwc main.go

clean:
	rm -rf bin/

test:
	go test -v